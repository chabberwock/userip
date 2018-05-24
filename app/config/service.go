package config

import (
	"os"
	"encoding/json"
	"fmt"
)

type IpStackConfig struct {
	RequestLimit int
	AccessKey string
	Enabled bool
}

type NekudoConfig struct {
	RequestLimit int
	Enabled bool
}


type Config struct {
	IpStack IpStackConfig
	Nekudo NekudoConfig
	TemplatePath string
	Bind string
	CacheTTL int64
}

func Read(filename string) Config {
	file, _ := os.Open(filename)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}