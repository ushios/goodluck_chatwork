package command

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/ushios/goodluck_chatwork/lib/chatwork"
)

// CmdLogin login
func CmdLogin(c *cli.Context) error {
	email := c.String("email")
	password := c.String("password")

	cred, err := chatwork.Login(email, password)
	if err != nil {
		return err
	}

	path := c.String("credential")
	if err := SaveCredential(path, cred); err != nil {
		return err
	}

	fmt.Printf("Credential file created to %s \n", path)

	return nil
}
