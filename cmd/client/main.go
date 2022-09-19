package main

import (
	"context"
	"flag"
	"log"
	"math"
	"time"

	"github.com/yangwawa0323/pcbook/pb"
	"github.com/yangwawa0323/pcbook/sample"
	"github.com/yangwawa0323/pcbook/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func laptop_main() {
	start := time.Now()
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	out := utils.DebugOutput{}
	log.Print(out.Debug("dial server %s", *serverAddress))

	conn, err := grpc.Dial(*serverAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(out.Panic("cannot dial server: ", err))
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()
	laptop.Id = ""

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	// Request timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Print(out.Warn("laptop already exists"))
		} else {
			log.Fatal(out.Panic("cannot create laptop: ", err))
		}
		return
	}

	log.Print(out.Debug("created laptop with id: %s", res.Id))

	log.Print(out.Debug("========================"))

	var laptopId string = res.Id

	req2 := &pb.FindLaptopRequest{
		Id: laptopId,
	}
	res2, err := laptopClient.FindLaptop(ctx, req2)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			log.Print(out.Warn("laptop NotFound"))
		} else {
			log.Fatal(out.Panic("cannot found laptop: ", err))
		}
		return
	}

	log.Print(out.Debug("find laptop with id: %#v", res2.Laptop))
	log.Print(out.Warn("time elapsed: %.2f microseconds\n", math.Abs(
		float64(time.Until(start).Milliseconds()))))

}

func main() {
	start := time.Now()
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	out := utils.DebugOutput{}
	log.Print(out.Debug("dial server %s", *serverAddress))

	conn, err := grpc.Dial(*serverAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(out.Panic("cannot dial server: %v", err))
	}
	userClient := pb.NewUserServiceClient(conn)

	if err != nil {
		log.Fatal(out.Panic("cannot generate user uuid %v", err))
	}

	req := &pb.CreateUserRequest{
		User: sample.NewUser(),
	}

	log.Print(out.Debug("req : %#v", req.User))

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	res, err := userClient.CreateUser(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Print(out.Warn("user already exists"))
		} else {
			log.Fatal(out.Panic("cannot create user: ", err))
		}
		return
	}

	log.Print(out.Debug("created user with id: %s", res.Id))

	log.Print(out.Debug("========================"))

	log.Print(out.Warn("time elapsed: %.2f microseconds\n", math.Abs(
		float64(time.Until(start).Milliseconds()))))
}
