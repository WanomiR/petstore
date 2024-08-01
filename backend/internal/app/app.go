package app

import (
	"backend/internal/lib/e"
	"backend/internal/lib/rr"
	"backend/internal/modules"
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// include to use db drivers
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Server interface {
	Serve()
	Signal() <-chan os.Signal
	readConfig(string) error
	init() error
	connectToDB() (*sql.DB, error)
	routes() *chi.Mux
}

type App struct {
	Host        string
	Port        string
	JWTSecret   string
	server      *http.Server
	signalChan  chan os.Signal
	DSN         string
	DB          repository.Repository
	services    *modules.Services
	controllers *modules.Controllers
}

func NewApp() (a *App, err error) {
	defer func() { err = e.WrapIfErr("failed to init app", err) }()

	a = &App{}

	if err = a.init(); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Serve() {
	defer a.DB.Connection().Close()

	fmt.Println("Started server on port", a.Port)
	if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

func (a *App) Signal() <-chan os.Signal {
	return a.signalChan
}

func (a *App) CloseDBConn() error {
	return a.DB.Connection().Close()
}

func (a *App) readConfig(envPath ...string) (err error) {
	if len(envPath) > 0 {
		err = godotenv.Load(envPath[0])
	} else {
		err = godotenv.Load()
	}

	if err != nil {
		return e.WrapIfErr("can't read .env file", err)
	}

	a.Host = os.Getenv("HOST")
	a.Port = os.Getenv("PORT")
	a.JWTSecret = os.Getenv("JWT_SECRET")

	a.DSN = fmt.Sprintf( // database source name
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5\n",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	return nil
}

func (a *App) init() error {
	if err := a.readConfig(); err != nil {
		return err
	}

	conn, err := a.connectToDB()
	if err != nil {
		return err
	}
	a.DB = dbrepo.NewPostgresDBRepo(conn)

	a.services = modules.NewServices(a.DB, a.Host, a.Host, a.JWTSecret, a.Host)
	a.controllers = modules.NewControllers(
		a.services,
		rr.NewReadRespond(rr.WithMaxBytes(1<<10)),
	)

	a.server = &http.Server{
		Addr:         ":" + a.Port,
		Handler:      a.routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	a.signalChan = make(chan os.Signal, 1)
	signal.Notify(a.signalChan, syscall.SIGINT, syscall.SIGTERM)

	return nil
}

func (a *App) connectToDB() (conn *sql.DB, err error) {
	defer func() { err = e.WrapIfErr("failed to connect to database", err) }()

	conn, err = sql.Open("pgx", a.DSN)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
