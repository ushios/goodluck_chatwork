package command

import (
	"fmt"
	"log"

	"github.com/urfave/cli"
	"github.com/ushios/goodluck_chatwork/lib/chatwork"
)

// CmdLogAll show chat log
func CmdLogAll(c *cli.Context) error {
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

	for rID, room := range contacts.RoomMap {
		log.Printf("Start - %s \n", room.Name)
		err = chatwork.LoadAndSaveAllChat(cred, contacts, rID, DefaultInterval)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}
