package main

import (
	"log"
	"os"
)

func info(msg string) {
	log.Print("INFO: " + msg)
}

func fail(msg string) {
	log.Print("ERROR: " + msg)
	os.Exit(1)
}

func warning(msg string) {
	log.Print("WARNING: " + msg)
}
