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
	// AccessTokenRegExp for getting access token from html
	AccessTokenRegExp *regexp.Regexp
	// MyIDRegExp for getting myid from html
	MyIDRegExp *regexp.Regexp
	// c is http client
	c *http.Client
)

func init() {
	AccessTokenRegExp = regexp.MustCompile(`var ACCESS_TOKEN = '(.+)'`)
	MyIDRegExp = regexp.MustCompile(`var myid = '(.+)'`)
}

type (
	// Credential have login info
	Credential struct {
		AccessToken string `json:"access_token"`
		MyID        string `json:"myid"`
	}

	// Contacts have contacts and room info
	Contacts struct {
		Contacts interface{}
		Rooms    interface{}
	}
)

// Login to chatwork and return ACCESSTOKEN
func Login(email, pass string) (*Credential, error) {
	cred := Credential{}

	values := url.Values{}
	values.Add("email", email)
	values.Add("password", pass)
	values.Add("autologin", "on")

	_, err := client().PostForm(u("/login.php"), values)
	if err != nil {
		return &cred, err
	}

	resp, err := client().Get(u("/"))
	if err != nil {
		return &cred, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &cred, err
	}

	tokenRegRes := AccessTokenRegExp.FindSubmatch(body)
	if len(tokenRegRes) < 2 {
		return &cred, fmt.Errorf("cannot found ACCESS_TOKEN in %s ", u("/"))
	}
	cred.AccessToken = string(tokenRegRes[1])

	myRegRes := MyIDRegExp.FindSubmatch(body)
	if len(myRegRes) < 2 {
		return &cred, fmt.Errorf("cannot found myid in %s", u("/"))
	}
	cred.MyID = string(myRegRes[1])

	return &cred, nil
}

func u(path string) string {
	return fmt.Sprintf("%s%s", BaseURL, path)
}

func client() *http.Client {
	if c != nil {
		return c
	}

	UsedJar, err := cookiejar.New(nil)
	if err != nil {
		return nil
	}
	c = &http.Client{
		Jar: UsedJar,
	}

	return c
}
