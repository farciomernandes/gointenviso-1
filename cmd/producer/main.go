package main

import (
	"encoding/json"
	"math/rand"

	"github.com/farciomernandes/gointensivo/internal/ordem/entity"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Pega uma order e trata ela para publicar no rabbitmq
func Publish(ch *amqp.Channel, order entity.Order) error {
	//Transforma os dados da order em um json
	body, err := json.Marshal(order)

	if err != nil {
		return err
	}

	err = ch.Publish(
		"amq.direct",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// Gera um order random
func GenerateOrders() entity.Order {
	return entity.Order{
		ID:    uuid.New().String(),
		Price: rand.Float64() * 100,
		Tax:   rand.Float64() * 10,
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	for i := 0; i < 100; i++ {
		Publish(ch, GenerateOrders())
	}

}
