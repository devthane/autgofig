package gofig

import (
	"runtime"
	"os"
	"gopkg.in/yaml.v2"
	"log"
	"os/exec"
	"reflect"
	"io/ioutil"
	"gopkg.in/AlecAivazis/survey.v1"
	"strconv"
	"fmt"
)

type Loader struct {
	ConfLocation string
	Separator    string
	Config       interface{}
	fields       map[string]reflect.Kind
	name         string
	HomeDir      string
}

// configOut must be a pointer to a struct containing the configuration fields, which must be exported fields.
// the struct passed will receive the values either given by prompts, or found in the .yml file.
// LoadConfig will set up a Loader, which will not be returned. Use NewLoader if you would like access to the Loader.
// The yml file's name is determined by the projectName parameter, e.g. "~/" + "." + projectName + ".yml"
// if any of the fields exist on the system as environment variables, they will be loaded as such over the yml config.
func LoadConfig(configOut interface{}, projectName string) error {
	_, err := NewLoader(configOut, projectName)
	return err
}

// Same as LoadConfig, except also returns the Loader, which contains useful OS specific variables.
// Such as directory separators and the users home directory.
func NewLoader(configOut interface{}, projectName string) (*Loader, error) {
	var l Loader

	if reflect.TypeOf(configOut).Kind() != reflect.Ptr || reflect.ValueOf(configOut).Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("the configOut paramater must be pointer to a struct")
	}

	l.name = projectName
	l.unpackFields(configOut)
	l.Config = configOut

	if runtime.GOOS == "windows" {
		l.Separator = "\\"
		l.HomeDir = os.Getenv("userprofile")

	} else {
		l.Separator = "/"
		l.HomeDir = os.Getenv("HOME")
	}

	l.ConfLocation = l.HomeDir + l.Separator + "." + l.name + ".yml"
	l.loadConfig()

	return &l, nil
}

func (l *Loader) unpackFields(fields interface{}) {
	var results = make(map[string]reflect.Kind)

	fieldsValue := reflect.ValueOf(fields).Elem()
	fieldsType := fieldsValue.Type()

	for i := 0; i < fieldsValue.NumField(); i++ {
		results[fieldsType.Field(i).Name] = fieldsType.Field(i).Type.Kind()
	}

	l.fields = results
}

func (l *Loader) loadConfig() error {
	data := make([]byte, 2048)
	data, err := ioutil.ReadFile(l.ConfLocation)
	if err != nil {
		f, err := os.Create(l.ConfLocation)
		if err != nil {
			log.Panicf("ERR: could not read file: %s", l.ConfLocation)
		}
		f.Read(data)
		f.Close()
	}

	configFromEnv := l.loadFromEnv()
	configFromFile := make(map[string]string)
	if err = yaml.Unmarshal(data, configFromFile); err != nil {
		return err
	}

	config := l.mergeConfigs(configFromEnv, configFromFile)
	if err = l.configure(config); err != nil {
		return err
	}


	return nil
}

func (l *Loader) mergeConfigs(primary map[string]string, secondary map[string]string) map[string]string {
	results := make(map[string]string)
	for fieldName := range l.fields {
		if value, exists := primary[fieldName]; exists && value != "" {
			results[fieldName] = value
		} else if value, exists := secondary[fieldName]; exists && value != "" {
			results[fieldName] = value
		}
	}

	return results
}

func (l *Loader) loadFromEnv() map[string]string {
	results := make(map[string]string)
	for fieldName := range l.fields {
		results[fieldName] = os.Getenv(fieldName)
	}

	return results
}

func (l *Loader) configure(config map[string]string) error {
	var (
		value  interface{}
		err    error
		exists bool
	)

	clear()
	reflectedConfig := reflect.ValueOf(l.Config).Elem()
	reflectedType := reflectedConfig.Type()

	for i := 0; i < reflectedConfig.NumField(); i++ {
		fieldKind := reflectedConfig.Field(i).Kind()
		fieldName := reflectedType.Field(i).Name

		if fieldKind == reflect.Int {
			value, exists = config[fieldName]
			if !exists {
				value = 0
			} else if value, err = strconv.Atoi(config[fieldName]); err != nil {
				return err
			}
		} else if fieldKind == reflect.String {
			value = config[fieldName]
		} else {
			return fmt.Errorf("type of field: %s must be string or int, not %s", fieldName, fieldKind)
		}

		if value == 0 || value == "" {
			prompt := &survey.Input{
				Message: fieldName + ":",
			}
			result := ""

			survey.AskOne(prompt, &result, nil)

			if fieldKind == reflect.Int {
				value, err = strconv.Atoi(result)
				if err != nil {
					return err
				}
			} else {
				value = result
			}
		}

		if fieldKind == reflect.Int {
			reflectedConfig.Field(i).SetInt(int64(value.(int)))
		} else {
			reflectedConfig.Field(i).SetString(value.(string))
		}
	}

	return nil
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
