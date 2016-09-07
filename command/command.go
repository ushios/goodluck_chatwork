package command

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ushios/goodluck_chatwork/lib/chatwork"
)

func init() {

}

// LoadCredential load from file
func LoadCredential(path string) (*chatwork.Credential, error) {
	var cred *chatwork.Credential

	js, err := ioutil.ReadFile(path)
	if err != nil {
		return cred, err
	}

	if err := json.Unmarshal(js, cred); err != nil {
		return cred, err
	}

	return cred, nil
}

// SaveCredential save to file
func SaveCredential(path string, cred *chatwork.Credential) error {
	js, err := json.Marshal(cred)
	if err != nil {
		return err
	}

	fmt.Println(string(js))

	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	writer := bufio.NewWriter(f)
	if _, err := writer.Write(js); err != nil {
		return err
	}
	writer.Flush()

	return nil
}
