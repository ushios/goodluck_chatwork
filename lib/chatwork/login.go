package chatwork

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
)

const (
	// BaseURL to chatwork
	BaseURL = "https://www.chatwork.com"
)

var (
	AccessTokenRegExp *regexp.Regexp
)

func init() {
	AccessTokenRegExp = regexp.MustCompile(`var ACCESS_TOKEN = '(.+)'`)
}

// Login to chatwork and return ACCESSTOKEN
func Login(email, pass string) ([]byte, error) {
	values := url.Values{}
	values.Add("email", email)
	values.Add("password", pass)
	values.Add("autologin", "on")

	jar, err := cookiejar.New(nil)
	if err != nil {
		return []byte{}, err
	}
	client := &http.Client{
		Jar: jar,
	}
	_, err = client.PostForm(u("/login.php"), values)
	if err != nil {
		return []byte{}, err
	}

	resp, err := client.Get(u("/"))
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	regRes := AccessTokenRegExp.FindSubmatch(body)
	if len(regRes) < 2 {
		return []byte{}, fmt.Errorf("cannot found ACCESS_TOKEN in %s ", u("/"))
	}

	return regRes[1], nil
}

func u(path string) string {
	return fmt.Sprintf("%s%s", BaseURL, path)
}
