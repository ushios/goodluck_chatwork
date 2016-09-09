package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
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

	if email == "" || password == "" {
		return fmt.Errorf("empty email or password")
	}

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

	contacts, err := chatwork.InitLoad(cred)
	if err != nil {
		fmt.Println(err)
		return err
	}

	cTable := tablewriter.NewWriter(os.Stdout)
	cTable.SetHeader([]string{"ID", "Name"})
	for _, contact := range contacts.ContactList {
		row := []string{
			contact.ID,
			contact.Name,
		}
		cTable.Append(row)
	}
	fmt.Println("\n\nContact List ======")
	cTable.Render()

	rTable := tablewriter.NewWriter(os.Stdout)
	rTable.SetHeader([]string{"ID", "Name"})
	for _, room := range contacts.RoomList {
		row := []string{
			room.ID,
			room.Name,
		}
		rTable.Append(row)
	}
	fmt.Println("\n\nRoom List ======")
	rTable.Render()

	return nil
}
