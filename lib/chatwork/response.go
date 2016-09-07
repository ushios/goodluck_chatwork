package chatwork

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/guregu/null"
)

type (
	// Response is api response root
	Response struct {
		Status       Status      `json:"status"`
		Result       interface{} `json:"result"`
		Message      null.String `json:"message,omitempty"`
		ResultString string      `json:"-"`
	}

	// Status have api response success or not
	Status struct {
		Success bool `json:"success"`
	}
)

// ReadResponse reading http response
func ReadResponse(r *http.Response) (*Response, error) {
	resp := Response{}

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(d, &resp); err != nil {
		return nil, err
	}

	if !resp.Status.Success {
		if resp.Message.IsZero() {
			return nil, fmt.Errorf("unknown api error")
		}
		return nil, fmt.Errorf(resp.Message.String)
	}

	resp.ResultString = string(d)

	return &resp, nil
}
