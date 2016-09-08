package chatwork

import "testing"

func TestInitLoad(t *testing.T) {
	test := func(id, pass string) {
		cred, err := Login(id, pass)
		if err != nil {
			t.Fatal(err)
		}

		_, err = InitLoad(cred)
		if err != nil {
			t.Fatal(err)
		}

	}

	test(email, pass)
}
