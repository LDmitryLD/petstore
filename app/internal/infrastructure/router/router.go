package router

import (
	"net/http"
	"projects/LDmitryLD/petstore/app/internal/modules"

	"github.com/go-chi/chi/v5"
)

func NewRouter(controllers *modules.Controllers) *chi.Mux {
	r := chi.NewRouter()

	setDefaultRoutes(r)

	r.Post("/pet/{petId}/uploadImage", controllers.Pet.UploadImage)
	r.Post("/pet", controllers.Pet.AddPet)
	r.Put("/pet", controllers.Pet.UpdateExistinPet)
	r.Get("/pet/findByStatus", controllers.Pet.FindByStatus)
	r.Get("/pet/{petId}", controllers.Pet.GetByID)
	r.Post("/pet/{petId}", controllers.Pet.UpdateByID)
	r.Delete("/pet/{petId}", controllers.Pet.DeletePet)

	r.Post("/user/createWithArray", controllers.User.CreateWithArray)
	r.Post("/user/createWithList", controllers.User.CreateWithArray)
	r.Get("/user/{username}", controllers.User.GetByUsername)
	r.Put("/user/{username}", controllers.User.Update)
	r.Delete("/user/{username}", controllers.User.Delete)
	r.Post("/user", controllers.User.Create)

	return r
}

func setDefaultRoutes(r *chi.Mux) {
	r.Get("/swagger", swaggerUI)
	r.Get("/public/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/public/", http.FileServer(http.Dir("../public"))).ServeHTTP(w, r)
	})
}
