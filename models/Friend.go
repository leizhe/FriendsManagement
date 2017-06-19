package models

import (
	"github.com/astaxie/beego/orm"
)

type Friend struct {
	Id       int   `orm:"column(Id);pk;auto"`
	UserId   int64 `orm:"column(UserId);"`
	FriendId int64 `orm:"column(FriendId);"`
}

func (t *Friend) TableName() string {
	return "Friend"
}

func init() {
	orm.RegisterModel(new(Friend))
}

// AddFriend insert a new Friend into database and returns
// last inserted Id on success.
func AddFriend(m *Friend) (created bool, id int64, err error) {
	o := orm.NewOrm()
	created, id, err = o.ReadOrCreate(m, "UserId", "FriendId")
	return
}

// GetFriendsByUserId retrieves Friends by UserId. Returns error if
// Id doesn't exist
func GetFriendsByUserId(userId int) (v []Friend, n int64, err error) {
	o := orm.NewOrm()
	//var lists []orm.ParamsList
	var friends []Friend
	num, err := o.QueryTable(new(Friend)).Filter("UserId", userId).All(&friends)

	if err == nil {
		return friends, num, nil
	}
	return nil, 0, err
}
