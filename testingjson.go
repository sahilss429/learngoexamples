package main
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)
type Configs []struct {
	Path string   `json:"path"`
	DC   []string `json:"DC"`
}
func (c Configs) toString() string {
	bytes, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(bytes)
}

func getConfigs() []Configs {
	configs := make([]Configs, 1)
	raw, err := ioutil.ReadFile("sample1.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &configs)
	return configs
}


func main() {
	configs := getConfigs()
	fmt.Println(configs)
	for _, value := range configs {
		fmt.Println(value.toString())
	}
}
