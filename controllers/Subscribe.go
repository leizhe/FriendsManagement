package controllers

import (
	"FriendsManagement/dtos"
	"FriendsManagement/models"
	"encoding/json"
	"errors"
	"strings"

	"github.com/astaxie/beego"
)

// SubscribeController --- Subscribe API
type SubscribeController struct {
	beego.Controller
}

// URLMapping ...
func (c *SubscribeController) URLMapping() {
	c.Mapping("AddSubscribe", c.AddSubscribe)
	c.Mapping("BlockSubscribe", c.BlockSubscribe)
	c.Mapping("RetrieveSubscribe", c.RetrieveSubscribe)
}

// AddSubscribe ...
// @Title AddSubscribe
// @Description subscribe to updates from an email address
// @Param	body		body 	dtos.SubscribeInput	true
// @Success
// @router /AddSubscribe [post]
func (c *SubscribeController) AddSubscribe() {
	var input dtos.SubscribeInput
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &input); err == nil {
		if result, err := BlackandWhiteSubscribe(input, 1); err == nil {
			c.Data["json"] = result
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// BlockSubscribe ...
// @Title BlockSubscribe
// @Description block updates from an email address.
// @Param	body		body 	dtos.SubscribeInput	true
// @router /BlockSubscribe [post]
func (c *SubscribeController) BlockSubscribe() {
	var input dtos.SubscribeInput
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &input); err == nil {
		if result, err := BlackandWhiteSubscribe(input, 0); err == nil {
			c.Data["json"] = result
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// RetrieveSubscribe ...
// @Title RetrieveSubscribe
// @Description retrieve all email addresses that can receive updates from an email address.
// @Param	body		body 	dtos.RetrieveSubscribeInput	true
// @router /RetrieveSubscribe [post]
func (c *SubscribeController) RetrieveSubscribe() {
	var input dtos.RetrieveSubscribeInput
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &input); err == nil {
		if v, err := GetRetrieveSubscriberIds(input); err == nil {
			if output, err := GetRetrieveSubscribeResult(v); err == nil {
				c.Data["json"] = output
			} else {
				c.Data["json"] = err.Error()
			}
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// BlackandWhiteSubscribe ...
func BlackandWhiteSubscribe(input dtos.SubscribeInput, status int) (result dtos.BaseResult, err error) {
	var (
		requestorID int64
		targetID    int64
	)
	if r, t, err := GetUserIDByInput(input); err == nil {
		requestorID = r
		targetID = t
	}
	result, err = AddSubscribeByStatus(requestorID, targetID, status)
	return result, err

}

// GetRetrieveSubscribeResult ...
func GetRetrieveSubscribeResult(ids []int64) (result dtos.RetrieveSubscribeOutput, err error) {
	if len(ids) > 0 {
		result.Success = true
		result.Recipients = GetEmailsByUserIds(ids)
	} else {
		return result, errors.New("Error: No subscriber")
	}

	return result, err
}

// GetRetrieveSubscriberIds ...
func GetRetrieveSubscriberIds(input dtos.RetrieveSubscribeInput) (result []int64, err error) {

	friendIds := []int64{}
	subscriberIds := []int64{}
	mentionedIds := []int64{}
	if v, err := GetFriendIdsByEmail(input.Sender); err == nil {
		friendIds = v
	}
	if v, err := GetSubscriberIdsByEmail(input.Sender); err == nil {
		subscriberIds = v
	}
	if v, err := GetMentionedIdsByText(input.Text); err == nil {
		mentionedIds = v
	}
	result = MergeArray(result, friendIds)
	result = MergeArray(result, subscriberIds)
	result = MergeArray(result, mentionedIds)
	result = ClearNoSubscriber(result, input.Sender)
	return result, err
}

// ClearNoSubscriber ...
func ClearNoSubscriber(ids []int64, mail string) (result []int64) {
	NoSubscriber := []int64{}
	if v, err := models.GetUserByEmail(mail); err == nil {
		id := v.Id
		if v, _, err := models.GetSubscribeByUserIDandStatus(id, 0); err == nil {
			for _, v := range v {
				NoSubscriber = append(NoSubscriber, v.UserId)
			}

			for _, y := range ids {
				if c := Contain(NoSubscriber, y); !c {
					result = append(result, y)
				}
			}

		}
	}

	return result
}

// GetEmailsByUserIds ...
func GetEmailsByUserIds(ids []int64) (result []string) {
	if users, _, err := models.GetUsersByIds(ids); err == nil {
		for _, v := range users {
			result = append(result, v.Email)

		}
	}
	return result
}

// GetMentionedIdsByText ...
func GetMentionedIdsByText(text string) (ids []int64, err error) {
	if v, _, err := models.GetAllUser(); err == nil {
		for _, u := range v {
			if strings.Contains(text, u.Email) {
				ids = append(ids, int64(u.Id))
			}
		}
	} else {
		return nil, errors.New("Error: Get Users fail")
	}
	return ids, err
}

// GetSubscriberIdsByEmail ...
func GetSubscriberIdsByEmail(email string) (v []int64, err error) {
	ids := []int64{}
	if v, err := models.GetUserByEmail(email); err == nil {
		id := v.Id
		if v, _, err := models.GetSubscribeByUserIDandStatus(id, 1); err == nil {
			for _, v := range v {
				ids = append(ids, v.UserId)
			}
		} else {
			return nil, errors.New("Error: Get firends fail")
		}
	} else {
		return nil, errors.New("Error: This Email does not exist")
	}
	return ids, err
}

// GetFriendIdsByEmail ...
func GetFriendIdsByEmail(email string) (ids []int64, err error) {
	if v, err := models.GetUserByEmail(email); err == nil {
		id := v.Id
		if v, _, err := models.GetFirendsByUserId(id); err == nil {
			for _, v := range v {
				ids = append(ids, v.FirendId)
			}
		} else {
			return nil, errors.New("Error: Get firends fail")
		}
	} else {
		return nil, errors.New("Error: This Email does not exist")
	}
	return ids, err
}

// AddSubscribeByStatus ...
func AddSubscribeByStatus(rid int64, tid int64, status int) (r dtos.BaseResult, err error) {
	var result dtos.BaseResult
	v, err := models.GetSubscribeByUserIDAndSubscriberID(rid, tid)
	resultSubscribe := v
	if resultSubscribe == nil {
		subscribe := models.Subscribe{UserId: rid, SubscriberId: tid, Status: status}
		if _, _, err := models.AddSubscribe(&subscribe); err == nil {
			result.Success = true
			return result, nil

		}
	} else {
		resultSubscribe.Status = status
		if err := models.UpdateSubscribeById(resultSubscribe); err == nil {
			result.Success = true
			return result, nil
		}
	}

	return result, err
}

// GetUserIDByInput ...
func GetUserIDByInput(input dtos.SubscribeInput) (requestorID int64, targetID int64, err error) {
	requestorID, err = GetUserIDByEmail(input.Requestor)
	targetID, err = GetUserIDByEmail(input.Target)
	return requestorID, targetID, err
}

// GetUserIDByEmail ...
func GetUserIDByEmail(mail string) (r int64, err error) {
	var resultID int64
	if v, err := models.GetUserByEmail(mail); err == nil {
		resultID = int64(v.Id)
	} else {
		user := models.User{Email: mail}
		if _, id, err := models.AddUser(&user); err == nil {
			resultID = id
		}
	}
	return resultID, err
}

// Contain ...
func Contain(nums []int64, num int64) (r bool) {
	result := false
	for _, n := range nums {
		if num == n {
			result = true
			return result
		}
	}
	return result
}

// MergeArray ...
func MergeArray(a1 []int64, a2 []int64) (r []int64) {
	result := []int64{}
	for _, n := range a1 {
		result = append(result, n)
	}
	for _, n := range a2 {
		if c := Contain(result, n); !c {
			result = append(result, n)
		}
	}
	return result
}
