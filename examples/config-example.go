package main

import (
	"github.com/thane421/gofig/pkg/gofig"
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
	err := gofig.LoadConfig(Config, "example-project")

	if err != nil {
		panic(err)
	}

	fmt.Println(
		Config.DatabaseName,
		Config.DatabaseHost,
		Config.DatabasePort,
		Config.DatabaseUser,
		Config.DatabasePassword,
	)

	// In another package you could import the Config object and use it's values which should have been loaded
	// 		from either ~/example-project.yml or from the prompts if that file did not exist.
	//		And if the yml file did not exist, it will have been created with the values provided.
}
