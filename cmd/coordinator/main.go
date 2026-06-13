package main

import (
	"log"
	"net"

	grpcpkg "github.com/ShriChinmay/Crucible/coordinator"
	"github.com/ShriChinmay/Crucible/coordinator/proto"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf(
			"failed to listen: %v",
			err,
		)
	}

	server := grpc.NewServer()

	proto.RegisterFleetCoordinatorServer(
		server,
		&grpcpkg.FleetCoordinatorServer{},
	)

	log.Println(
		"Fleet Coordinator listening on :50051",
	)

	if err := server.Serve(lis); err != nil {

		log.Fatalf(
			"failed to serve: %v",
			err,
		)
	}
}