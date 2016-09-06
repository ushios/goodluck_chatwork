package cw

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	// BaseURL to chatwork
	BaseURL = "https://www.chatwork.com"
)

// Login to chatwork
func Login(email, pass string) (string, error) {
	values := url.Values{}
	values.Add("email", email)
	values.Add("password", pass)
	values.Add("autologin", "on")

	client := &http.Client{}
	resp, err := client.Post(
		chatworkURL("/login.php"),
		"application/x-www-form-urlencoded",
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		return "", err
	}

	fmt.Println(resp)

	return "", nil
}

func chatworkURL(path string) string {
	return fmt.Sprintf("%s%s", BaseURL, path)
}
