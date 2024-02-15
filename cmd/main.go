package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mustthink/microblog/internal/app"
)

func main() {
	application := app.New()
	go application.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.Stop()
}
