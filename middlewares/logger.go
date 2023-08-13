package middlewares

import (
	"log"
	"time"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcLogger(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,

) (resp interface{}, err error) {
	startTime   := time.Now()
	result, err := handler(ctx, req)
	duration    := time.Since(startTime)

	statusCode := codes.Unknown
	if st, ok := status.FromError(err); ok {
		statusCode = st.Code()
	}

	log.Printf(
		"protocol=grpc method=%s status_code=%d status_text=%v duration=%d", 
		info.FullMethod, int(statusCode), statusCode.String(), duration,
	)

	return result, err
}