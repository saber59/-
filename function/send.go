package function

import (
	"log"

	"github.com/streadway/amqp"
)

var S = make(chan []byte, 100)

func Fsend() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	log.Print(err)
	inerrworke(500, err, "amqp连接出错")
	defer conn.Close()
	ch, err := conn.Channel()
	inerrworke(500, err, "队列创建出错")
	defer ch.Close()
	q, err := ch.QueueDeclare("upbase", true, false, false, false, nil)
	inerrworke(500, err, "队列创建出错2")
	for {
		body := <-S
		ch.Publish("", q.Name, false, false, amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	}

}
