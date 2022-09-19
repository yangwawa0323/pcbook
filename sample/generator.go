package sample

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/jaswdr/faker"
	"github.com/yangwawa0323/pcbook/pb"
)

var fkr faker.Faker

func init() {
	fkr = faker.New()
}

func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

// NewCPU returns a new sample CPU
func NewCpu() *pb.Cpu {
	brand := randomCpuBrand()
	name := randomCpuName(brand)

	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)

	cpu := &pb.Cpu{
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
func NewGpu() *pb.Gpu {
	brand := randomGpuBrand()
	name := randomGpuName(brand)

	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)

	memory := &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}

	gpu := &pb.Gpu{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}

	return gpu
}

// NewRAM returns a sample RAM
func NewRAM() *pb.Memory {
	ram := &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}

	return ram
}

// NewSSD returns a sample SSD storage
func NewSSD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	return ssd
}

// NewHDD returns a sample SSD storage
func NewHDD() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1, 4)),
			Unit:  pb.Memory_TERABYTE,
		},
	}
	return hdd
}

func NewScreen() *pb.Screen {

	screen := &pb.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &pb.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCpu(),
		Ram:      NewRAM(),
		Gpus:     []*pb.Gpu{NewGpu()},
		Storages: []*pb.Storage{NewSSD(), NewHDD()},
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

func NewEmail() *pb.Email {
	return &pb.Email{
		Id:    randomID(),
		Email: fkr.Internet().Email(),
	}
}

func NewUser() (user *pb.User) {

	userID := randomID()
	user = &pb.User{
		Id:     userID,
		Name:   fkr.Person().Name(),
		Age:    uint32(randomInt(16, 35)),
		Emails: []*pb.Email{NewEmail(), NewEmail()},
	}
	return
}
