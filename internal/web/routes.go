package web

import (
	"app/internal/store"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, hitsStore *store.HitsStore) {
	app.Get("/", GetMain(hitsStore))
}
