package chatwork

import (
	"fmt"
	"io/ioutil"
)

// InitLoad loading contact info
func InitLoad(cred *Credential) (*Contacts, error) {
	path := fmt.Sprintf(
		"/gateway.php?cmd=init_load&myid=%s&_v=1.80a&_av=4&_t=%s&ln=ja&rid=0&type=&new=1",
		cred.MyID,
		cred.AccessToken,
	)
	resp, err := client().Get(u(path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// TODO: response to json
	d, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(d))

	return nil, nil
}
