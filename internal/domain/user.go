package domain

import "time"

type Role string

const (
	RoleUser      Role = "user"
	RoleModerator Role = "moderator"
	RoleAdmin     Role = "admin"
)

var RolesPermMap = map[Role]int{
	RoleUser:      0,
	RoleModerator: 1,
	RoleAdmin:     2,
}

type User struct {
	ID        uint64
	Username  string
	Role      Role
	IsActive  bool
	CreatedAt time.Time
}

func (u User) HasRequiredPerms(minRole Role) bool {
	return RolesPermMap[u.Role] >= RolesPermMap[minRole]
}
