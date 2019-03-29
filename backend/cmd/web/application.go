package main

import (
	"github.com/gorilla/mux"
	"github.com/isa-mrs-tim6/Projekat/pkg/db/gorm/postgre"
	"log"
)

type Application struct {
	ErrorLog      *log.Logger
	InfoLog       *log.Logger
	Router        *mux.Router
	Store         *postgre.Store
	EmailAddress  string
	EmailPassword string
}
