package main

import (
	"bufio"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"runtime"
	"strings"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s %s", err, msg)
	}
}

type RabbitMQClient struct {
	RmqUrl string
	Conn   *amqp.Connection
	Ch     *amqp.Channel
}

type RabbitMQConsumer struct {
	RmqUrl string
	Conn   *amqp.Connection
	Ch     *amqp.Channel
}

func (rmqClient *RabbitMQClient) ConnectToRabbit() {
	conn, err := amqp.Dial(rmqClient.RmqUrl)
	FailOnError(err, "failed to connect to RabbitMQ")
	rmqClient.Conn = conn

	ch, err := conn.Channel()
	FailOnError(err, "failed to open a channel")
	rmqClient.Ch = ch

	fmt.Printf("connection to rabbit has been created")

	err = ch.ExchangeDeclare("logs", "fanout", true, false, false, false, nil)
	FailOnError(err, "failed to create an exchange")
}

func (rmqConsumer *RabbitMQConsumer) ConnectToRabbit() {
	conn, err := amqp.Dial(rmqConsumer.RmqUrl)
	FailOnError(err, "failed to connect to RabbitMQ")
	rmqConsumer.Conn = conn

	ch, err := conn.Channel()
	FailOnError(err, "failed to open a channel")
	rmqConsumer.Ch = ch

	fmt.Printf("connection to rabbit has been created")

	err = ch.ExchangeDeclare("logs", "fanout", true, false, false, false, nil)
	FailOnError(err, "failed to create an exchange")
}

func (rmqConsumer *RabbitMQConsumer) ConsumeFromRabbit() {
	rmqConsumer.ConnectToRabbit()
	q, err := rmqConsumer.Ch.QueueDeclare("", false, false, true, false, nil)
	FailOnError(err, "error due to setup consumer")
	err = rmqConsumer.Ch.QueueBind(q.Name, "", "logs", false, nil)
	FailOnError(err, "failed to bind a queue")

	msgs, err := rmqConsumer.Ch.Consume(q.Name, "", true, false, false, false, nil)

	forever := make(chan bool)

	go func() {
		runtime.Gosched()
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
			runtime.Gosched()
		}
	}()

	log.Printf(" [*] waiting for logs")
	<-forever
}

func (rmqClient *RabbitMQClient) PublishMessageToRabbit() {
	err := rmqClient.Ch.Publish("logs", "", false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte{1, 2, 3}})
	FailOnError(err, "unable to send message to rabbit!")
	fmt.Println("SENT!")
}

func (rmqClient *RabbitMQClient) CloseConnectionToRabbit() {
	rmqClient.Ch.Close()
	rmqClient.Conn.Close()
	rmqClient.Ch = nil
	rmqClient.Conn = nil
	fmt.Printf("connection to rabbit has been closed")
}

type Application struct {
	Name        string
	RMQClient   RabbitMQClient
	RMQConsumer RabbitMQConsumer
}

func (App *Application) CloseConnectionToRMQ() {
	App.RMQClient.CloseConnectionToRabbit()
}

func (RabbitMQConsumer) Run() {

}

func (App *Application) RunApp() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("1 - run rabbitMQ, 2 - stop rabbit")
		answer, err := reader.ReadString('\n')
		answer = strings.Trim(answer, "\n")
		FailOnError(err, "unable to read")
		if answer == "1" {
			App.RMQClient.ConnectToRabbit()
		} else if answer == "2" {
			App.RMQClient.CloseConnectionToRabbit()
		} else if answer == "3" {
			if (App.RMQClient.Ch != nil) && (App.RMQClient.Conn != nil) {
				App.RMQClient.PublishMessageToRabbit()
			} else {
				fmt.Println("Channel and conn is not created!")
			}
		} else if answer == "4" {
			App.RMQConsumer.ConsumeFromRabbit()
		}

	}
}

func (App *Application) SetupResources() {
	App.RMQClient.ConnectToRabbit()
	App.RMQConsumer.ConsumeFromRabbit()
	fmt.Println("resources has been set up")
}

func main() {
	RmqURL := "amqp://guest:guest@0.0.0.0:5672"
	RmqCLient := RabbitMQClient{RmqURL, nil, nil}
	RmqConsumer := RabbitMQConsumer{RmqURL, nil, nil}

	app := Application{"myApp", RmqCLient, RmqConsumer}
	app.RunApp()
}
