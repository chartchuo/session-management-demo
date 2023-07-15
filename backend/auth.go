package main

import (
	"backend/model"
	"errors"

	"github.com/ucarion/urlpath"
)

type auth struct {
	method string
	path   string
	roles  []string
}

var authTable []auth = []auth{
	//require userid matching
	{"GET", "/user/:userid/*", []string{"user"}},
	//admin role
	{"GET", "/admin/*", []string{"admin"}},
	{"GET", "/user/*", []string{"admin"}},
}

var authTablePaths []urlpath.Path

func init() {
	authTablePaths = make([]urlpath.Path, len(authTable))
	for i, a := range authTable {
		authTablePaths[i] = urlpath.New(a.path)
	}
}

func findRole(role string, roles []string) bool {
	for _, r := range roles {
		if role == r {
			return true
		}
	}
	return false
}

// return true when authorize
func authorize(u *model.User, method string, path string) (allow bool, err error) {
	for i, a := range authTablePaths {
		match, ok := a.Match(path)
		if !ok {
			continue
		}

		//verify role
		if !findRole(u.Role, authTable[i].roles) {
			continue
		}

		//verify userid
		if match.Params["userid"] != "" && match.Params["userid"] != u.UserID {
			continue
		}

		return true, nil
	}

	return false, errors.New("unauthorize")
}
