package chatwork

import (
	"fmt"
	"testing"
)

func TestLoadOldChat(t *testing.T) {
	test := func(id, pass string, roomID int64) {
		cred, err := Login(id, pass)
		if err != nil {
			t.Fatal(err)
		}

		res, err := LoadOldChat(cred, roomID)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(res)
	}

	test(email, pass, 57468721)

}
