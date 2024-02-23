package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Ell534/go-dexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	// Set up channel to listen for interrupt and hang-up signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	// Run the REPL in a separate goroutine so that signal handling can proceed concurrently
	go startRepl(&cfg)

	// Block until a signal is received
	sig := <-sigChan

	// Perform any cleanup if necessary
	fmt.Printf("\nReceived %v signal. Exiting pokedex...\n", sig)

	// Exit the program
	os.Exit(0)
}
