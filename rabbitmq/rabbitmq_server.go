package rabbitmq

import (
	"fmt"
	"log"
	"os"

	auth "github.com/charitan-go/auth-server/internal/auth/service"
	rabbitmqservice "github.com/charitan-go/auth-server/rabbitmq/service"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitmqServer struct {
	rabbitmqSvc rabbitmqservice.RabbitmqService
	authSvc     auth.AuthService
}

func NewRabbitmqServer(rabbitmqSvc rabbitmqservice.RabbitmqService, authSvc auth.AuthService) *RabbitmqServer {
	return &RabbitmqServer{rabbitmqSvc, authSvc}
}

func (srv *RabbitmqServer) startRabbitmqConsumer() error {
	// ch, err := srv.rabbitmqSvc.ConnectRabbitmq()
	log.Println("In function startRabbitmqConsumer")

	// time.Sleep(5 * time.Second)

	amqpConnectionStr := fmt.Sprintf("amqp://%s:%s@message-broker:5672",
		os.Getenv("MESSAGE_BROKER_USER"),
		os.Getenv("MESSAGE_BROKER_PASSWORD"))
	log.Printf("Connection str: %s\n", amqpConnectionStr)
	conn, err := amqp.Dial(amqpConnectionStr)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return err
	}
	defer ch.Close()

	msgs, err := srv.setupGetPrivateKeyConsumer(ch)
	if err != nil {
		log.Fatalf("Setup get private key consumer failed %v\n", err)
		return err
	}

	forever := make(chan bool)
	go func() {
		log.Println("Inside the loop to process exchange topics")
		for d := range msgs {
			switch d.Exchange {
			case KEY_EXCHANGE_NAME:
				{
					log.Printf("Received message from exchange GET_PRIVATE_KEY: %s\n", d.Body)
					srv.authSvc.HandleGetPrivateKeyRabbitmq()
				}
			}
		}
	}()

	<-forever

	return nil

}

func (s *RabbitmqServer) Run() {
	// Start rabbitmq consumer
	s.startRabbitmqConsumer()
}
