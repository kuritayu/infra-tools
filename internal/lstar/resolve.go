package lstar

import (
	"os/user"
	"strconv"
)

func ResolveUser(id int) string {
	var owner string
	uid := strconv.Itoa(id)
	u, err := user.LookupId(uid)
	if err != nil {
		owner = uid
	} else {
		owner = u.Username
	}
	return owner
}

func ResolveGroup(id int) string {
	var group string
	gid := strconv.Itoa(id)
	g, err := user.LookupGroupId(gid)
	if err != nil {
		group = gid
	} else {
		group = g.Name
	}
	return group
}
