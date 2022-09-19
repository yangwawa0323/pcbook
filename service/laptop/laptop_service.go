package laptop

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/yangwawa0323/pcbook/pb"
	"github.com/yangwawa0323/pcbook/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	Store LaptopStore
	pb.UnimplementedLaptopServiceServer
}

// NewLaptopServer return a new LaptopServer
func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{Store: store}
}

// CreateLaptop is a unary RPC to create a new laptop
func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {

	out := utils.DebugOutput{}

	laptop := req.GetLaptop()
	log.Print(out.Debug("receive a create-laptop request with id: %s", laptop.Id))

	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop ID: %v", err)
		}
		laptop.Id = id.String()
	}

	// some heavy processing
	// time.Sleep(6 * time.Second)
	// save the laptop to in-memory or db store
	err := server.Store.Save(context.Background(), laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}

		return nil, status.Errorf(code, "cannot save laptop to the store: %v", err)
	}

	log.Print(out.Debug("saved laptop with id %s", laptop.Id))
	response := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}

	return response, nil
}

func (server *LaptopServer) FindLaptop(
	ctx context.Context,
	req *pb.FindLaptopRequest) (*pb.FindLaptopResponse, error) {
	laptopId := req.Id
	laptop, err := server.Store.Find(ctx, laptopId)
	if err != nil {
		code := codes.NotFound
		return nil, status.Errorf(code, "laptop with id %s is not found: %v", req.Id, err)
	}
	return &pb.FindLaptopResponse{
		Laptop: laptop,
	}, nil
}
