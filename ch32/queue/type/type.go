package _type

type HandlerFunc func(args interface{}) (traceId string, err error)

func (h HandlerFunc) Call(args interface{}) (traceId string, err error) {
	return h(args)
}

type Data struct {
	Body          []byte
	MessageHandel interface{}
}
type Sender struct {
	QueueName string `json:"queue_name,omitempty"` //队列名称
	Job       string `json:"job,omitempty"`        //任务名称
	TraceID   string `json:"trace_id,omitempty"`
	Msg       string `json:"msg,omitempty"` //消息体,回调时使用
}
