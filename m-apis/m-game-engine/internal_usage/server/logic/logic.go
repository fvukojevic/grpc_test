package logic

import (
	"google.golang.org/grpc"
)

var pastFourScores = []float64{5.0, 4.0, 2.0, 1.0}

type Grpc struct {
	address string //address where the gRPC will listen at
	srv     *grpc.Server
}

func GetSize() float64 {
	oldScores := pastFourScores[0] + pastFourScores[1]
	newScores := pastFourScores[2] + pastFourScores[3]

	diff := newScores - oldScores

	if diff > 0.0 {
		size := 600.0 + diff*60.0
		if size < 2000.0 {
			return size
		}
		return 2000.0
	}

	if diff > -5.0 && diff <= 0.0 {
		return 100.0 + 18.0*diff
	}

	return 10.0
}

func SetScore(x float64) {
	pastFourScores = append(pastFourScores, x)
	pastFourScores = pastFourScores[1:]
}
