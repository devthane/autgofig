# autgofig

[![GoDoc](https://godoc.org/github.com/devthane/autgofig?status.svg)](https://godoc.org/github.com/devthane/autgofig)

## Information

Configuration library which only requires:
* A pointer to a struct with exported fields which your configuration fields.
* The name of your project.

After passing these two parameters, the passed struct will be filled with the configuration values found in "~/.projectName.yml".

If that file does not exist, it will be created.

If any of the configuration variables do not exist in the yml file, the variables will be requested via [AlecAivazis/survey](https://github.com/AlecAivazis/survey) prompts.

Configuration variables will also be read from environment variables. Environment variables will take preference over .yml variables.

If any matching environment variables are found, the .yml file will not be written to or changed at all, other than to create it if it does not exist.

## Example

```
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
```



## Comments

The purpose of this library is to make configuration easy.

I wanted a library that would make yml configuration easy to set up and use.

I also wanted a configuration library which would lend itself to use in a container deployment, so overriding configuration via environment variables is supported.

Dependencies are managed via [dep](https://github.com/golang/dep).

#### FYI

This is my first open source repo so any constructive criticism is welcome, I'm rather new to golang and developing in general.

Contributions are welcome as well.