package main

import (
	"github.com/devthane/autgofig/pkg/autgofig"
	"fmt"
)

type ExampleConfig struct {
	DatabaseName     string
	DatabaseHost     string
	DatabasePort     int
	DatabaseUser     string
	DatabasePassword string
}

var Config *ExampleConfig

func main() {
	Config = new(ExampleConfig)

	if err := autgofig.LoadConfig(Config, "example-project"); err != nil {
		panic(err)
	}

	fmt.Println(
		Config.DatabaseName,
		Config.DatabaseHost,
		Config.DatabasePort,
		Config.DatabaseUser,
		Config.DatabasePassword,
	)
}
