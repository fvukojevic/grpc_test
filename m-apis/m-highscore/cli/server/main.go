package main

import (
	"flag"
	grpcSetup "github.com/fvukojevic/grpc_test/m-apis/m-highscore/internal/server/grpc"
	"github.com/rs/zerolog/log"
)

func main() {
	var addressPtr = flag.String("address", ":50051", "address where you can connect to gRPC m-highscore service")

	flag.Parse()

	s := grpcSetup.NewServer(*addressPtr)

	if err := s.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start grpc server")
	}
}