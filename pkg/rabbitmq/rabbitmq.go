package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch, nil
}

/*
out -> É uma threading que vai ficar responsável por ler as msgs do rabbitmq
*/
func Consume(ch *amqp.Channel, out chan amqp.Delivery) error {
	msgs, err := ch.Consume(
		"orders",      //nome da fila
		"go-consumer", //nome da aplicação
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	//for infinito que encerra o trabalho nessa threading e envia para outra
	for msg := range msgs {
		out <- msg
	}
	return nil
}
