package main

import (
	"fmt"

	"github.com/wbartholomay/gatorcli/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil { panic(err) }

	err = cfg.SetUser("will")
	if err != nil { panic(err) }

	cfg, err = config.Read()
	if err != nil { panic(err) }

	fmt.Println(cfg)
}