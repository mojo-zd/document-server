package service

import (
	"reflect"

	"github.com/astaxie/beego/orm"
	_ "github.com/document-server/engine"
)

type BaseService struct {
}

func (base BaseService) Read(v interface{}, cols ...string) (err error) {
	err = orm.NewOrm().Read(v, cols...)
	return
}

func (base BaseService) Insert(v interface{}) (err error) {
	_, err = orm.NewOrm().Insert(v)
	return
}

func (base BaseService) InsertMulti(v interface{}) (err error) {
	_, err = orm.NewOrm().InsertMulti(1, v)
	return
}

func (base BaseService) Update(v interface{}, cols ...string) (err error) {
	_, err = orm.NewOrm().Update(v, cols...)
	return
}

func (base BaseService) UpdateWithMap(m map[string]interface{}, tableBean interface{}, conditions ...interface{}) (err error) {
	table := reflect.New(reflect.TypeOf(tableBean)).Interface()
	_, err = orm.NewOrm().QueryTable(table).Filter("Id", 1).Update(m)
	return
}

func (base BaseService) InsertOrUpdate(v interface{}) (err error) {
	_, err = orm.NewOrm().InsertOrUpdate(v)
	return
}

//func getPrivateKey(bean interface{}) {
//	ty := getType(bean)
//}

func getType(bean interface{}) (t reflect.Type) {
	ty := reflect.TypeOf(bean)
	if ty.Kind() == reflect.Ptr {
		t = ty.Elem()
		return
	}
	t = ty
	return
}
