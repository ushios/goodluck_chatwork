package chatwork

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"time"
)

const (
	// AttachementDirectoryName file dir.
	AttachementDirectoryName = "./attachements"
	// LogRootDirectoryName log root
	LogRootDirectoryName = "./chatwork_log"
)

var (
	// Retry setting
	Retry = 5
	// ChatLength max
	ChatLength = 40
	// FilenameRegexp from header
	FilenameRegexp *regexp.Regexp
	// DownloadRegexp from chat message
	DownloadRegexp *regexp.Regexp
	// ErrFilenameNotFound filename not found in header
	ErrFilenameNotFound = errors.New("ErrFilenameNotFound")
)

func init() {
	FilenameRegexp = regexp.MustCompile(`filename\*=UTF-8''(.+)`)
	DownloadRegexp = regexp.MustCompile(`\[download:(\d+)\].+\[\/download\]`)
}

// LoadAndSaveAllChat .
func LoadAndSaveAllChat(cred *Credential, contacts *Contacts, roomID int64, interval time.Duration) error {
	chatID := int64(0)

	accounts, err := AccountInfo(cred, contacts)
	if err != nil {
		return err
	}

	// get csv file handler
	room, ok := contacts.RoomMap[roomID]
	if !ok {
		return fmt.Errorf("room (%d) not found", roomID)
	}

	dirname := fmt.Sprintf("%s(%s)", room.Name, room.ID)
	dirpath := filepath.Join(LogRootDirectoryName, dirname)
	if err := checkDir(dirpath); err != nil {
		return err
	}
	filename := filepath.Join(dirpath, "messages.csv")
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	// converter := transform.NewWriter(file, japanese.ShiftJIS.NewEncoder())
	writer := csv.NewWriter(file)

	// create csv header
	if err = file.Truncate(0); err != nil {
		return err
	}
	writer.Write([]string{"chat_id", "time", "name", "account_id", "message"})

	// logging
	for {
		log.Printf("Downloading %d - %d messages \n", roomID, chatID)
		res, err := LoadOldChat(cred, roomID, chatID)
		if err != nil {
			return err
		}

		// output body
		for _, chat := range res.ChatList {
			acc, ok := (*accounts)[chat.AID]
			if !ok {
				acc = Account{
					ID:   chat.AID,
					Name: "deleted user",
				}
			}

			// Buffer to csv writer
			row, err := createRow(roomID, &chat, &acc)
			if err != nil {
				return err
			}
			if err := writer.Write(row); err != nil {
				return err
			}

			// download if file exists
			if err := download(roomID, chat.Message, dirpath); err != nil {
				return err
			}
		}

		if len(res.ChatList) < ChatLength {
			break
		}

		time.Sleep(interval)
		chatID = res.ChatList[len(res.ChatList)-1].ID
	}

	// Flush
	writer.Flush()

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

	// sort by ID
	sort.Sort(sort.Reverse(result))

	return &result, nil
}

// DownloadFile get file info
func DownloadFile(fID int64, dirpath string) error {
	path := fmt.Sprintf(
		"/gateway.php?cmd=download_file&bin=1&file_id=%d",
		fID,
	)
	rawResp, err := client().Get(u(path))
	if err != nil {
		return err
	}
	defer rawResp.Body.Close()

	// create filename
	filename, err := filenameFromResponse(rawResp)
	if err != nil {
		switch err {
		case ErrFilenameNotFound:
			filename = fmt.Sprintf("unknown-file-%d", fID)
		default:
			return err
		}
	}

	// Download file
	d, err := ioutil.ReadAll(rawResp.Body)
	if err != nil {
		return err
	}

	if err := checkDir(dirpath); err != nil {
		return err
	}

	dest := filepath.Join(dirpath, filename)
	if err := ioutil.WriteFile(dest, d, 0644); err != nil {
		return err
	}

	return nil
}

func filenameFromResponse(resp *http.Response) (string, error) {
	cd := resp.Header.Get("Content-disposition")
	if cd == "" {
		return "", fmt.Errorf("Content-disposition was empty")
	}
	res := FilenameRegexp.FindStringSubmatch(cd)
	if len(res) < 2 {
		return "", ErrFilenameNotFound
	}

	filename, err := url.QueryUnescape(res[1])
	if err != nil {
		return "", err
	}

	return filename, nil
}
