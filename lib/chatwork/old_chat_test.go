package chatwork

import (
	"fmt"
	"testing"
)

func TestLoadOldChat(t *testing.T) {
	test := func(id, pass string, roomID, fcID int64) {
		cred, err := Login(id, pass)
		if err != nil {
			t.Fatal(err)
		}

		res, err := LoadOldChat(cred, roomID, fcID)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(res)
	}

	test(email, pass, 57468721, 0)

}

func TestLoadAndSaveAllChat(t *testing.T) {
	test := func(id, pass string, roomID int64) {
		cred, err := Login(id, pass)
		if err != nil {
			t.Fatal(err)
		}

		if err := LoadAndSaveAllChat(cred, roomID); err != nil {
			t.Fatal(err)
		}

	}

	test(email, pass, 57468721)
}
