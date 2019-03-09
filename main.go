package main

import (
	"log"
	"flag"

	"github.com/emersion/go-smtp"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:1025", "Address to bind to")
	username := flag.String("username", "username", "Username to authenticate with")
	password := flag.String ("password", "password", "Password to authenticate with")
	outputPluginString := flag.String("output-plugin", "", "Output plugin configuration")
	flag.Parse()

	outputPlugin, err := parseOutputPlugin(*outputPluginString)
	if err != nil {
		log.Fatal(err)
	}

	backend := &Backend{Username: *username, Password: *password, Output: outputPlugin}

	server := smtp.NewServer(backend)
	server.Addr = *addr
	server.MaxIdleSeconds = 300
	server.MaxMessageBytes = 1024 * 1024
	server.MaxRecipients = 50
	server.AllowInsecureAuth = true

	log.Println("Starting server at", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
