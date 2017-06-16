package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id    int    `orm:"column(Id);pk;auto"`
	Email string `orm:"column(Email);null"`
}

func (t *User) TableName() string {
	return "User"
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUsers insert a new User into database and returns
// last inserted Id on success.
func AddUsers(u1 User, u2 User) (u1id int64, u2id int64, err error) {
	_, u1id, err = AddUser(&u1)
	_, u2id, err = AddUser(&u2)
	return
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (created bool, id int64, err error) {
	o := orm.NewOrm()
	created, id, err = o.ReadOrCreate(m, "Email")
	return
}

// GetUserByEmail retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserByEmail(email string) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Email: email}
	if err = o.Read(v, "Email"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetUsersByIds ...
func GetUsersByIds(ids []int64) (v []User, n int64, err error) {
	o := orm.NewOrm()
	var users []User
	num, err := o.QueryTable(new(User)).Filter("Id__in", ids).All(&users)

	if err == nil {
		return users, num, nil
	}
	return nil, 0, err
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser() (v []User, n int64, err error) {
	o := orm.NewOrm()
	var users []User
	num, err := o.QueryTable(new(User)).All(&users)
	if err == nil {
		return users, num, nil
	}
	return nil, 0, err

}
