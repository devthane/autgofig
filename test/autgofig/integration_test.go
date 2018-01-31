package autgofig

import (
	"github.com/devthane/autgofig/pkg/autgofig"
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
		Separator    string
		HomeDir      string
		ConfLocation string
	)

	name := "config-yaml-test"

	if runtime.GOOS == "windows" {
		Separator = "\\"
		HomeDir = os.Getenv("userprofile")
	} else {
		Separator = "/"
		HomeDir = os.Getenv("HOME")
	}

	ConfLocation = HomeDir + Separator + "." + name + ".yml"

	file, err := os.Create(ConfLocation)
	if err != nil {
		t.Fatal(err)
	}

	file.Write([]byte(
		fmt.Sprintf("StringField: YamlTest\nIntField: \"5\"\n"),
	))
	file.Close()
	defer os.Remove(ConfLocation)

	config := new(TestObj)
	autgofig.LoadConfig(config, name)

	if config.StringField != "YamlTest" || config.IntField != 5 {
		t.Fail()
	}
}

func TestIntegrationEnv(t *testing.T) {
	os.Setenv("StringField", "EnvTest")
	os.Setenv("IntField", "10")
	defer os.Clearenv()

	config := new(TestObj)
	autgofig.LoadConfig(config, "config-env-test")

	if config.StringField != "EnvTest" || config.IntField != 10 {
		t.Fail()
	}
}

func TestIntegrationPriority(t *testing.T) {
	var (
		Separator    string
		HomeDir      string
		ConfLocation string
	)

	os.Setenv("StringField", "PriorityTest")
	defer os.Clearenv()

	name := "config-priority-test"

	if runtime.GOOS == "windows" {
		Separator = "\\"
		HomeDir = os.Getenv("userprofile")
	} else {
		Separator = "/"
		HomeDir = os.Getenv("HOME")
	}

	ConfLocation = HomeDir + Separator + "." + name + ".yml"

	file, err := os.Create(ConfLocation)
	if err != nil {
		t.Fatal(err)
	}

	file.Write([]byte(
		fmt.Sprintf("StringField: YamlTest\nIntField: \"5\"\n"),
	))
	file.Close()
	defer os.Remove(ConfLocation)

	config := new(TestObj)
	autgofig.LoadConfig(config, name)

	if config.StringField != "PriorityTest" || config.IntField != 5 {
		t.Fail()
	}
}
