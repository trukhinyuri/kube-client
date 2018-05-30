package model

type UserGroupAccess string

const (
	AdminAccess  UserGroupAccess = "admin"
	MasterAccess UserGroupAccess = "master"
	MemberAccess UserGroupAccess = "member"
	GuestAccess  UserGroupAccess = "guest"
	NoAccess     UserGroupAccess = "none"
)

// UserGroup -- group of users
//
// swagger:model
type UserGroup struct {
	ID           string           `json:"id,omitempty"`
	Label        string           `json:"label"`
	OwnerID      string           `json:"owner_user_id,omitempty"`
	Members      UserGroupMembers `json:"members,omitempty"`
	MembersCount uint             `json:"members_count,omitempty"`
	//creation date in RFC3339 format
	CreatedAt string `json:"created_at,omitempty"`
}

// UserGroupMembers -- list of user group members
//
// swagger:model
type UserGroupMembers struct {
	Members []UserGroupMember `json:"members"`
}

// UserGroupMember -- group member
//
// swagger:model
type UserGroupMember struct {
	Username string          `json:"username"`
	Access   UserGroupAccess `json:"access"`
}