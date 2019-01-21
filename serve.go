package main
import (
	"encoding/json"
	"fmt"
	"time"
//	"io/ioutil"
//	"os"
)
func rebuildVersionScheduler(filename string) error {
	fmt.Println("Debug")
	byteValue, err := GetFileData(filename)
		if err != nil {
			fmt.Errorf("data not found %v", err)
			return err
		}
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
	return nil
}
func scheduler() error {
	filename := "./config.json"
	pollInterval := 300
	tiktok := time.Tick(time.Duration(pollInterval) * time.Millisecond)
	for range tiktok {
		err := rebuildVersionScheduler(filename)
			if err != nil {
				return err
			}
	}
	return nil
}


func main() {
	go scheduler()
	fmt.Println("this is thread from Main")
	fmt.Scanln()
	fmt.Println("done")
}
