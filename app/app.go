package app

import (
	"log"

	"github.com/charitan-go/auth-server/external/email"
	"github.com/charitan-go/auth-server/external/key"
	"github.com/charitan-go/auth-server/external/profile"
	"github.com/charitan-go/auth-server/internal/auth"
	"github.com/charitan-go/auth-server/pkg/database"
	"github.com/charitan-go/auth-server/rabbitmq"
	"github.com/charitan-go/auth-server/rest"
	"github.com/charitan-go/auth-server/rest/api"

	"go.uber.org/fx"
)

// Run both servers concurrently
func runServers(restSrv *rest.RestServer, rabbitmqSrv *rabbitmq.RabbitmqServer) {
	log.Println("In invoke")

	// Start REST server
	go func() {
		log.Println("In goroutine of rest")
		restSrv.Run()
	}()

	// Start RabbitMQ server
	go func() {
		log.Println("In goroutine of rabbitmq")
		rabbitmqSrv.Run()
	}()

}

func Run() {
	// Connect to db
	database.SetupDatabase()

	fx.New(
		auth.AuthModule,
		profile.ProfileModule,
		key.KeyModule,
		email.EmailModule,
		rabbitmq.RabbitmqModule,
		fx.Provide(
			rest.NewRestServer,
			rest.NewEcho,
			api.NewApi,
		),
		fx.Invoke(runServers),
	).Run()
}
