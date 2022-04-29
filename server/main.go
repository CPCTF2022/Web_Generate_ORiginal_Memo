package main

import (
	"os"
	"strconv"
)

func main() {
	flag, ok := os.LookupEnv("FLAG")
	if !ok {
		panic("FLAG is not set")
	}

	adminPassword, ok := os.LookupEnv("ADMIN_PASSWORD")
	if !ok {
		panic("ADMIN_PASSWORD is not set")
	}

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

	err = dbInit(flag, adminPassword, user, pass, host, DBPort, database)
	if err != nil {
		panic(err)
	}

	sessionSecret, ok := os.LookupEnv("SESSION_SECRET")
	if !ok {
		panic("SESSION_SECRET is not set")
	}

	staticRoot, ok := os.LookupEnv("STATIC_ROOT")
	if !ok {
		panic("STATIC_ROOT is not set")
	}

	addr, ok := os.LookupEnv("APP_ADDR")
	if !ok {
		panic("APP_ADDR is not set")
	}

	err = startServer(addr, staticRoot, sessionSecret)
	if err != nil {
		panic(err)
	}
}
