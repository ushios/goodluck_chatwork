package chatwork

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strconv"
	"time"
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

	c, err := createContacts(cred, &result)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func createContacts(cred *Credential, res *InitLoadResult) (*Contacts, error) {
	cs := Contacts{
		ContactList: []Contact{},
		RoomList:    []Room{},
	}

	cMap := res.ContactDat
	for k, con := range cMap {
		c := Contact{
			ID:   k,
			AID:  con.AID,
			Name: con.Name,
		}
		cs.ContactList = append(cs.ContactList, c)
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

		var name string
		nameTemplate := "%s"
		switch rm.TP {
		case 1:
			name = rm.Name
		case 2:
			for key := range rm.M {
				if key != cred.MyID {
					if con, ok := cMap[key]; !ok {
						name = fmt.Sprintf(nameTemplate, "unknown user")
					} else {
						name = fmt.Sprintf(nameTemplate, con.Name)
					}
				}
			}
		case 3:
			name = "My Chat"
		}

		r := Room{
			ID:      k,
			AIDList: aIDList,
			Name:    name,
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

func createRow(roomID int64, chat *ChatMessage, acc *Account) ([]string, error) {
	chatID := strconv.FormatInt(chat.ID, 10)
	tm := time.Unix(int64(chat.TM), 0)
	name := acc.Name
	accID := strconv.FormatInt(acc.ID, 10)
	message := chat.Message

	download(roomID, message)

	// fmt.Println(chat.ID, tm.Format(time.RFC3339), acc.Name, acc.ID, chat.Message)
	return []string{chatID, tm.Format(time.RFC3339), name, accID, message}, nil
}

func download(roomID int64, message string) error {
	res := DownloadRegexp.FindStringSubmatch(message)
	if len(res) < 2 {
		return nil
	}

	fID, err := strconv.ParseInt(res[1], 10, 64)
	if err != nil {
		return err
	}

	if err := DownloadFile(fID, downloadDirname(roomID)); err != nil {
		return err
	}

	return nil
}

func downloadDirname(roomID int64) string {
	dir := fmt.Sprintf("%s/%d/%s",
		LogRootDirectoryName,
		roomID,
		AttachementDirectoryName,
	)

	return dir
}
