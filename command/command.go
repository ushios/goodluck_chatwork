package command

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/ushios/goodluck_chatwork/lib/chatwork"
)

const (
	// LogRootDirectoryName log root
	LogRootDirectoryName = "./chatwork_log"
	// AttachementDirectoryName file dir.
	AttachementDirectoryName = "./attachements"
)

func init() {

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

func createRow(roomID int64, chat *chatwork.ChatMessage, acc *chatwork.Account) ([]string, error) {
	chatID := strconv.FormatInt(chat.ID, 10)
	tm := time.Unix(int64(chat.TM), 0)
	name := acc.Name
	accID := strconv.FormatInt(acc.ID, 10)
	message := chat.Message

	// fmt.Println(chat.ID, tm.Format(time.RFC3339), acc.Name, acc.ID, chat.Message)
	return []string{chatID, tm.Format(time.RFC3339), name, accID, message}, nil
}

func download(roomID int64, message string, parentDirName string) error {
	fID, err := chatwork.FileIDFromMessage(message)
	if err != nil {
		switch err {
		case chatwork.ErrFileIDNotFound:
			return nil
		default:
			return err
		}
	}

	log.Printf("file(%d) downloading... \n", fID)
	if err := chatwork.DownloadFile(fID, downloadDirname(parentDirName)); err != nil {
		return err
	}

	return nil
}

func downloadDirname(parentDirName string) string {
	dir := filepath.Join(
		parentDirName,
		AttachementDirectoryName,
	)

	return dir
}
