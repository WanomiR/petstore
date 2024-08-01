package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *App) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)

	r.Route("/pet", func(r chi.Router) {
		r.Use(a.requireAuthentication)
		r.Get("/{petId}", a.controllers.Pet.GetById)
		r.Post("/{petId}", a.controllers.Pet.UpdateWithForm)
		r.Delete("/{petId}", a.controllers.Pet.DeleteById)
		r.Post("/{petId}/uploadImage", a.controllers.Pet.UploadImage)
		r.Post("/", a.controllers.Pet.CreatePet)
		r.Put("/", a.controllers.Pet.UpdatePet)
		r.Get("/findByStatus", a.controllers.Pet.GetByStatus)
	})

	r.Route("/store", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(a.requireAuthentication)
			r.Get("/inventory", a.controllers.Store.GetInventory)
		})
		r.Post("/order", a.controllers.Store.CreateOrder)
		r.Get("/order/{orderId}", a.controllers.Store.GetOrderById)
		r.Delete("/order/{orderId}", a.controllers.Store.DeleteOrder)

	})

	r.Route("/user", func(r chi.Router) {
		r.Get("/{username}", a.controllers.User.GetByUsername)
		r.Put("/{username}", a.controllers.User.Update)
		r.Delete("/{username}", a.controllers.User.Delete)
		r.Post("/", a.controllers.User.Create)
		r.Post("/createWithArray", a.controllers.User.CreateWithArray)
		r.Get("/login", a.controllers.User.Login)
		r.Get("/logout", a.controllers.User.Logout)
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://%s:%s/swagger/doc.json", a.Host, a.Port)),
	))

	return r
}
