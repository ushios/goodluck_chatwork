package chatwork

import (
	"fmt"
	"io/ioutil"
)

// LoadOldChat loading chat logs
func LoadOldChat(cred *Credential, roomID int64) (interface{}, error) {
	return loadOldChat(cred, roomID, 0)
}

func loadOldChat(cred *Credential, roomID, firstChatID int64) (interface{}, error) {
	path := fmt.Sprintf(
		"/gateway.php?cmd=load_old_chat&myid=%s&_v=1.80a&_av=4&_t=%s&ln=ja&room_id=%d&last_chat_id=0&first_chat_id=%d&jump_to_chat_id=0&unread_num=0&file=1&desc=1",
		cred.MyID,
		cred.AccessToken,
		roomID,
		firstChatID,
	)
	rawResp, err := client().Get(u(path))
	if err != nil {
		return nil, err
	}
	defer rawResp.Body.Close()

	d, _ := ioutil.ReadAll(rawResp.Body)
	fmt.Println(string(d))

	return nil, nil
}
