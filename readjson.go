package main
import (
	"encoding/json"
	"fmt"
	"time"
	"io/ioutil"
	"os"
)
type Configs struct {
	Configs []Config `json:"configs"`
}

type Config struct {
	Path   string   `json:"path"`
	Dc     []string `json:"dc"`
}

func getFileData(filename string) []byte {
	FileData, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened " + filename)
	defer FileData.Close()
	byteValue, _ := ioutil.ReadAll(FileData)
	return byteValue
}

func rebuildVersionScheduler(filename string){
	fmt.Println("Debug")
	byteValue := getFileData(filename)
	var configs Configs
	json.Unmarshal(byteValue, &configs)
	for i := 0; i < len(configs.Configs); i++ {
                for j := 0; j < len(configs.Configs[i].Dc); j++ {
                path := configs.Configs[i].Path
                dc := configs.Configs[i].Dc[j]
                fmt.Println("run rebuild version: "+ path  + " " + dc)
/*
		config := &api.Config {
			Address: "localhost:8500",
			Scheme: "http",
			Datacenter: dc,
		}
		consulClient, err := api.NewClient(config)
		if err != nil {
			log.Error(err)
		}
		port := strconv.Itoa(webConfig.ClientVersionsPort)
		_, err = kv.RebuildVersions(consulClient, path, port)
		if err != nil {
			log.Errorf("Error rebuilding versions for %s in %s: %v", path, dc, err)
		}
                }
		return
*/
		}
	}
	return
}
func main() {
	filename := "./config.json"
	pollInterval := 300
	tiktok := time.Tick(time.Duration(pollInterval) * time.Millisecond)
	for range tiktok {
		go rebuildVersionScheduler(filename)
	}
	fmt.Println("this is thread from Main")
	fmt.Scanln()
	fmt.Println("done")
}
