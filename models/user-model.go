package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id       int64     `orm:"auto;pk" form:"-"`
	Email    string    `orm:"size(100);unique" form:"email,text,login" valid:"MaxSize(100)"`
	Password string    `orm:"size(100)" form:"password,text,password" valid:"MaxSize(100)"`
	Created  time.Time `orm:"type(datetime);auto_now_add" form:"-"`
	Updated  time.Time `orm:"type(datetime);null"  form:"-"`
}

func AddUser(user *User) (int64, error) {
	var err error
	var id int64
	o := orm.NewOrm()
	o.Using("default")
	beego.Trace("Try to create new user:", user.Email)

	if id, err = o.Insert(user); err != nil {
		beego.Error("Error:", err)
		return 0, err
	}
	return id, err
}

func init() {
	orm.RegisterModel(new(User))
}
