package main
import (
//        "encoding/json"
        "fmt"
//        "time"
	log "github.com/sirupsen/logrus"
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

func GetFileData(filename string) ([]byte, error) {
        FileData, err := os.Open(filename)
        if err != nil {
                log.Errorf("Error!! Sorry could not locate file at location: %v", err)
		return nil, err
        }
        fmt.Println("Successfully Opened " + filename)
        defer FileData.Close()
        byteValue, _ := ioutil.ReadAll(FileData)
        return byteValue, nil
}
