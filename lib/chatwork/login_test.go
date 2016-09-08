package chatwork

import (
	"fmt"
	"testing"
)

var (
	// Test account
	email = "ushio.s+goodluck@gmail.com"
	pass  = "goodluck"
)

func TestLogin(t *testing.T) {
	test := func(id, pass string) {
		res, err := Login(id, pass)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println("ACCESS_TOKEN: ", res.AccessToken)
		fmt.Println("myid:", res.MyID)
	}

	test(email, pass)
}
