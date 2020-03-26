package utils

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/streadway/amqp"
	"github.com/yossefaz/go_utils/queue"
	globalUtils "github.com/yossefaz/go_utils/utils"
)

var pFile = globalUtils.PickFile{
	Id:   0,
	Path: "",
	Result: map[string]int{
		"BorderExist": 1,
		"InsideJer":   0,
	},
	Status: 0,
}

var message, _ = json.Marshal(pFile)

var m = amqp.Delivery{
	Acknowledger:    nil,
	Headers:         nil,
	ContentType:     "",
	ContentEncoding: "",
	DeliveryMode:    0,
	Priority:        0,
	CorrelationId:   "",
	ReplyTo:         "",
	Expiration:      "",
	MessageId:       "",
	Timestamp:       time.Time{},
	Type:            "",
	UserId:          "",
	AppId:           "",
	ConsumerTag:     "",
	MessageCount:    0,
	DeliveryTag:     0,
	Redelivered:     false,
	Exchange:        "",
	RoutingKey:      "",
	Body:            message,
}

var rmq = queue.Rabbitmq{
	&amqp.Connection{},
	&amqp.Channel{},
	map[string]amqp.Queue{},
}

func TestCheckResultsFromWorker(t *testing.T) {
	res := CheckResultsFromWorker(&pFile)
	fmt.Println(res)

}

func TestGetMessageFromWorker(t *testing.T) {
	GetMessageFromWorker(m, &rmq)

}

func TestCreateErrorsInDB(t *testing.T) {
	CreateErrorsInDB(&pFile)
}
