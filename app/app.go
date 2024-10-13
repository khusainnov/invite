package app

import (
	"context"

	"gitlab.com/khusainnov/invite-app/app/api"
	"gitlab.com/khusainnov/invite-app/app/config"
	"gitlab.com/khusainnov/invite-app/app/infra/server"
	"gitlab.com/khusainnov/invite-app/app/infra/storage"
	"gitlab.com/khusainnov/invite-app/app/processor/invite"
	"gitlab.com/khusainnov/invite-app/app/repository"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type App struct {
	Cfg    *config.Config
	Log    *zap.Logger
	Server *server.Server
	done   chan struct{}
}

func New(cfg *config.Config) *App {
	app := &App{
		Cfg: cfg,
	}

	log, err := app.CreateLogger()
	if err != nil {
		panic(err)
	}

	app.Log = log

	app.Log.Info("connecting to db")
	conn, err := storage.New(app.Log, cfg.DB)
	if err != nil {
		app.Log.Error("failed to connect to db")
		panic(err)
	}

	app.Log.Info("creating new handler")
	handlers := api.New(
		invite.New(
			app.Log,
			conn,
			repository.NewCustomerRepo(),
			repository.NewEventRepo(),
		),
	)

	app.Log.Info("creating new server")
	srv := server.New(app.Cfg.Server)
	app.Server = srv
	if err := app.Server.Init(handlers); err != nil {
		panic(err)
	}

	return app
}

func (a *App) CreateLogger() (*zap.Logger, error) {
	logCfg := zap.NewProductionConfig()
	logCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logCfg.EncoderConfig.TimeKey = "timestamp"

	return logCfg.Build(zap.AddCaller())
}

func (a *App) Run() {
	if a.Server != nil {
		go func() {
			defer func() { a.done <- struct{}{} }()
			a.Log.Info("starting server")
			err := a.Server.Run()
			if err != nil {
				a.Log.Error(err.Error(), zap.Error(err))
			}
		}()
	}

	a.Log.Info("invite-app running!")
	<-a.done
	a.stop()
}

func (a *App) stop() {
	a.Log.Info("invite-app stopping...")

	if err := a.Server.Stop(context.Background()); err != nil {
		a.Log.Error(err.Error(), zap.Error(err))
	}
}
