package command

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/ushios/goodluck_chatwork/lib/chatwork"
)

// LoginInfo .
type LoginInfo struct {
	Credential *chatwork.Credential `json:"credential"`
}

// CmdLogin login
func CmdLogin(c *cli.Context) error {
	email := c.String("email")
	password := c.String("password")

	cred, err := chatwork.Login(email, password)
	if err != nil {
		fmt.Println(err)
		return err
	}

	path := c.String("credential")
	info := LoginInfo{
		Credential: cred,
	}
	if err := SaveCredential(path, &info); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Credential file created to %s \n", path)

	_, err = chatwork.InitLoad(cred)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
