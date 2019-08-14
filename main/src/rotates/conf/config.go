package conf

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Service Service
	Alerts  Alert
	Checks  Check
}

type Service struct {
	Port            int
	Interval        int
	SuppressionTime int `toml:"suppression_time"`
	AlertTimes      int `toml:"alert_times"`
}

type Alert struct {
	Dings []Ding
	Alis  []Ali
}

type Ding struct {
	Name  string
	Token string
	AtAll bool `toml:"at_all"`
	At    []int
}

type Ali struct {
	Name          string
	Key           string
	Secret        string
	ShowNumber    string   `toml:"show_number"`
	CalledNumbers []string `toml:"called_numbers"`
	TtsCode       string   `toml:"tts_code"`
	Params        map[string]string
}

type Check struct {
	Urls  []Url
	Ports []Port
}

type Url struct {
	Name   string
	Url    string
	Expect string
}

type Port struct {
	Name string
	Host string
	Port int
}

func ParseConfig(file string) Config {
	conf := Config{}
	_, err := toml.DecodeFile(file, &conf)
	if err != nil {
		log.Fatal(err)
	}
	return conf
}
