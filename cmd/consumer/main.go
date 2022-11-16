package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/farciomernandes/gointensivo/internal/ordem/infra/database"
	"github.com/farciomernandes/gointensivo/internal/ordem/usecase"
	"github.com/farciomernandes/gointensivo/pkg/rabbitmq"
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.CalculateFinalPriceUseCase{OrderRepository: repository}
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	//Depois que ele executar, ele fecha o canal
	defer ch.Close()

	/*
		Agora cria-se um canal de comunicação para conseguir falar com o rabbitmq e consumir as mensagens
	*/
	out := make(chan amqp.Delivery) // chanel
	go rabbitmq.Consume(ch, out)    //Consume na THREADING 2

	for msg := range out {
		var inputDTO usecase.OrderInputDTO

		err := json.Unmarshal(msg.Body, &inputDTO)
		//A msg do Body está sendo colocado dentro do inputDTO no seu formato

		if err != nil {
			panic(err)
		}

		inputDTO.Tax = 10.0

		outputDTO, err := uc.Execute(inputDTO)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Println(outputDTO) // Imprime na THREADING 1
	}

}
