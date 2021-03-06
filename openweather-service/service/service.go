package service

import (
	"context"
	"log"
	"net/http"
	"time"
)

func GracefullyShutDown(ctx context.Context) (err error) {
	mux := http.NewServeMux()

	mux.Handle("/api/v1/weather/current", http.HandlerFunc(CurrentWeatherHandler))
	mux.Handle("/api/v1/weather/forecast", http.HandlerFunc(FiveDaysForecastHandler))
	mux.Handle("/health", http.HandlerFunc(HealthCheck))

	srv := &http.Server{
		Addr:    ":80",
		Handler: mux,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen:%s\n", err)
		}
	}()

	log.Printf("Server started on port 8080")


	<-ctx.Done()

	log.Printf("Server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("Server shutdown failed:%+s", err)
	}

	log.Printf("Server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}
