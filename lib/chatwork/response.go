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
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	// InitLoadResult is result of InitLoad
	InitLoadResult struct {
		RoomDat         map[string]RoomDat    `json:"room_dat"`
		ContactDat      map[string]ContactDat `json:"contact_dat"`
		MentionDat      interface{}           `json:"mention_dat"`
		MyRequestDat    interface{}           `json:"myrequest_dat"`
		RequestDat      interface{}           `json:"request_dat"`
		SettingData     interface{}           `json:"setting_data"`
		CatDat          interface{}           `json:"cat_dat"`
		AnnounceDat     interface{}           `json:"announce_dat"`
		AnnounceID      int                   `json:"announce_id"`
		LastID          string                `json:"last_id"`
		Storage         interface{}           `json:"storage"`
		StorageLimit    string                `json:"storage_limit"`
		ChatworkID      string                `json:"chatwork_id"`
		Plan            string                `json:"plan"`
		PayPlanName     string                `json:"pay_plan_name"`
		PayType         string                `json:"pay_type"`
		StartTime       int64                 `json:"start_time"`
		IsBusiness      bool                  `json:"is_business"`
		IsSecurity      bool                  `json:"is_security"`
		IsAdmin         bool                  `json:"is_admin"`
		IsAdminUser     bool                  `json:"is_admin_user"`
		IsEnterprise    bool                  `json:"is_enterprise"`
		Limit           interface{}           `json:"limit"`
		AvailableOption interface{}           `json:"available_option"`
		ContactLimitNum interface{}           `json:"contact_limit_num"`
		GroupLimitNum   int                   `json:"group_limit_num"`
	}

	// AccountResult is result of account info
	AccountResult struct {
		AccountDat map[string]AccountDat `json:"account_dat"`
	}

	// RoomDat is room data
	RoomDat struct {
		MID  int64          `json:"mid"`
		R    int            `json:"r"`
		S    int            `json:"s"`
		TP   int            `json:"tp"` // 1: group, 2:DM, 3: self
		C    int            `json:"c"`
		F    int            `json:"f"`
		T    int            `json:"t"`
		LT   int64          `json:"lt"`
		Name string         `json:"n,omitempty"`
		LN   string         `json:"ln,omitempty"`
		IC   string         `json:"ic,omitempty"`
		M    map[string]int `json:"m"`
	}

	// ContactDat is contacts data
	ContactDat struct {
		AID       int64  `json:"aid"`
		GID       int64  `json:"gid"`
		Nm        string `json:"nm"`
		CWID      string `json:"cwid"`
		ONM       string `json:"onm"`
		DP        string `json:"dp"`
		SP        string `json:"sp"`
		FB        string `json:"fb"`
		TW        string `json:"tw"`
		UD        int    `json:"ud"`
		AV        string `json:"av"`
		CV        string `json:"cv"`
		IsDeleted bool   `json:"is_deleted"`
		Name      string `json:"name"`
		AVID      string `json:"avid"`
		MRID      int64  `json:"mrid"`
		RID       int64  `json:"rid"`
	}

	// AccountDat is account
	AccountDat struct {
		ContactDat
	}

	// LoadOldChatResult is result
	LoadOldChatResult struct {
		ChatList []ChatMessage `json:"chat_list"`
	}

	// ChatMessage is a message of chat
	ChatMessage struct {
		ID      int64  `json:"id"`
		AID     int64  `json:"aid"`
		Message string `json:"msg"`
		TM      int    `json:"tm"`
		UTM     int    `json:"utm"`
	}
)

// ReadResponse reading http response
func ReadResponse(r *http.Response, result interface{}) (*Response, error) {
	resp := Response{
		Result: result,
	}

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	// fmt.Println(string(d))
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

// Len for sort
func (l LoadOldChatResult) Len() int {
	return len(l.ChatList)
}

// Swap for sort
func (l LoadOldChatResult) Swap(i, j int) {
	l.ChatList[i], l.ChatList[j] = l.ChatList[j], l.ChatList[i]
}

// Less for sort
func (l LoadOldChatResult) Less(i, j int) bool {
	return l.ChatList[i].ID < l.ChatList[j].ID
}
