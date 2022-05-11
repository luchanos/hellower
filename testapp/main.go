package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"runtime"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	myParam := r.URL.Query().Get("param")
	if myParam != "" {
		fmt.Fprintln(w, "myParam is ", myParam)
	}

	key := r.FormValue("key")
	if key != "" {
		fmt.Fprintln(w, "key is", key)
	}
}

func rabbitSendMsgHandler(w http.ResponseWriter, r *http.Request) {
	// надо прокинуть канал, через который будут общаться горутины
	fmt.Fprintln(w, "Success!")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type RabbitMQClient struct {
	URL  string
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

func (rmqClient *RabbitMQClient) SetupRabbitMQClient() {
	conn, err := amqp.Dial(rmqClient.URL)
	failOnError(err, "failed to connect to RabbitMQ")
	rmqClient.Conn = conn

	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	rmqClient.Ch = ch

	fmt.Printf("connection to rabbit has been created")
}

func (rmqClient *RabbitMQClient) SendMessage(queue, msg string) {
	// объявляем очередь
	q, err := rmqClient.Ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = rmqClient.Ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	time.Sleep(1 * time.Second)
	failOnError(err, "Failed to publish a message")
	runtime.Gosched()

}

func (rmqClient *RabbitMQClient) ConsumeMessages(queue string) {
	q, err := rmqClient.Ch.QueueDeclare(queue, false, false, false, false, nil)
	failOnError(err, "error due to setup consumer")
	//err = rmqClient.Ch.QueueBind(q.Name, "", "", false, nil)
	//failOnError(err, "failed to bind a queue")

	msgs, err := rmqClient.Ch.Consume(q.Name, "", true, false, false, false, nil)

	go func() {
		runtime.Gosched()
		for d := range msgs {
			time.Sleep(1 * time.Second)
			log.Printf(" [x] %s", d.Body)
			runtime.Gosched()
		}
	}()

	log.Printf(" [*] waiting for logs")
}

func main() {
	url := "amqp://guest:guest@0.0.0.0:5672"
	rmqSender := RabbitMQClient{URL: url, Conn: nil, Ch: nil}
	rmqSender.SetupRabbitMQClient()

	rmqConsumer := RabbitMQClient{URL: url, Conn: nil, Ch: nil}
	rmqConsumer.SetupRabbitMQClient()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("1")
			runtime.Gosched()
		}
	}()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("2")
			runtime.Gosched()
		}
	}()

	go func() {
		for {
			rmqSender.SendMessage("hello", "тестовое сообщение")
		}
	}()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			rmqConsumer.ConsumeMessages("hello")
		}
	}()
	fmt.Scanln()
}
