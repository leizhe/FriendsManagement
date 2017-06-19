package controllers

import (
	"encoding/json"
	"errors"
	"strings"

	"FriendsManagement/dtos"
	"FriendsManagement/models"

	"github.com/astaxie/beego"
)

// FriendController --- Friend API
type FriendController struct {
	beego.Controller
}

// URLMapping ...
func (c *FriendController) URLMapping() {
	c.Mapping("AddFriend", c.AddFriend)
	c.Mapping("GetAllFriends", c.GetAllFriends)
	c.Mapping("GetCommonFriends", c.GetCommonFriends)
}

// AddFriend ...
// @Title AddFriend
// @Description create a friend connection between two email addresses
// @Param	body		body 	dtos.AddFriendInput	true
// @Success
// @router /AddFriends [post]
func (c *FriendController) AddFriend() {
	var input dtos.AddFriendInput
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &input); err == nil {
		emails := &input.Friends
		u1 := models.User{Email: emails[0]}
		u2 := models.User{Email: emails[1]}
		u1id, u2id, _ := models.AddUsers(u1, u2)
		if result, err := CheckAndAddFriend(u1id, u2id); err == nil {
			c.Data["json"] = result
		} else {
			c.Data["json"] = err.Error()
		}

	} else {
		c.Data["json"] = err.Error()
	}

	c.ServeJSON()

}

// GetAllFriends ...
// @Title GetAllFriends
// @Description retrieve the friends list for an email address
// @Param	body		body 	dtos.GetAllFriendsInput 	true
// @router /GetFriends [post]
func (c *FriendController) GetAllFriends() {
	var input dtos.GetAllFriendsInput
	var output dtos.GetAllFriendsOutput
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &input); err == nil {
		email := input.Email
		if v, n, err := GetFriendsByEmail(email); err == nil {
			output.Success = true
			output.Count = n
			output.Friends = v
			c.Data["json"] = output
		}

	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetCommonFriends ...
// @Title GetCommonFriends
// @Description retrieve the common friends list between two email addresses.
// @Param	body		body 	dtos.GetCommonFriendsInput 	true
// @router /GetCommonFriends [post]
func (c *FriendController) GetCommonFriends() {
	var input dtos.GetCommonFriendsInput
	var output dtos.GetCommonFriendsOutput
	user1Friend := []string{}
	user2Friend := []string{}
	commoFrinds := []string{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &input); err == nil {
		emails := &input.Friends
		if v, _, err := GetFriendsByEmail(emails[0]); err == nil {
			user1Friend = v
			if v, _, err := GetFriendsByEmail(emails[1]); err == nil {
				user2Friend = v
				commoFrinds = CommonFriends(user1Friend, user2Friend)
				output.Success = true
				output.Count = len(commoFrinds)
				output.Friends = commoFrinds
				c.Data["json"] = output
			} else {
				c.Data["json"] = err.Error()
			}

		} else {
			c.Data["json"] = err.Error()
		}

	}
	c.ServeJSON()
}

//CommonFriends ...
func CommonFriends(user1Friend []string, user2Friend []string) (result []string) {
	for _, v1 := range user1Friend {
		for _, v2 := range user2Friend {
			if strings.Compare(v1, v2) == 0 {
				result = append(result, v1)
			}
		}

	}

	return result
}

// GetEmailsByMyFriends ...
func GetEmailsByMyFriends(friends []models.Friend) (v []string) {
	ids := []int64{}
	for _, v := range friends {
		ids = append(ids, v.FriendId)

	}
	result := []string{}
	if users, _, err := models.GetUsersByIds(ids); err == nil {
		for _, v := range users {
			result = append(result, v.Email)

		}
	}

	return result
}

// GetFriendsByEmail ...
func GetFriendsByEmail(mail string) (result []string, num int64, err error) {
	if v, err := models.GetUserByEmail(mail); err == nil {
		id := v.Id
		if v, n, err := models.GetFriendsByUserId(id); err == nil {
			result = GetEmailsByMyFriends(v)
			num = n
		} else {
			return nil, num, errors.New("Error: Get friends fail")
		}
	} else {
		return nil, num, errors.New("Error: This Email does not exist")
	}
	return result, num, err
}

// CheckBlocksSubscribe ...
func CheckBlocksSubscribe(u1id int64, u2id int64) (r bool) {

	result := false
	if v, err := models.GetSubscribeByUserIDAndSubscriberID(u1id, u2id); err == nil {
		if v.Status == 0 {
			result = true
		}
	}
	if v, err := models.GetSubscribeByUserIDAndSubscriberID(u2id, u2id); err == nil {
		if v.Status == 0 {
			result = true
		}
	}
	return result
}

// CheckAndAddFriend ...
func CheckAndAddFriend(u1id int64, u2id int64) (r dtos.BaseResult, err error) {

	var result dtos.BaseResult
	tag := CheckBlocksSubscribe(u1id, u2id)
	if !tag {
		friend := models.Friend{UserId: u1id, FriendId: u2id}
		if _, _, err := models.AddFriend(&friend); err == nil {
			result.Success = true

		}
	} else {
		return result, errors.New("Error: Unable to add friend")
	}
	return result, err
}
