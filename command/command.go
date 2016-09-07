package command

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
)

func init() {

}

// LoadCredential load from file
func LoadCredential(path string) (*LoginInfo, error) {
	var info *LoginInfo

	js, err := ioutil.ReadFile(path)
	if err != nil {
		return info, err
	}

	if err := json.Unmarshal(js, info); err != nil {
		return info, err
	}

	return info, nil
}

// SaveCredential save to file
func SaveCredential(path string, li *LoginInfo) error {
	js, err := json.Marshal(li)
	if err != nil {
		return err
	}

	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	writer := bufio.NewWriter(f)
	if _, err := writer.Write(js); err != nil {
		return err
	}
	writer.Flush()

	return nil
}
