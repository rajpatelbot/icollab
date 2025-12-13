package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rajpatelbot/icollab/internal/config"
	"github.com/rajpatelbot/icollab/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set log format to include date, time with microseconds, and file line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)

	// Load environment variables
	config.InitEnv(".env")
	env := config.EnvConfig

	// Define GIN_MODE based on environment variable
	switch env.GIN_MODE {
	case "release":
		os.Setenv("GIN_MODE", "release")
	case "debug":
		os.Setenv("GIN_MODE", "debug")
	default:
		log.Fatalf("Invalid GIN_MODE: %s. Must be 'release' or 'debug'.", env.GIN_MODE)
	}

	// Server start
	srv := StartServer(env)

	// Connect to DB
	db := database.NewDatabase()
	if db != nil {
		log.Print("✅ Database connected successfully")
	}

	// Graceful shutdown handling can be added here if needed
	WaitForShutdown(srv)

	log.Println("✅ All active connections finished. Resources cleaned up.")
	log.Println("👋 Server exiting.")
}

func StartServer(env *config.Env) *http.Server {
	// Server initialization
	router := gin.Default()

	// Create an http server
	srv := &http.Server{
		Addr:    ":" + env.APP_PORT,
		Handler: router,
	}

	// Graceful server start in a goroutine
	go func() {
		log.Printf("🚀 Auth service is running at %s", srv.Addr)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("🚀 Auth service listen error: %v\n", err)
		}
	}()

	return srv
}

func WaitForShutdown(srv *http.Server) {
	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)

	// Catch OS interrupt signals
	signal.Notify(quit, os.Interrupt)

	// Block until a signal is received (e.g. Ctrl+C)
	<-quit

	log.Println("⚠️ Shutdown signal received. Started graceful shutdown...")

	// The context is used to inform the server; it has timeout to finish
	// Gives running requests a deadline to complete within timeout
	const timeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Attempt graceful server shutdown and stop accepting new requests
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Server forced to shutdown after: %v %v", timeout, err)
	}
}
