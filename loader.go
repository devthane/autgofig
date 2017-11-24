package gofig

import (
	"runtime"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
	"log"
	"bufio"
)

type Loader struct {
	confLocation string
	Separator    string
	Config       *Config
	fields       []string
	name         string
}

func NewLoader(fields []string, name string) *Loader {
	var l Loader

	l.fields = fields
	l.name = name

	if runtime.GOOS == "windows" {
		l.Separator = "\\"
		l.confLocation = os.Getenv("userprofile") + l.Separator + l.name + ".yml"
	} else {
		l.Separator = "/"
		l.confLocation = os.Getenv("HOME") + l.Separator + l.name + ".yml"
	}

	l.loadConfig()

	return &l
}

func (l *Loader) loadConfig() {
	data := make([]byte, 2048)
	data, err := ioutil.ReadFile(l.confLocation)
	if err != nil {
		f, err := os.Create(l.confLocation)
		if err != nil {
			log.Panicf("ERR: could not read file: %s", l.confLocation)
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

	Clear()
	for field, value := range mutatingRawConfig {
		if value == "" {
			mutatingRawConfig[field] = Input(field)
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

	file, err := os.Create(l.confLocation)
	if err != nil {
		log.Fatalln("ERR: could not open Config file for writing.", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		log.Fatalln("ERR: Could not write to Config file.", err)
	}
}

func (l *Loader) requestConf(field string) string {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	w.WriteString(fmt.Sprintf("%s> ", field))
	w.Flush()

	cr := byte('\n')
	answer, err := r.ReadBytes(cr)
	if err != nil {
		log.Fatalln("Could not read input:", err)
	}

	return string(answer[:len(answer)-1])
}
