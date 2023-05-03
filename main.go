package main

import (
	"net/http"

	"github.com/fiware/VCVerifier/logging"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fiber := fiber.New()

	//tokenHandler := ishare.NewTokenHandler()

	// TODO inject error handler
	issuerService := IssuerService{tokenHandler: nil}
	router := NewApiController(&issuerService)

	for _, route := range router.Routes() {
		switch route.Method {
		case http.MethodGet:
			logging.Log().Infof("Setup path: %s", route.Pattern)
			fiber.Get(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			fiber.Post(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			fiber.Put(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			fiber.Patch(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			fiber.Delete(route.Pattern, route.HandlerFunc)
		}
	}

	fiber.Listen(":9090")
}
