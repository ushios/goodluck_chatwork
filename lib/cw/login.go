package chatwork

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

const (
	// BaseURL to chatwork
	BaseURL = "https://www.chatwork.com"

	// Wait for interval
	Wait = "1s"
)

var (
	wait time.Duration
)

func init() {
	wait = waitInterval()
}

// Login to chatwork
func Login(email, pass string) (string, error) {
	values := url.Values{}
	values.Add("email", email)
	values.Add("password", pass)
	values.Add("autologin", "on")

	// Client making
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Transport: tr,
		Jar:       jar,
	}

	// Login action
	resp, err := client.PostForm(u("/login.php"), values)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	fmt.Println(resp)

	// Sleep
	time.Sleep(wait)

	// Login check
	resp, err = client.Get(u("/"))
	if err != nil {
		return "", nil
	}

	fmt.Println(resp)

	return "", nil
}

func u(path string) string {
	return fmt.Sprintf("%s%s", BaseURL, path)
}

func waitInterval() time.Duration {
	d, err := time.ParseDuration(Wait)
	if err != nil {
		panic(err)
	}
	return d
}
