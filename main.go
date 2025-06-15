package main

import (
	"fmt"
	"os"

	"github.com/wbartholomay/gatorcli/internal/config"
)

var handlers map[string]func(*state, command) error = map[string]func(*state, command) error{
	"login" : handlerLogin,
}

func main() {
	cfg, err := config.Read()
	if err != nil { panic(err) }

	//intialize state struct
	s := state{
		cfg : &cfg,
	}

	//initialize commands struct
	cmds := commands{
		handlers : handlers,
	}

	//read args from command line, create command struct
	if len(os.Args) < 2 {
		fmt.Println("No arguments were provided.")
		os.Exit(1)
	}

	//initialize cmd name and args
	cmdName := os.Args[1]
	var cmdArgs []string
	if len(os.Args) == 2 {
		cmdArgs = []string{}
	} else {
		cmdArgs = os.Args[2:]
	}


	cmd := command{
		name: cmdName,
		args: cmdArgs,
	}

	err = cmds.run(&s, cmd)
	if err != nil { panic(err) }

}