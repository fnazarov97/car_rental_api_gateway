package clients

import (
	"car_rental/config"
	"car_rental/genprotos/authorization"
	"car_rental/genprotos/brand"
	"car_rental/genprotos/car"
	"car_rental/genprotos/rental"

	"google.golang.org/grpc"
)

type GrpcClients struct {
	Car           car.CarServiceClient
	Brand         brand.BrandServiceClient
	Rental        rental.RentalServiceClient
	Authorization authorization.AuthServiceClient
	conns         []*grpc.ClientConn
}

func NewGrpcClients(cfg config.Config) (*GrpcClients, error) {
	connCar, err := grpc.Dial(cfg.CarServiceGrpcHost+cfg.CarServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	car := car.NewCarServiceClient(connCar)

	connBrand, err := grpc.Dial(cfg.BrandServiceGrpcHost+cfg.BrandServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	brand := brand.NewBrandServiceClient(connBrand)

	connRental, err := grpc.Dial(cfg.RentalServiceGrpcHost+cfg.RentalServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	rental := rental.NewRentalServiceClient(connRental)

	connAuthorization, err := grpc.Dial(cfg.AuthorizationServiceGrpcHost+cfg.AuthorizationServiceGrpcPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	authorization := authorization.NewAuthServiceClient(connAuthorization)
	conns := make([]*grpc.ClientConn, 0)
	return &GrpcClients{
		Car:           car,
		Brand:         brand,
		Rental:        rental,
		Authorization: authorization,
		conns:         append(conns, connCar, connRental, connAuthorization),
	}, nil
}

func (c *GrpcClients) Close() {
	for _, v := range c.conns {
		v.Close()
	}
}
