package chatwork

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strconv"
)

type (
	// Contacts have contacts and room info
	Contacts struct {
		ContactList []Contact
		RoomList    []Room
	}

	// Contact data
	Contact struct {
		ID   string
		AID  int64
		Name string
	}

	// Room data
	Room struct {
		ID      string
		AIDList []int64
		Name    string
	}
)

// AIDs .
func (c *Contacts) AIDs() []int64 {
	ids := []int64{}
	m := map[int64]bool{}

	for _, con := range c.ContactList {
		m[con.AID] = true
	}

	for _, room := range c.RoomList {
		for _, aid := range room.AIDList {
			m[aid] = true
		}
	}

	for key := range m {
		ids = append(ids, key)
	}

	return ids
}

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

	c, err := createContacts(&result)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func createContacts(res *InitLoadResult) (*Contacts, error) {
	cs := Contacts{
		ContactList: []Contact{},
		RoomList:    []Room{},
	}

	cMap := res.ContactDat
	for k, con := range cMap {
		if !con.IsDeleted {
			c := Contact{
				ID:   k,
				AID:  con.AID,
				Name: con.Name,
			}

			cs.ContactList = append(cs.ContactList, c)
		}
	}

	rMap := res.RoomDat
	for k, rm := range rMap {
		aIDList := []int64{}
		for key := range rm.M {
			aID, err := strconv.ParseInt(key, 10, 64)
			if err != nil {
				return nil, err
			}
			aIDList = append(aIDList, aID)
		}

		r := Room{
			ID:      k,
			AIDList: aIDList,
			Name:    rm.Name,
		}

		cs.RoomList = append(cs.RoomList, r)
	}

	return &cs, nil
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
