package chatwork

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"sort"
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
