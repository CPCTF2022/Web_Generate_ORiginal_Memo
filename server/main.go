package main

import (
	"os"
	"strconv"
)

func main() {
	user, ok := os.LookupEnv("DB_USERNAME")
	if !ok {
		panic("DB_USERNAME is not set")
	}

	pass, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		panic("DB_PASSWORD is not set")
	}

	host, ok := os.LookupEnv("DB_HOSTNAME")
	if !ok {
		panic("DB_HOSTNAME is not set")
	}

	strDBPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		panic("DB_PORT is not set")
	}
	DBPort, err := strconv.Atoi(strDBPort)
	if err != nil {
		panic("invalid DB_PORT")
	}

	database, ok := os.LookupEnv("DB_DATABASE")
	if !ok {
		panic("DB_DATABASE is not set")
	}

	err = dbInit(user, pass, host, DBPort, database)
	if err != nil {
		panic(err)
	}

	staticRoot, ok := os.LookupEnv("STATIC_ROOT")
	if !ok {
		panic("STATIC_ROOT is not set")
	}

	addr, ok := os.LookupEnv("APP_ADDR")
	if !ok {
		panic("DB_PORT is not set")
	}

	err = startServer(addr, staticRoot)
	if err != nil {
		panic(err)
	}
}
