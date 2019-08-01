package main

//
// Simple struct for use as a config dictionary.
//
type Config struct {
	items map[string]interface{}
}

func (c *Config) Set(k string, v interface{}) {
	c.items[k] = v
}

func (c Config) Get(k string) interface{} {
	return c.items[k]
}
