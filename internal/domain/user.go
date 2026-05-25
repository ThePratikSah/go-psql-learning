package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Name      string    `json:"name"`
	Role      UserRole  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRole string

const (
	UserRoleAdmin   UserRole = "admin"
	UserRoleMember  UserRole = "member"
	UserRoleViewer  UserRole = "viewer"
)

type UserCreateInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// IsAdmin returns true if the user has admin role.
func (u User) IsAdmin() bool {
	return u.Role == UserRoleAdmin
}

// CanPerform checks if the user can perform the given action.
// Admins can do everything. Members can create/update. Viewers can only read.
//
// NOTE: Value receiver — only reading fields. This is the foundation for RBAC.
// Compare to Node.js: `if (user.role === 'admin' || action === 'read')`
func (u User) CanPerform(action string) bool {
	if u.Role == UserRoleAdmin {
		return true
	}
	switch action {
	case "create", "update", "delete":
		return u.Role == UserRoleMember
	case "read", "view", "list":
		return true
	}
	return false
}
