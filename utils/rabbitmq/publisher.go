package rabbitmq

// import (
// 	"fmt"
// 	"greebel.core.be/config"
// 	"log"

// 	"github.com/streadway/amqp"
// )

// func Publish(body []byte, contentType, publishName string) error {
// 	conn, _ := config.GetRabbitmqConnection()
// 	defer conn.Close()

// 	ch, err := conn.Channel()
// 	if err != nil {
// 		log.Fatalf("cannot create channel: %v", err)
// 	}
// 	defer ch.Close()

// 	q, err := ch.QueueDeclare(
// 		publishName, // name
// 		false,       // durable
// 		false,       // delete when unused
// 		false,       // exclusive
// 		false,       // no-wait
// 		nil,         // arguments
// 	)
// 	if err != nil {
// 		fmt.Println(err, "Failed to declare a queue")
// 	}

// 	// defer cancel()
// 	err = ch.Publish(
// 		"",     // exchange
// 		q.Name, // routing key
// 		false,  // mandatory
// 		false,  // immediate
// 		amqp.Publishing{
// 			ContentType: contentType,
// 			Body:        []byte(body),
// 		})
// 	if err != nil {
// 		fmt.Println(err, "Failed to publish a message")
// 	}
// 	log.Printf(" [x] Sent %s\n", body)

// 	return err
// }
