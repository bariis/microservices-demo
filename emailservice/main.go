package main

import (
	"log"

	rabbitmq "github.com/bariis/rabbitmq-fanout-library"
)

const (
	AMQPUri      = "amqp://guest:guest@rabbitmq:5672/"
	ExchangeName = "CartItems"
	ConsumerName = "PaymentService"
	QueueName    = "PaymentQueue"
	ContentType  = "application/json"
)

type Checkout struct {
	Email      string
	CreditCart CreditCart
	Address    Address
	Paid       float64
}

type Address struct {
	StreetAddress string
	City          string
	State         string
	ZipCode       string
	Country       string
}

type CreditCart struct {
	Number          string
	ExpirationMonth int
	ExpirationYear  int
	Cvv             int
}

func main() {

	rmq := rabbitmq.New(AMQPUri)
	if err := rmq.Connect(); err != nil {
		log.Fatalln(err)
	}

	channel, err := rmq.Channel()
	if err != nil {
		log.Fatalln(err)
	}

	amqp := rabbitmq.NewAMQP(rmq, ExchangeName, QueueName, "text/plain")
	err = amqp.StartQueue(channel)
	if err != nil {
		log.Fatalln(err)
	}

	forever := make(chan bool)
	go func() {
		items, err := amqp.Consume(channel, ConsumerName)
		log.Println(err)
		for _ = range items {
			log.Println("Email is sent.")
		}
	}()
	<-forever

}
