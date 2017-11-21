package gofig

type Config map[string]string

func newConfig(fields []string) *Config {
	c := make(Config)

	for _, field := range fields {
		c[field] = ""
	}

	return &c
}
