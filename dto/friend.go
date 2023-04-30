package dto

type PushRequest struct {
	To uint `json:"to" form:"to"` //请求好友id
}

type AcceptRequest struct {
	ReqUid uint `json:"req_uid" form:"req_uid"` //请求者id
}
