package routs

import (
	"Pet_store/internal/controller/app"

	"github.com/go-chi/chi/v5"
	//"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
)

// func NewRouter(h Handler) *chi.Mux {
// 	r := chi.NewRouter()

// 	r.Get("/swagger/*", httpSwagger.WrapHandler)
// 	r.Post("/api/login", h.LoginHandler)
// 	r.Post("/api/register", h.RegisterHandler)

// 	r.Group(func(r chi.Router) {
// 		tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
// 		r.Use(jwtauth.Verifier(tokenAuth))
// 		r.Use(jwtauth.Authenticator)
// 		r.Post("/api/address/search", h.AddressSearchHandler)
// 		r.Post("/api/address/geocode", h.AddressGeocodeHandler)
// 	})
// 	return r
// }

func InitAllRouts(app *app.App) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Get("/user/login", app.UH.Login())
	r.Get("/user/logout", app.UH.Logout())

	r.Group(func(r chi.Router) {
		r.Post("/pet", app.PH.Add())
		r.Put("/pet/", app.PH.FullUpdate())
		r.Get("/pet/findByStatus", app.PH.GetByStatus())
		r.Get("/pet/{petId}", app.PH.GetById())
		r.Delete("/pet/{petId}", app.PH.DeleteById())
		r.Post("/pet/{petId}", app.PH.FormUpdate())
	})

	r.Group(func(r chi.Router) {
		r.Get("/store/inventory", app.SH.InventoryHandler())
		r.Post("/store/order", app.SH.CreateOrderHandler())
		r.Get("/store/order/{orderId}", app.SH.GetOrderByIdHandler())
		r.Delete("/store/order/{orderId}", app.SH.DeleteOrderHandler())
	})

	r.Group(func(r chi.Router) {
		r.Get("/user/{username}", app.UH.GetByUsername())
		r.Put("/user/{username}", app.UH.Update())
		r.Delete("/user/{username}", app.UH.Delete())
		r.Post("/user", app.UH.Create())
		r.Post("/user/createWithArray", app.UH.CreateWithInputArray())
	})

	return r
}
