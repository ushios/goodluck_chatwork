package chatwork

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

var (
	// Retry setting
	Retry = 5
	// ChatLength max
	ChatLength = 40
	// FilenameRegExp from header
	FilenameRegExp *regexp.Regexp
)

func init() {
	FilenameRegExp = regexp.MustCompile(`filename\*=UTF-8''(.+)`)
}

// FileInfo .
type FileInfo struct {
	Filename string
	URL      *url.URL
}

// LoadAndSaveAllChat .
func LoadAndSaveAllChat(cred *Credential, roomID int64, interval time.Duration) error {
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

		time.Sleep(interval)
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

// DownloadFileInfo get file info
func DownloadFileInfo(fID int64) (*FileInfo, error) {
	fi := FileInfo{}

	path := fmt.Sprintf(
		"/gateway.php?cmd=download_file&bin=1&file_id=%d",
		fID,
	)
	rawResp, err := client().Get(u(path))
	if err != nil {
		return nil, err
	}
	defer rawResp.Body.Close()

	// // create url
	// fi.URL, err = rawResp.Location()
	// if err != nil {
	// 	return nil, err
	// }

	// create filename
	fi.Filename, err = filename(rawResp)
	if err != nil {
		return nil, err
	}

	return &fi, nil
}

func filename(resp *http.Response) (string, error) {
	cd := resp.Header.Get("Content-disposition")
	if cd == "" {
		return "", fmt.Errorf("Content-disposition was empty")
	}
	res := FilenameRegExp.FindStringSubmatch(cd)
	if len(res) < 2 {
		return "", fmt.Errorf("filename not found")
	}

	filename, err := url.QueryUnescape(res[1])
	if err != nil {
		return "", err
	}

	return filename, nil
}
