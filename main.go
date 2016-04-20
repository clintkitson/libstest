package main

import (
	"bytes"
	"fmt"

	"github.com/akutz/gofig"
	apiserver "github.com/emccode/libstorage/api/server"
	"gopkg.in/yaml.v2"
)

func main() {
	driver := "mock"

	libstorageConfigMap := map[string]interface{}{
		"driver": driver,
		"server": map[string]interface{}{
			"services": map[string]interface{}{
				driver: nil,
			},
		},
	}

	libstorageConfig := map[string]interface{}{
		"libstorage": libstorageConfigMap,
	}

	yamlBuf, err := yaml.Marshal(libstorageConfig)
	if err != nil {
		panic(err)
	}

	config := gofig.New()

	if err := config.ReadConfig(bytes.NewReader(yamlBuf)); err != nil {
		panic(err)
	}

	server, err := apiserver.Serve(config)
	if err != nil {
		panic(err)
	}

	fmt.Println(server)
}
