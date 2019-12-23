package team

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Value struct {
	Member []string `yaml:"member"`
}

type Team map[string]Value

func Load(file string) (*Team, error){
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var d Team
	if err := yaml.Unmarshal(buf, &d); err != nil {
		return nil, err
	}
	return &d, err
}