package chatwork

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type (
	// AccountInfoRequest .
	AccountInfoRequest struct {
		AID            []int64 `json:"aid"`
		GetPrivateData int     `json:"get_private_data"`
	}

	// Account .
	Account struct {
		ID   int64
		Name string
	}
)

// AccountInfo .
func AccountInfo(cred *Credential, c *Contacts) (*map[int64]Account, error) {
	postStruct := AccountInfoRequest{
		AID:            c.AIDs(),
		GetPrivateData: 0,
	}

	postJSON, err := json.Marshal(postStruct)
	if err != nil {
		return nil, err
	}

	postData := string(postJSON)
	path := fmt.Sprintf("/gateway.php?cmd=get_account_info&myid=%s&_v=1.80a&_av=4&_t=%s&ln=ja",
		cred.MyID,
		cred.AccessToken,
	)

	values := url.Values{}
	values.Add("pdata", postData)

	rawResp, err := client().PostForm(u(path), values)
	if err != nil {
		return nil, err
	}
	defer rawResp.Body.Close()

	result := AccountResult{}
	if _, err := ReadResponse(rawResp, &result); err != nil {
		return nil, err
	}

	m := map[int64]Account{}
	for _, acc := range result.AccountDat {
		m[acc.AID] = Account{
			ID:   acc.AID,
			Name: acc.Name,
		}
	}

	return &m, nil
}
