package session

import (
	"logf"
	"reflect"
)

const (
	BeforeQuery  = "BeforeQuery "
	AfterQuery   = "AfterQuery"
	BeforeInsert = "BeforeInsert"
	AfterInsert  = "AfterInsert"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate  = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete  = "AfterDelete"
)

func (s *Session) CallMethod(methodName string, val any) {
	method := reflect.ValueOf(s.RefTable().Model).MethodByName(methodName)
	if val != nil {
		method = reflect.ValueOf(val).MethodByName(methodName)
	}
	param := []reflect.Value{reflect.ValueOf(s)}
	if method.IsValid() {
		if call := method.Call(param); len(call) > 0 {
			if err := call[0].Interface().(error); err != nil {
				logf.Error(err)
			}
		}
	}
}
