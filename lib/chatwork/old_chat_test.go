package chatwork

import (
	"fmt"
	"testing"
	"time"
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

		interval := 1 * time.Second
		if err := LoadAndSaveAllChat(cred, roomID, interval); err != nil {
			t.Fatal(err)
		}

	}

	test(email, pass, 57468721)
}

func TestDownloadFile(t *testing.T) {
	test := func(id, pass string, fID int64) {
		_, err := Login(id, pass)
		if err != nil {
			t.Fatal(err)
		}

		err = DownloadFile(1, fID)
		if err != nil {
			t.Fatal(err)
		}
	}

	test(email, pass, 102484735)
}
