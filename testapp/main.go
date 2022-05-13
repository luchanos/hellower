package main

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"runtime"
	"time"
)

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

	fmt.Println("connection to rabbit has been created")
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

func (rmqClient *RabbitMQClient) ConsumeMessages(queue string, ctx context.Context) {
	q, err := rmqClient.Ch.QueueDeclare(queue, false, false, false, false, nil)
	failOnError(err, "error due to setup consumer")
	msgs, err := rmqClient.Ch.Consume(q.Name, "", true, false, false, false, nil)

	go func() {
		runtime.Gosched()
		for d := range msgs {
			select {
			case <-ctx.Done():
				rmqClient.Ch.Close()
				fmt.Println("Consumer stopped!")
				return
			default:
				log.Printf(" [x] %s", d.Body)
				runtime.Gosched()
			}
		}
	}()
	log.Printf(" [*] waiting for logs")
}

func StartSender(param string, rmqClient *RabbitMQClient, ctx context.Context) http.HandlerFunc {
	f := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(param)
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Fprintln(w, "RMQ sender had been finished!!!", r.URL.String())
					return
				default:
					rmqClient.SendMessage("hello", "тестовое сообщение")
				}
			}
		}()
		fmt.Fprintln(w, "RMQ sender had been start!!!", r.URL.String())
	}
	return http.HandlerFunc(f)
}

func StartConsumer(param string, rmqClient *RabbitMQClient, ctx context.Context) http.HandlerFunc {
	f := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(param)
		rmqClient.Ch, _ = rmqClient.Conn.Channel()
		rmqClient.ConsumeMessages("hello", ctx)
		fmt.Fprintln(w, "RMQ consumer had been start!!!", r.URL.String())
	}
	return http.HandlerFunc(f)
}

func StopSender(param string, rmqClient *RabbitMQClient, finish context.CancelFunc) http.HandlerFunc {
	f := func(w http.ResponseWriter, r *http.Request) {
		finish()
		fmt.Fprintln(w, "RMQ sender had been stopped!!!", r.URL.String())
	}
	return http.HandlerFunc(f)
}

func StopConsumer(param string, rmqClient *RabbitMQClient, finish context.CancelFunc) http.HandlerFunc {
	f := func(w http.ResponseWriter, r *http.Request) {
		finish()
		fmt.Fprintln(w, "RMQ consumer had been stopped!!!", r.URL.String())
	}
	return http.HandlerFunc(f)
}

func makeServer(addr string, rmqSender, rmqConsumer *RabbitMQClient) *http.ServeMux {
	mux := http.NewServeMux()

	senderFinishCtx, finishSender := context.WithCancel(context.Background())
	consumerFinishCtx, finishConsumer := context.WithCancel(context.Background())
	mux.HandleFunc("/start_sender", StartSender("try sender start", rmqSender, senderFinishCtx))
	mux.HandleFunc("/start_consumer", StartConsumer("try consumer start", rmqConsumer, consumerFinishCtx))
	mux.HandleFunc("/stop_sender", StopSender("try sender stop", rmqSender, finishSender))
	mux.HandleFunc("/stop_consumer", StopConsumer("try consumer stop", rmqConsumer, finishConsumer))
	return mux
}

func runServer(addr string, rmqSender, rmqConsumer *RabbitMQClient) {
	mux := makeServer(addr, rmqSender, rmqConsumer)

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	fmt.Println("starting a new server at", addr)
	server.ListenAndServe()
}

func main() {
	url := "amqp://guest:guest@0.0.0.0:5672"
	rmqSender := RabbitMQClient{URL: url, Conn: nil, Ch: nil}
	rmqConsumer := RabbitMQClient{URL: url, Conn: nil, Ch: nil}
	rmqConsumer.SetupRabbitMQClient()
	rmqSender.SetupRabbitMQClient()

	go runServer(":8081", &rmqSender, &rmqConsumer)
	fmt.Scanln()
}
