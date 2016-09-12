package chatwork

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

// Download file to dest path
func Download(c *http.Client, url url.URL, dest string) error {
	res, err := c.Get(url.String())
	if err != nil {
		return err
	}
	defer res.Body.Close()

	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(dest, d, 0644); err != nil {
		return err
	}

	return nil
}
