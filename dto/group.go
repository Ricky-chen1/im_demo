package dto

type CreateGroup struct {
	Name string `json:"name" form:"name"`
}

type JoinGroup struct {
	Gid uint `json:"gid" form:"gid"`
}
