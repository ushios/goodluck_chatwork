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

func TestFileInfo(t *testing.T) {
	test := func(id, pass string, fID int64) {
		cred, err := Login(id, pass)
		if err != nil {
			t.Fatal(err)
		}

		res, err := DownloadFileInfo(cred, fID)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println("filename:", res.Filename)
		fmt.Println("url:", res.URL)
	}

	test(email, pass, 102484735)
}
