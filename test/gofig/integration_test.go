package gofig

import (
	"github.com/thane421/gofig/pkg/gofig"
	"testing"
	"os"
	"runtime"
	"fmt"
)

type TestObj struct {
	StringField string
	IntField    int
}

func TestIntegrationYaml(t *testing.T) {
	var (
		stringField  string
		intField     int
		Separator    string
		HomeDir      string
		ConfLocation string
	)

	name := ".config-test"

	if runtime.GOOS == "windows" {
		Separator = "\\"
		HomeDir = os.Getenv("userprofile")
		ConfLocation = HomeDir + Separator + name + ".yml"
	} else {
		Separator = "/"
		HomeDir = os.Getenv("HOME")
		ConfLocation = HomeDir + Separator + name + ".yml"
	}

	file, err := os.Create(ConfLocation)
	if err != nil {
		t.Fatal(err)
	}

	file.Write([]byte(
		fmt.Sprintf("StringField: test\nIntField: \"5\"\n"),
	))
	file.Close()
	defer os.Remove(ConfLocation)

	config := new(TestObj)
	gofig.LoadConfig(config, name)

	stringField = config.StringField
	intField = config.IntField

	if stringField != "test" || intField != 5 {
		t.Fail()
	}

}
