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
	Credential struct {
		AccessToken string `json:"access_token"`
		MyID        string `json:"myid"`
	}

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

	client := Client()
	_, err := client.PostForm(u("/login.php"), values)
	if err != nil {
		return &cred, err
	}

	resp, err := client.Get(u("/"))
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

func InitLoad(cred *Credential) (*Contacts, error) {
	client := Client()
	path := fmt.Sprintf(
		"/gateway.php?cmd=init_load&myid=%s&_v=1.80a&_av=4&_t=%s&ln=ja&rid=0&type=&new=1",
		cred.MyID,
		cred.AccessToken,
	)
	resp, err := client.Get(u(path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// TODO: response to json

	return nil, nil
}

func u(path string) string {
	return fmt.Sprintf("%s%s", BaseURL, path)
}

func Client() *http.Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil
	}
	client := &http.Client{
		Jar: jar,
	}

	return client
}
