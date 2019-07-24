package main

import (
	"net"

	serviceHandler "github.com/chrusty/go-grpc-service/internal/handler"
	storage "github.com/chrusty/go-grpc-service/internal/storage"
	storageMocks "github.com/chrusty/go-grpc-service/internal/storage/mocks"
	serviceProto "github.com/chrusty/go-grpc-service/pkg/api/v1/go"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	listenAddress = ":6666"
)

// main() is where we put all our packages together to form the service binary:
func main() {

	// Prepare some dependencies:
	logger := logrus.New()
	storer := new(storageMocks.FakeStorer)

	// Program the storer mock to respond with _something_:
	storer.CreateCruftReturns("12345", nil)
	storer.ReadCruftReturns(nil, storage.ErrNotFound)

	// Inject the dependencies into a new Handler:
	handler := serviceHandler.New(logger, storer)

	// Make a new GRPC Server (usually I would have this in a common / shared library, and pre-load it with middleware built from our logger / instrumenter / tracer interfaces):
	grpcServer := grpc.NewServer()

	// Register our Handler and GRPC Server with our generated service-proto code:
	serviceProto.RegisterExampleServer(grpcServer, handler)

	// Listen for connections:
	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		logger.Fatalf("Unable to start GRPC server on TCP address %s", listenAddress)
	}

	// Start the GRPC server:
	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatalf("Unable to start the GRPC server: %v", err)
	}
}
