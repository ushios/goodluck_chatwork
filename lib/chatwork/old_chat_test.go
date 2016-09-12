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

		contacts, err := InitLoad(cred)
		if err != nil {
			t.Fatal(err)
		}

		interval := 1 * time.Second
		if err := LoadAndSaveAllChat(cred, contacts, roomID, interval); err != nil {
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

		err = DownloadFile(fID, downloadDirname(1))
		if err != nil {
			t.Fatal(err)
		}
	}

	test(email, pass, 102484735)
}

func TestDownloadRegexp(t *testing.T) {
	test := func(message string, ei string) {
		res := DownloadRegexp.FindStringSubmatch(message)
		if len(res) < 2 {
			t.Fatalf("download regexp match error")
		}

		if ei != res[1] {
			t.Errorf("expected(%s) but (%s)", ei, res[1])
		}
	}

	test(
		`[info][title][dtext:file_uploaded][/title][preview id=102484735 ht=150][download:102484735]\u30b9\u30af\u30ea\u30fc\u30f3\u30b7\u30e7\u30c3\u30c8 2016-09-06 19.28.59.png (132.74 KB)[/download][/info]`,
		"102484735",
	)
}
