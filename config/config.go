package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// template struct
type template struct {
	Type       string
	RemoteName string
	LocalName  string
	Port       string
}

// Docker struct
type docker struct {
	GhKey          string
	SshPort        string
	SshHost        string
	SshUsername    string
	DockerPassword string
	DockerUsername string
	SshKey         string
}

type config struct {
	Template template
	Docker   docker
}

func Configuration() (map[string]config, error) {

	yfile, err := ioutil.ReadFile("atomatiki.yaml")
	if err != nil {
		return nil, err
	}

	data := make(map[string]config)

	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		return nil, err
	}

	return data, nil

}
