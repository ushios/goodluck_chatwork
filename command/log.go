package command

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/cli"
	"github.com/ushios/goodluck_chatwork/lib/chatwork"
)

var (
	// DefaultInterval wait second
	DefaultInterval = 1 * time.Second
)

// CmdLog show chat log
func CmdLog(c *cli.Context) error {
	email := c.String("email")
	password := c.String("password")

	if email == "" || password == "" {
		return fmt.Errorf("empty email or password")
	}

	cred, err := chatwork.Login(email, password)
	if err != nil {
		fmt.Println(err)
		return err
	}

	contacts, err := chatwork.InitLoad(cred)
	if err != nil {
		fmt.Println(err)
		return err
	}

	roomID := c.Int("room")
	err = LoadAndSaveAllChat(cred, contacts, int64(roomID), DefaultInterval)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// LoadAndSaveAllChat .
func LoadAndSaveAllChat(cred *chatwork.Credential, contacts *chatwork.Contacts, roomID int64, interval time.Duration) error {
	chatID := int64(0)

	accounts, err := chatwork.AccountInfo(cred, contacts)
	if err != nil {
		return err
	}

	// get csv file handler
	room, ok := contacts.RoomMap[roomID]
	if !ok {
		return fmt.Errorf("room (%d) not found", roomID)
	}

	dirname := fmt.Sprintf("%s(%s)", room.Name, room.ID)
	dirpath := filepath.Join(
		LogRootDirectoryName,
		dirname,
	)
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
		res, err := chatwork.LoadOldChat(cred, roomID, chatID)
		if err != nil {
			return err
		}

		// output body
		for _, chat := range res.ChatList {
			acc, ok := (*accounts)[chat.AID]
			if !ok {
				acc = chatwork.Account{
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

		if len(res.ChatList) < chatwork.ChatLength {
			break
		}

		time.Sleep(interval)
		chatID = res.ChatList[len(res.ChatList)-1].ID
	}

	// Flush
	writer.Flush()

	return nil
}
