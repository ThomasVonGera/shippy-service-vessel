package main

import (
	"context"
	"errors"
	"log"

	pb "github.com/ThomasVonGera/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
)

const ServiceName = "shippy.service.vessel"

type Repository interface {
	FindAvaiable(*pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

func (repo *VesselRepository) FindAvaiable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("Kein Vessel mit dieser Specification gefunden")
}

type vesselService struct {
	repo Repository
}

func (s *vesselService) FindAvaiable(ctx context.Context, request *pb.Specification, response *pb.Response) error {

	vessel, err := s.repo.FindAvaiable(request)
	if err != nil {
		return err
	}

	response.Vessel = vessel
	return nil
}

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatFace", MaxWeight: 200000, Capacity: 500},
	}
	repo := &VesselRepository{vessels}

	service := micro.NewService(
		micro.Name(ServiceName),
	)

	service.Init()

	if err := pb.RegisterVesselServiceHandler(service.Server(), &vesselService{repo}); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
