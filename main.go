package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/wbartholomay/gatorcli/internal/config"
	"github.com/wbartholomay/gatorcli/internal/database"
)

var handlers map[string]func(*state, command) error = map[string]func(*state, command) error{
	"login" : handlerLogin,
	"register" : handlerRegister,
	"reset" : handlerReset,
	"users" : handlerUsers,
	"agg" : handlerAgg,
	"addfeed" : handlerAddFeed,
	"feeds" : handlerFeeds,
	"follow" : handlerFollow,
}

func main() {
	cfg, err := config.Read()
	if err != nil { panic(err) }

	//connect to database
	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil { panic(err) }

	//intialize state struct
	dbQueries := database.New(db)
	s := state{
		cfg : &cfg,
		db : dbQueries,
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
	if err != nil { 
		fmt.Println(err)
		os.Exit(1)
	 }

}