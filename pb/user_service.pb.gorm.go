package pb

import (
	context "context"
	gorm "github.com/jinzhu/gorm"
)

type UserServiceDefaultServer struct {
	DB *gorm.DB
}

// CreateUser ...
func (m *UserServiceDefaultServer) CreateUser(ctx context.Context, in *CreateUserRequest) (*CreateUserResponse, error) {
	out := &CreateUserResponse{}
	return out, nil
}

// FindUser ...
func (m *UserServiceDefaultServer) FindUser(ctx context.Context, in *FindUserRequest) (*FindUserResponse, error) {
	out := &FindUserResponse{}
	return out, nil
}
