package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

var wb *Wolfbot

func main() {
	wb = &Wolfbot{}
	log.Println("Starting up...")
	wb.Configure()
	wb.Connect()
	wb.Run()

	log.Println("Wolfbot is running.  Press CTRL-C to exit.")

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	wb.Stop()
}
