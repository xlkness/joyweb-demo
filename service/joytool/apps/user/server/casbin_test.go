package server

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"testing"
)

func TestCasbin(t *testing.T) {
	e, err := casbin.NewEnforcer("path/model.conf", "path/policy.csv")
	if err != nil {
		panic(err)
	}

	e.EnableAutoSave(false)

	enforce(e, "alice", "/createuser", "read", false)
	enforce(e, "alice", "/createuser", "write", true)
	enforce(e, "bob", "/createuser", "read", false)
	enforce(e, "bob", "/createuser", "write", false)

	e.AddPolicy("userAdmin", "/createuser", "read")
	enforce(e, "bob", "/createuser", "read", true)
	enforce(e, "bob", "/createuser", "write", false)

	enforce(e, "marry", "/createuser", "write", false)
	e.AddGroupingPolicy("marry", "superAdmin")
	enforce(e, "marry", "/createuser", "write", true)

	e.SavePolicy()
}

func enforce(e *casbin.Enforcer, sub, obj, act string, expected bool) bool {
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		panic(err)
	}
	if ok == expected {
		return true
	}
	panic(fmt.Errorf("%v %v %v not passed, expected %v but %v", sub, obj, act, expected, ok))
	return ok
}
