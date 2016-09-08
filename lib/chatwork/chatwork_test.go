package chatwork

import (
	"fmt"
	"testing"
)

func TestInitLoad(t *testing.T) {
	test := func(id, pass string) {
		cred, err := Login(id, pass)
		if err != nil {
			t.Fatal(err)
		}

		res, err := InitLoad(cred)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(res)
	}

	test(email, pass)
}
