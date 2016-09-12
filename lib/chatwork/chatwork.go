package chatwork

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
)

// InitLoad loading contact info
func InitLoad(cred *Credential) (*Contacts, error) {
	path := fmt.Sprintf(
		"/gateway.php?cmd=init_load&myid=%s&_v=1.80a&_av=4&_t=%s&ln=ja&rid=0&type=&new=1",
		cred.MyID,
		cred.AccessToken,
	)
	rawResp, err := client().Get(u(path))
	if err != nil {
		return nil, err
	}
	defer rawResp.Body.Close()

	result := InitLoadResult{}
	_, err = ReadResponse(rawResp, &result)
	if err != nil {
		return nil, err
	}

	c := createContacts(&result)

	return c, nil
}

func createContacts(res *InitLoadResult) *Contacts {
	cs := Contacts{
		ContactList: []Contact{},
		RoomList:    []Room{},
	}

	cMap := res.ContactDat
	for k, con := range cMap {
		if !con.IsDeleted {
			c := Contact{
				ID:   k,
				Name: con.Name,
			}

			cs.ContactList = append(cs.ContactList, c)
		}
	}

	rMap := res.RoomDat
	for k, rm := range rMap {
		r := Room{
			ID:   k,
			Name: rm.Name,
		}

		cs.RoomList = append(cs.RoomList, r)
	}

	return &cs
}

func u(path string) string {
	return fmt.Sprintf("%s%s", BaseURL, path)
}

func client() *http.Client {
	if c != nil {
		return c
	}

	UsedJar, err := cookiejar.New(nil)
	if err != nil {
		return nil
	}
	c = &http.Client{
		Jar: UsedJar,
	}

	return c
}

func checkDir(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
		return nil
	}

	return err
}
