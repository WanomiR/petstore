package app

import (
	"backend/internal/lib/e"
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
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
	Host       string
	Port       string
	server     *http.Server
	signalChan chan os.Signal
	DSN        string
	DB         repository.Repository
	// services
	// controllers
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
	log.Println("Connected to database")
	a.DB = dbrepo.NewPostgresDBRepo(conn)

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

func (a *App) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://%s:%s/swagger/doc.json", a.Host, a.Port)),
	))

	return r
}
