package main

import (
	"context"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Msg1 struct {
	Id   int    `gorm:"id"`   //gorm:"primaryKey"`
	Data string `gorm:"data"` // gorm:"type:varchar(256)"`
	Age  string `gorm:"age" ` //gorm:"type:int(32)"`
}
type Msg struct {
	Id   int    `gorm:"id"`   // gorm:"primaryKey"`
	Data string `gorm:"data"` // gorm:"type:varchar(256)"`
	Age  string `gorm:"age"`  // gorm:"type:int(32)"`
}

var msg1 Msg1

func main() {
	//Enter the message

	msg := Msg{
		Id:   0,
		Data: "vamsi",
		Age:  "20",
	}
	a, _ := json.Marshal(msg)
	_ = json.Unmarshal(a, &msg1)
	//fmt.Println("Enter the message:")
	//fmt.Scanln(&msg)
	//fmt.Println("Enter the age:")
	//fmt.Scanln(&age)
	url := "amqp://guest:guest@localhost:5672/"
	//Establishing connection to RMQ
	conn, err := amqp.Dial(url)
	defer conn.Close()
	//Creating a channel
	ch, err := conn.Channel()
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"MessageQ", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments

	)
	if err != nil {
		fmt.Println("Failed to declare a queue")
		return
	}
	fmt.Println(msg1.Age)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//var k Msg
	body := a
	//fmt.Println(k)

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
			//Body:       []byte(body1),
		})
	if err != nil {
		fmt.Println("Failed to publish a message")
		return
	}
	log.Printf(" [x] Sent %s\n", body)
	//	log.Printf(" [x] Sent %s\n", body1)
	//log.Printf(" [x] Sent %s\n", body1)

	// Subscribing to MessageQ for getting messages.
	messagess, err := ch.Consume(
		"MessageQ", // queue name
		"",         // consumer
		true,       // auto-ack
		//false, // auto-ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   // arguments
	)

	if err != nil {
		log.Println(err)
	}
	// Make a channel to receive messages.
	foreve := make(chan bool)
	DB := Init()
	go func() {
		for message := range messagess {
			// For example, show received message in a console.
			log.Printf(" > consumer Received message: %s\n", message.Body)
			//log.Printf(" > consumer Received message: %s\n", message.Body1)
			var m, z Msg
			//m.Data = string(message.Body)
			m.Data = string(message.Body)
			m.Age = string(message.Body)
			//fmt.Println(m)
			fmt.Println("hereeeeee.")
			// insert to the  table
			z.Data = m.Data
			z.Age = m.Age

			if result := DB.Create(&msg1); result.Error != nil {
				fmt.Println("vvvv")
			}
		}

		<-foreve

	}()

}
