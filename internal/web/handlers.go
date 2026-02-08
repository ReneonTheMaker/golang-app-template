package web

import (
	"app/internal/store"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// index page handler
func GetMain(hitStore *store.HitsStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		// get number of hits from the store
		id := c.Cookies("id")
		if id == "" {
			id = uuid.New().String()
			c.Cookie(&fiber.Cookie{
				Name:  "id",
				Value: id,
			})
		}
		hits, exists := hitStore.GetHits(id)
		if !exists {
			hitStore.SetHits(id, 1)
			hits = 1
		} else {
			hitStore.SetHits(id, hits+1)
		}
		return c.Render("index", fiber.Map{
			"Hits":    hits,
		})
	}
}

