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
	MyIDRegExp        *regexp.Regexp
)

func init() {
	AccessTokenRegExp = regexp.MustCompile(`var ACCESS_TOKEN = '(.+)'`)
	MyIDRegExp = regexp.MustCompile(`var myid = '(.+)'`)
}

type (
	LoginResponse struct {
		AccessToken string
		MyID        string
	}
)

// Login to chatwork and return ACCESSTOKEN
func Login(email, pass string) (*LoginResponse, error) {
	lr := LoginResponse{}

	values := url.Values{}
	values.Add("email", email)
	values.Add("password", pass)
	values.Add("autologin", "on")

	jar, err := cookiejar.New(nil)
	if err != nil {
		return &lr, err
	}
	client := &http.Client{
		Jar: jar,
	}
	_, err = client.PostForm(u("/login.php"), values)
	if err != nil {
		return &lr, err
	}

	resp, err := client.Get(u("/"))
	if err != nil {
		return &lr, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &lr, err
	}

	tokenRegRes := AccessTokenRegExp.FindSubmatch(body)
	if len(tokenRegRes) < 2 {
		return &lr, fmt.Errorf("cannot found ACCESS_TOKEN in %s ", u("/"))
	}
	lr.AccessToken = string(tokenRegRes[1])

	myRegRes := MyIDRegExp.FindSubmatch(body)
	if len(myRegRes) < 2 {
		return &lr, fmt.Errorf("cannot found myid in %s", u("/"))
	}
	lr.MyID = string(myRegRes[1])

	return &lr, nil
}

func u(path string) string {
	return fmt.Sprintf("%s%s", BaseURL, path)
}
