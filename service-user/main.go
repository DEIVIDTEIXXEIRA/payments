package serviceuser

import (
	"log"

	"github.com/payments/service-user/src/config"
	"github.com/payments/service-user/src/repository"
	"github.com/payments/service-user/src/server"
	"github.com/payments/service-user/src/service"
)

func main() {
	config.LoadConfig("development")
	log.Println("Create new httpServer")
	http := server.NewServer()
	conn, _ := repository.OpenConnection()

	serv := server.NewServerList(service.NewService(repository.NewRepository(conn)))

	http.NewRoutes(serv)
	http.StartAPI()
}
