package utils

import (
	"github.com/cjrd/allocate"
	"reflect"
)

// PageOptions 分页表单
type PageOptions struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

// PipePagerFn 填充默认分页
func PipePagerFn(dataSource interface{}) {
	_ = allocate.Zero(dataSource)
	t := reflect.ValueOf(dataSource).Elem()
	page := t.FieldByName("Page")
	size := t.FieldByName("Size")
	if page.Int() < 1 {
		page.SetInt(1)
	}
	if size.Int() < 1 {
		size.SetInt(10)
	}
}

// PipeLimitFn 分页方法
func PipeLimitFn(db interface{}, dataSource interface{}) {
	data := reflect.ValueOf(dataSource).Elem()
	page := data.FieldByName("Page").Int()
	size := data.FieldByName("Size").Int()
	offset := (page - 1) * size
	refDb := reflect.ValueOf(db)
	refDb.MethodByName("Limit").Call([]reflect.Value{reflect.ValueOf(int(size))})
	refDb.MethodByName("Offset").Call([]reflect.Value{reflect.ValueOf(int(offset))})
}
