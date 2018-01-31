package autgofig

import (
	"testing"
	"reflect"
)

type TestObj struct {
	StringField string
	IntField    int
}

func TestUnpackFields(t *testing.T) {
	l := new(Loader)

	l.unpackFields(new(TestObj))
	if l.fields["StringField"] != reflect.String {
		t.Fail()
	}
	if l.fields["IntField"] != reflect.Int {
		t.Fail()
	}
}

func TestConfigure(t *testing.T) {
	l := new(Loader)

	testObj := new(TestObj)
	l.Config = testObj
	l.unpackFields(testObj)

	presetConfig := make(map[string]string)
	presetConfig["IntField"] = "5"
	presetConfig["StringField"] = "test"

	l.configure(presetConfig)

	config, ok := l.Config.(*TestObj)
	if !ok {
		t.Fail()
	}
	if config.IntField != 5 || config.StringField != "test" {
		t.Fail()
	}
}
