package sample

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/jaswdr/faker"
	pb_laptop "github.com/yangwawa0323/pcbook/pb/laptop/v1"
	pb_user "github.com/yangwawa0323/pcbook/pb/user/v1"
)

var fkr faker.Faker

func init() {
	fkr = faker.New()
}

func NewKeyboard() *pb_laptop.Keyboard {
	keyboard := &pb_laptop.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

// NewCPU returns a new sample CPU
func NewCpu() *pb_laptop.Cpu {
	brand := randomCpuBrand()
	name := randomCpuName(brand)

	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)

	cpu := &pb_laptop.Cpu{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}

	return cpu
}

// NewGPU return a sample GPU
func NewGpu() *pb_laptop.Gpu {
	brand := randomGpuBrand()
	name := randomGpuName(brand)

	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)

	memory := &pb_laptop.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  pb_laptop.Memory_GIGABYTE,
	}

	gpu := &pb_laptop.Gpu{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}

	return gpu
}

// NewRAM returns a sample RAM
func NewRAM() *pb_laptop.Memory {
	ram := &pb_laptop.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  pb_laptop.Memory_GIGABYTE,
	}

	return ram
}

// NewSSD returns a sample SSD storage
func NewSSD() *pb_laptop.Storage {
	ssd := &pb_laptop.Storage{
		Driver: pb_laptop.Storage_SSD,
		Memory: &pb_laptop.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit:  pb_laptop.Memory_GIGABYTE,
		},
	}
	return ssd
}

// NewHDD returns a sample SSD storage
func NewHDD() *pb_laptop.Storage {
	hdd := &pb_laptop.Storage{
		Driver: pb_laptop.Storage_HDD,
		Memory: &pb_laptop.Memory{
			Value: uint64(randomInt(1, 4)),
			Unit:  pb_laptop.Memory_TERABYTE,
		},
	}
	return hdd
}

func NewScreen() *pb_laptop.Screen {

	screen := &pb_laptop.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

func NewLaptop() *pb_laptop.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &pb_laptop.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCpu(),
		Ram:      NewRAM(),
		Gpus:     []*pb_laptop.Gpu{NewGpu()},
		Storages: []*pb_laptop.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		// Weight: &pb.Laptop_WeightKg{
		// 	WeightKg: randomFloat64(1.0, 3.0),
		// },
		PriceUsd:    randomFloat64(1500, 3000),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdatedAt:   ptypes.TimestampNow(),
	}

	return laptop
}

func NewEmail() *pb_user.Email {
	return &pb_user.Email{
		Id:    randomID(),
		Email: fkr.Internet().Email(),
	}
}

func NewUser() (user *pb_user.User) {

	userID := randomID()
	user = &pb_user.User{
		Id:     userID,
		Name:   fkr.Person().Name(),
		Age:    uint32(randomInt(16, 35)),
		Emails: []*pb_user.Email{NewEmail(), NewEmail()},
	}
	return
}
