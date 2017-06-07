package webserver

type BaseJsonBean struct {
	Code    int         `jsontest:"code"`
	Data    interface{} `jsontest:"data"`
	Message string      `jsontest:"message"`
}

func NewBaseJsonBean() *BaseJsonBean {
	return &BaseJsonBean{}
}