package config

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

// Store is the global session store
var Store *session.Store

func InitSession() {
	Store = session.New(session.Config{
		Expiration:     24 * time.Hour,
		Storage:        nil, // defaults to memory
		CookieHTTPOnly: true,
		CookieSecure:   false, // Set to true in production with HTTPS
	})
}
