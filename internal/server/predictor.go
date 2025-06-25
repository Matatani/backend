package server

import (
	"context"
	ml2 "www.github.com/Maevlava/Matatani/backend/internal/ml"
)

type Predictor struct {
	ml2.UnimplementedPredictorServer
}

func (p Predictor) Hello(ctx context.Context, request *ml2.HelloRequest) (*ml2.HelloResponse, error) {
	greet := request.Greeting.Greeting + request.From + "!"
	return &ml2.HelloResponse{
		Greeting: greet,
	}, nil
}
