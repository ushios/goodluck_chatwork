package chatwork

import (
	"fmt"
	"testing"
)

func TestAccountInfo(t *testing.T) {
	test := func(id, pass string) {
		cred, err := Login(id, pass)
		if err != nil {
			t.Fatal(err)
		}

		contacts, err := InitLoad(cred)
		if err != nil {
			t.Fatal(err)
		}

		res, err := AccountInfo(cred, contacts)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(res)

	}

	test(email, pass)
}
