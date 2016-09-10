package chatwork

import (
	"fmt"
	"time"
)

var (
	// Retry setting
	Retry = 5
	// ChatLength max
	ChatLength = 40
)

// LoadAndSaveAllChat .
func LoadAndSaveAllChat(cred *Credential, roomID int64) error {
	chatID := int64(0)

	for {
		res, err := LoadOldChat(cred, roomID, chatID)
		if err != nil {
			return err
		}

		for _, chat := range res.ChatList {
			tm := time.Unix(int64(chat.TM), 0)
			fmt.Println(chat.ID, tm.Format(time.RFC3339), chat.Message)
		}

		if len(res.ChatList) < ChatLength {
			break
		}

		chatID = res.ChatList[len(res.ChatList)-1].ID
	}
	return nil
}

// LoadOldChat loading chat logs
func LoadOldChat(cred *Credential, roomID, firstChatID int64) (*LoadOldChatResult, error) {
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

	result := LoadOldChatResult{}
	_, err = ReadResponse(rawResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func FileInfo(cred *Credential) {

}
