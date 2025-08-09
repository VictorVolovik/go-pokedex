# Go Pokedex

A CLI tool built following Boot.Dev's [Build a Pokedex in Go](https://www.boot.dev/courses/build-pokedex-cli-golang), using the [PokeAPI](https://pokeapi.co/).

## Features

- Browse Pokemon locations
- Catch Pokemon with probability-based mechanics
- Inspect caught Pokemon details
- Maintain a personal Pokedex collection

## Prerequisites

- Go version 1.23.4 or later installed

## Getting Started

- `go run .`: Run the application
- `go build`: Build an executable (see [Go documentation](https://go.dev/doc/go1.23) for more info)
- `go test ./...`: Run all tests

## CLI Commands

- `help`: Displays a help message
- `map`: Get the next page of locations
- `mapb`: Get the previous page of locations
- `explore <location id or name>`: List all Pokemon in the specified location
- `catch <pokemon id or name>`: Attempt to catch the specified Pokemon
- `inspect <pokemon name>`: Check details of a caught Pokemon
- `pokedex`: List all caught Pokemon
- `exit`: Exit the Pokedex

## Notes

This project demonstrates foundational Go concepts such as working with APIs, building CLI tools, and handling user input effectively. It is a learning exercise and a fun way to explore Go while interacting with the PokeAPI.
