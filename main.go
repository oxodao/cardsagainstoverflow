package main

import (
	"fmt"
	"os"

	"github.com/oxodao/cao/cmd"
	"github.com/oxodao/cao/config"
	"github.com/oxodao/cao/services"
)

func main() {
	err := config.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = services.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd.Execute()
}
