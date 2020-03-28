package main

import (
	"fmt"
	"net/http"

	dbConn "github.com/krishnakumarkp/goapp/adapter/gorm"
	"github.com/krishnakumarkp/goapp/app/app"
	"github.com/krishnakumarkp/goapp/app/router"
	"github.com/krishnakumarkp/goapp/config"
	lr "github.com/krishnakumarkp/goapp/util/logger"
	vr "github.com/krishnakumarkp/goapp/util/validator"
)

func main() {
	appConf := config.AppConfig()

	logger := lr.New(appConf.Debug)

	db, err := dbConn.New(appConf)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
		return
	}
	if appConf.Debug {
		db.LogMode(true)
	}

	validator := vr.New()
    application := app.New(logger, db, validator)

	appRouter := router.New(application)

	address := fmt.Sprintf(":%d", appConf.Server.Port)
	logger.Info().Msgf("Starting server %s\n", address)
	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failed")
	}
}
