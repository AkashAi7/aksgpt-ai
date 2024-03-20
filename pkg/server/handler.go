package server

import (
	rpc "buf.build/gen/go/aksgpt-ai/aksgpt/grpc/go/schema/v1/schemav1grpc"
)

type handler struct {
	rpc.UnimplementedServerServiceServer
}
