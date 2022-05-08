package main

import (
	"bufio"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
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

func (rmqClient *RabbitMQClient) CloseConnectionToRabbit() {
	rmqClient.Ch.Close()
	rmqClient.Conn.Close()
	fmt.Printf("connection to rabbit has been closed")
}

type Application struct {
	Name      string
	RMQClient RabbitMQClient
}

func (App *Application) CloseConnectionToRMQ() {
	App.RMQClient.CloseConnectionToRabbit()
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
			break
		}

	}
}

func (App *Application) SetupResources() {
	App.RMQClient.ConnectToRabbit()
	fmt.Println("resources has been set up")
}

func main() {
	RmqURL := "amqp://guest:guest@0.0.0.0:5672"
	RmqCLient := RabbitMQClient{RmqURL, nil, nil}

	app := Application{"myApp", RmqCLient}
	app.RunApp()
}
