package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/kioku-project/kioku/services/frontend/handler"
	pb "github.com/kioku-project/kioku/services/frontend/proto"
	pblogin "github.com/kioku-project/kioku/services/login/proto"
	pbregister "github.com/kioku-project/kioku/services/register/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"

	_ "github.com/go-micro/plugins/v4/registry/kubernetes"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
)

var (
	service        = "frontend"
	version        = "latest"
	serviceAddress = fmt.Sprintf("%s%s", os.Getenv("HOSTNAME"), ":8080")
)

func main() {

	logger.Info("Trying to listen on: ", serviceAddress)

	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer(server.Address(serviceAddress))),
		micro.Client(grpcc.NewClient()),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
		micro.Address(serviceAddress),
	)

	// Create a new instance of the service handler with the initialized database connection
	svc := handler.New(
		pblogin.NewLoginService("login", srv.Client()),
		pbregister.NewRegisterService("register", srv.Client()),
	)

	app := fiber.New()
	app.Post("/api/login", svc.LoginHandler)
	app.Post("/api/register", svc.RegisterHandler)

	// Register handler
	if err := pb.RegisterHealthHandler(srv.Server(), new(handler.Health)); err != nil {
		logger.Fatal(err)
	}

	go func() {
		if err := app.Listen(":80"); err != nil {
			logger.Fatal(err)
		}
	}()

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
