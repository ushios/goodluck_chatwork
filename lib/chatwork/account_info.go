package chatwork

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type (
	// AccountInfoRequest .
	AccountInfoRequest struct {
		AID            []int64 `json:"aid"`
		GetPrivateData int     `json:"get_private_data"`
	}

	// AccountInfoResponse .
	AccountInfoResponse struct {
	}
)

// AccountInfo .
func AccountInfo(cred *Credential, c *Contacts) (*[]AccountInfoResponse, error) {
	postStruct := AccountInfoRequest{
		AID:            c.AIDs(),
		GetPrivateData: 0,
	}

	postJSON, err := json.Marshal(postStruct)
	if err != nil {
		return nil, err
	}

	postData := string(postJSON)
	fmt.Println(postData)
	path := fmt.Sprintf("/gateway.php?cmd=get_account_info&myid=%s&_v=1.80a&_av=4&_t=%s&ln=ja",
		cred.MyID,
		cred.AccessToken,
	)

	values := url.Values{}
	values.Add("pdata", postData)

	resp, err := client().PostForm(u(path), values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	d, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(d))

	return nil, nil
}
