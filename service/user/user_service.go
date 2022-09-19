package user

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/yangwawa0323/pcbook/pb"
)

type UserServer struct {
	Store UserStore
	pb.UnimplementedUserServiceServer
}

func NewUserServer(store UserStore) *UserServer {
	return &UserServer{Store: store}
}

func (us *UserServer) CreateUser(ctx context.Context,
	req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := req.GetUser().ToORM(ctx)
	log.Print(out.Debug("Get user request: %#v", user))
	if err != nil {
		return nil, err
	}
	if err = us.Store.Save(&user); err != nil {
		log.Fatal(out.Panic("Cannot save the user to DB : %#v", err))
	}

	fmt.Fprint(os.Stdout, out.Debug("user is created : %#v\n", user))

	return &pb.CreateUserResponse{
		Id: user.Id,
	}, nil
}

func (us *UserServer) FindUser(ctx context.Context,
	req *pb.FindUserRequest) (*pb.FindUserResponse, error) {

	userID := req.GetId()
	userORM, err := us.Store.Find(userID)
	if err != nil {
		return nil, err
	}
	user, _ := userORM.ToPB(ctx)

	return &pb.FindUserResponse{
		User: &user,
	}, nil
}
