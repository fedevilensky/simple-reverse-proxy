package main

type SimpleProxy struct {
	Port   int     `yaml:"port"`
	Routes []Route `yaml:"routes"`
}

type Route struct {
	Name      string  `yaml:"omitempty"`
	Subdomain *string `yang:"subdomain"`
	To        string  `yaml:"to"`
}
