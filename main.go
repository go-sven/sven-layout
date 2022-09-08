package main

import (
	"github/go-sven/sven-layout/app/conf"
	"github/go-sven/sven-layout/cmd"
)

func main()  {
	//init config
	config := conf.InitConf("./config")

	// run
	cmd.Run(config)
}
