package service

type msgService struct{}

var msgInstance *msgService

func NewMsg() *msgService {
	if msgInstance == nil {
		msgInstance = &msgService{}
	}
	return msgInstance
}

func (*msgService) SaveMsg() {

}

func (*msgService) FindMsgsByTime() {

}
