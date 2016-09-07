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

	// InitLoadResult is result of InitLoad
	InitLoadResult struct {
		RoomDat         interface{} `json:"room_dat"`
		ContactDat      interface{} `json:"contact_dat"`
		MentionDat      interface{} `json:"mention_dat"`
		MyRequestDat    interface{} `json:"myrequest_dat"`
		RequestDat      interface{} `json:"request_dat"`
		SettingData     interface{} `json:"setting_data"`
		CatDat          interface{} `json:"cat_dat"`
		AnnounceDat     interface{} `json:"announce_dat"`
		AnnounceID      int         `json:"announce_id"`
		LastID          string      `json:"last_id"`
		Storage         interface{} `json:"storage"`
		StorageLimit    int         `json:"storage_limit"`
		ChatworkID      int         `json:"chatwork_id"`
		Plan            string      `json:"plan"`
		PayPlanName     string      `json:"pay_plan_name"`
		PayType         string      `json:"pay_type"`
		StartTime       int64       `json:"start_time"`
		IsBusiness      bool        `json:"is_business"`
		IsSecurity      bool        `json:"is_security"`
		IsAdmin         bool        `json:"is_admin"`
		IsAdminUser     bool        `json:"is_admin_user"`
		IsEnterprise    bool        `json:"is_enterprise"`
		Limit           interface{} `json:"limit"`
		AvailableOption interface{} `json:"available_option"`
		ContactLimitNum int         `json:"contact_limit_num"`
		GroupLimitNum   int         `json:"group_limit_num"`
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
