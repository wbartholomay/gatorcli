package main

import (
	"github.com/wbartholomay/gatorcli/internal/config"
	"github.com/wbartholomay/gatorcli/internal/database"
)

type state struct {
	cfg *config.Config
	db *database.Queries
}