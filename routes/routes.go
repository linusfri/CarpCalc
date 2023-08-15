package routes

import (
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	authController "github.com/linusfri/calc-api/controllers/auth"
	userController "github.com/linusfri/calc-api/controllers/user"
	"github.com/linusfri/calc-api/middleware"
)

func InitRoutes() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   GetAllowedOrigins(),
		AllowCredentials: true,
	}))

	r.Use(middleware.Logger)

	// Users
	r.Route("/users", func(r chi.Router) {
		r.Post("/", userController.CreateUser)
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthJWT)
		r.Get("/users", userController.GetUsers)
	})

	r.Group(func(r chi.Router) {
		r.Post("/auth/login", authController.Login)
	})
	http.ListenAndServe(":8080", r)
}

func GetAllowedOrigins() []string {
	return strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
}
