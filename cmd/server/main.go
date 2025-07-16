package main

import (
	"log"
	"os"

	"git.eveutil.org/fabric/everate/internal/logger"
	"git.eveutil.org/fabric/everate/internal/middleware"
	"go.uber.org/zap"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func main() {
	// 1. Load config
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// 2. Create the logger instance
	isDev := os.Getenv("DEV") == "true"
	appLogger, err := logger.New(isDev)
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	defer appLogger.Sync() // This is still important!

	// 3. Create Echo instance
	e := echo.New()

	// 4. Register your custom Zap logger middleware
	// This replaces `e.Use(echomiddleware.Logger())`
	e.Use(middleware.ZapRequestLogger(appLogger))

	// 5. Register other middleware
	sessionSecret := os.Getenv("SESSION_SECRET")
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(sessionSecret))))

	// 6. Routes


	// Start the server
	appLogger.Info("Starting the server...", zap.String("address", ":1323"))
	if err := e.Start(":1323"); err != nil {
		appLogger.Fatal("Server failed to start", zap.Error(err))

	}
}
