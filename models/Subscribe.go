package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Subscribe struct {
	Id           int   `orm:"column(Id);pk;auto"`
	UserId       int64 `orm:"column(UserId);"`
	SubscriberId int64 `orm:"column(SubscriberId);"`
	Status       int   `orm:"column(Status);null"`
}

func (t *Subscribe) TableName() string {
	return "Subscribe"
}

func init() {
	orm.RegisterModel(new(Subscribe))
}

// AddSubscribe insert a new Subscribe into database and returns
// last inserted Id on success.
func AddSubscribe(m *Subscribe) (created bool, id int64, err error) {
	o := orm.NewOrm()
	created, id, err = o.ReadOrCreate(m, "UserId", "SubscriberId")
	return
}

// GetSubscribeByUserIDandStatus retrieves Firends by userID status. Returns error if
// Id doesn't exist
func GetSubscribeByUserIDandStatus(userID int, status int) (v []Subscribe, n int64, err error) {
	o := orm.NewOrm()
	var subscribers []Subscribe
	num, err := o.QueryTable(new(Subscribe)).Filter("SubscriberId", userID).Filter("Status", status).All(&subscribers)

	if err == nil {
		return subscribers, num, nil
	}
	return nil, 0, err
}

// GetSubscribeByUserIDAndSubscriberID retrieves Subscribe by UserIDAndSubscriberID. Returns error if
// Id doesn't exist
func GetSubscribeByUserIDAndSubscriberID(userID int64, subscriberID int64) (v *Subscribe, err error) {
	o := orm.NewOrm()
	v = &Subscribe{UserId: userID, SubscriberId: subscriberID}
	if err = o.Read(v, "UserId", "SubscriberId"); err == nil {
		return v, nil
	}
	return nil, err
}

// UpdateSubscribe updates Subscribe by Id and returns error if
// the record to be updated doesn't exist
func UpdateSubscribeById(m *Subscribe) (err error) {
	o := orm.NewOrm()
	v := Subscribe{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}
