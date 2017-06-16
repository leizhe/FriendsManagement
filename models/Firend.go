package models

import (
	"github.com/astaxie/beego/orm"
)

type Firend struct {
	Id       int   `orm:"column(Id);pk;auto"`
	UserId   int64 `orm:"column(UserId);"`
	FirendId int64 `orm:"column(FirendId);"`
}

func (t *Firend) TableName() string {
	return "Firend"
}

func init() {
	orm.RegisterModel(new(Firend))
}

// AddFirend insert a new Firend into database and returns
// last inserted Id on success.
func AddFirend(m *Firend) (created bool, id int64, err error) {
	o := orm.NewOrm()
	created, id, err = o.ReadOrCreate(m, "UserId", "FirendId")
	return
}

// GetFirendsByUserId retrieves Firends by UserId. Returns error if
// Id doesn't exist
func GetFirendsByUserId(userId int) (v []Firend, n int64, err error) {
	o := orm.NewOrm()
	//var lists []orm.ParamsList
	var firends []Firend
	num, err := o.QueryTable(new(Firend)).Filter("UserId", userId).All(&firends)

	if err == nil {
		return firends, num, nil
	}
	return nil, 0, err
}
