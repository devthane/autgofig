package gofig_test

import (
	"github.com/thane421/gofig"
	"testing"
	"os"
)

func TestIntegrationYaml(t *testing.T) {
	var (
		var1  string
		var2  string
		var3  string
		var4  string
		var5  string
		var6  string
		var7  string
		var8  string
		var9  string
		var10 string
	)

	l := gofig.NewLoader([]string{
		"var1",
		"var2",
		"var3",
		"var4",
		"var5",
		"var6",
		"var7",
		"var8",
		"var9",
		"var10",
	}, "config-test")

	config := *l.Config
	defer os.Remove(l.ConfLocation)

	var1 = config["var1"]
	var2 = config["var2"]
	var3 = config["var3"]
	var4 = config["var4"]
	var5 = config["var5"]
	var6 = config["var6"]
	var7 = config["var7"]
	var8 = config["var8"]
	var9 = config["var9"]
	var10 = config["var10"]

	if var1 == "" {
		t.Fail()
	}
	if var2 == "" {
		t.Fail()
	}
	if var3 == "" {
		t.Fail()
	}
	if var4 == "" {
		t.Fail()
	}
	if var5 == "" {
		t.Fail()
	}
	if var6 == "" {
		t.Fail()
	}
	if var7 == "" {
		t.Fail()
	}
	if var8 == "" {
		t.Fail()
	}
	if var9 == "" {
		t.Fail()
	}
	if var10 == "" {
		t.Fail()
	}

}
