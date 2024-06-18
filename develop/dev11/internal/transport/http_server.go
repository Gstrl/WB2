package transport

import (
	"WB2/develop/dev11/config"
	"WB2/develop/dev11/internal/database"
	"WB2/develop/dev11/internal/services"
	"net/http"
	"time"
)

func ServerRun(cfg *config.Config, cache *database.CacheEventRep) error {
	router := http.NewServeMux()
	app := services.NewApplication(cache)
	h := NewHTTPCalendarHandler(app)
	CustomRegisterHandlers(router, h)
	server := &http.Server{
		Addr:         cfg.Port,
		Handler:      router,
		ReadTimeout:  180 * time.Second,
		WriteTimeout: 180 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
