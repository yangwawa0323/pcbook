package user

import (
	"context"
	"fmt"
	"log"
	"os"

	pb_user "github.com/yangwawa0323/pcbook/pb/user/v1"
)

type UserServer struct {
	Store UserStore
	pb_user.UnimplementedUserServiceServer
}

func NewUserServer(store UserStore) *UserServer {
	return &UserServer{Store: store}
}

func (us *UserServer) CreateUser(ctx context.Context,
	req *pb_user.CreateUserRequest) (*pb_user.CreateUserResponse, error) {
	user, err := req.GetUser().ToORM(ctx)
	log.Print(out.Debug("Get user request: %#v", user))
	if err != nil {
		return nil, err
	}
	if err = us.Store.Save(&user); err != nil {
		log.Fatal(out.Panic("Cannot save the user to DB : %#v", err))
	}

	fmt.Fprint(os.Stdout, out.Debug("user is created : %#v\n", user))

	return &pb_user.CreateUserResponse{
		Id: user.Id,
	}, nil
}

func (us *UserServer) FindUser(ctx context.Context,
	req *pb_user.FindUserRequest) (*pb_user.FindUserResponse, error) {

	userID := req.GetId()
	userORM, err := us.Store.Find(userID)
	if err != nil {
		return nil, err
	}
	user, _ := userORM.ToPB(ctx)

	return &pb_user.FindUserResponse{
		User: &user,
	}, nil
}
