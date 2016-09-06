package chatwork

import (
	"fmt"
	"testing"
)

var (
	email = ""
	pass  = ""
)

func TestLogin(t *testing.T) {
	test := func(id, pass string) {
		res, err := Login(id, pass)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(res)
	}

	test(email, pass)
}
