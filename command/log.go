package command

import (
	"fmt"
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
	err = chatwork.LoadAndSaveAllChat(cred, contacts, int64(roomID), DefaultInterval)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
