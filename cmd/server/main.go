package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/yangwawa0323/pcbook/pb"
	"github.com/yangwawa0323/pcbook/service/laptop"
	"github.com/yangwawa0323/pcbook/service/user"
	"github.com/yangwawa0323/pcbook/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func laptop_main() {

	out := &utils.DebugOutput{}

	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Print(out.Debug("Start server on port %d\n", *port))

	// laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())

	dbLaptopStore := laptop.NewDbLaptopStore()

	// log.Printf(out.Debug("%v", dbLaptopStore))
	laptopServer := laptop.NewLaptopServer(dbLaptopStore)

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	// _, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(out.Error("cannot start server:", err))
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(out.Error("cannot start server"))
	}
}

func main() {
	out := utils.NewDebugOutput()

	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Print(out.Debug("Start server on port %d\n", *port))

	dbUserStore := user.NewUserDBStore()

	userServer := user.NewUserServer(dbUserStore)
	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, userServer)

	reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	// _, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(out.Error("cannot start server:", err))
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(out.Error("cannot start server"))
	}

}
