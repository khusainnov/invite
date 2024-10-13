package main

import (
	"gitlab.com/khusainnov/invite-app/app"
	"gitlab.com/khusainnov/invite-app/app/config"
)

func main() {
	cfg := config.NewFromEnv()
	inviteapp := app.New(cfg)
	inviteapp.Run()
}
