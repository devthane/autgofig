package gofig

import (
	"runtime"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
	"gopkg.in/AlecAivazis/survey.v1"
	"os/exec"
)

type Loader struct {
	ConfLocation string
	Separator    string
	Config       *Config
	fields       []string
	name         string
	HomeDir      string
}

func NewLoader(fields []string, name string) *Loader {
	var l Loader

	l.fields = fields
	l.name = name

	if runtime.GOOS == "windows" {
		l.Separator = "\\"
		l.HomeDir = os.Getenv("userprofile")
		l.ConfLocation = l.HomeDir + l.Separator + l.name + ".yml"
	} else {
		l.Separator = "/"
		l.HomeDir = os.Getenv("HOME")
		l.ConfLocation = l.HomeDir + l.Separator + l.name + ".yml"
	}

	l.loadConfig()

	return &l
}

func (l *Loader) loadConfig() {
	data := make([]byte, 2048)
	data, err := ioutil.ReadFile(l.ConfLocation)
	if err != nil {
		f, err := os.Create(l.ConfLocation)
		if err != nil {
			log.Panicf("ERR: could not read file: %s", l.ConfLocation)
		}
		defer f.Close()
		f.Read(data)
	}

	c := newConfig(l.fields)
	err = yaml.Unmarshal(data, c)
	l.configure(c)
}

func (l *Loader) configure(c *Config) {
	if c == nil {
		c = newConfig(l.fields)
	}

	mutatingRawConfig := *c

	clear()
	for field, value := range mutatingRawConfig {
		if value == "" {
			prompt := &survey.Input{
				Message: field + ":",
			}
			result := ""

			survey.AskOne(prompt, &result, nil)

			mutatingRawConfig[field] = result
		}
	}

	l.writeRawConfig(mutatingRawConfig)
	l.Config = c
}

func (l *Loader) writeRawConfig(m map[string]string) {
	data, err := yaml.Marshal(m)
	if err != nil {
		log.Fatalln("ERR: Could not prepare Config for write.", err)
	}

	file, err := os.Create(l.ConfLocation)
	if err != nil {
		log.Fatalln("ERR: could not open Config file for writing.", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		log.Fatalln("ERR: Could not write to Config file.", err)
	}
}

func clear() {
	cmd := new(exec.Cmd)
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
	cmd.Wait()
}
