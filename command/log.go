package command

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/ushios/goodluck_chatwork/lib/chatwork"
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

	roomID := c.Int64("room")
	err = chatwork.LoadAndSaveAllChat(cred, roomID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
