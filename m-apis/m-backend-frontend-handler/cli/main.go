package main

import (
	"flag"
	"github.com/fvukojevic/m-apis/m-backend-frontend-handler/domain"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	grpcAddressHighScore := flag.String("address-m-highscore", "localhost:50051", "The gRPC server address for m-highscore microservice")
	grpcAddressGameEngine := flag.String("address-m-game-engine", "localhost:60051", "The gRPC server address for m-game-engine microservice")
}
