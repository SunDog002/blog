package config

import "fmt"

type System struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Env       string `yaml:"env"`
	Directory string `yaml:"directory"`
	Router    string `yaml:"router"`
}

func (s System) Address() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
