package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitlab.com/seqone/mailtick/app"
	"gitlab.com/seqone/mailtick/db"
	"gitlab.com/seqone/mailtick/mailer"
	"gitlab.com/seqone/mailtick/scheduler"
)

var schedulerPause = flag.String("d", "1m", "pause between mail send")

func main() {
	flag.Parse()
	p, err := time.ParseDuration(*schedulerPause)
	if err != nil {
		panic(err)
	}
	log.Printf("Scheduler pause set to %s", p)

	// Setup app with dependencies
	db, err := db.New()
	if err != nil {
		panic(err)
	}
	app := app.New(db)
	// Start listening for request
	go app.Listen(":8181")

	// Setup and start sheduler
	scheduler := scheduler.New(db, mailer.New())
	go scheduler.Start(p)

	log.Printf("mailtick up and running on port 8181")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// Wait for interupr or sigterm
	<-c
	// Ensure graceful shutdown
	app.Shutdown()
}
