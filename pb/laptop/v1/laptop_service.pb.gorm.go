package pb

import (
	context "context"
	gorm "github.com/jinzhu/gorm"
)

type LaptopServiceDefaultServer struct {
	DB *gorm.DB
}

// CreateLaptop ...
func (m *LaptopServiceDefaultServer) CreateLaptop(ctx context.Context, in *CreateLaptopRequest) (*CreateLaptopResponse, error) {
	out := &CreateLaptopResponse{}
	return out, nil
}

// FindLaptop ...
func (m *LaptopServiceDefaultServer) FindLaptop(ctx context.Context, in *FindLaptopRequest) (*FindLaptopResponse, error) {
	out := &FindLaptopResponse{}
	return out, nil
}
