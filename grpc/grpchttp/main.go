package main

import (
    "context"
    "fmt"
    "github.com/grpc-ecosystem/grpc-gateway/runtime"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    helloworldpb "grpc2/proto"
    "log"
    "net"
    "net/http"
)


const (
 port = ":50051"
)

type server struct {
    helloworldpb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
 fmt.Printf("Received: %v", in.GetName())
 return &helloworldpb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main () {

    // Create a listener on TCP port
    lis, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalln("Failed to listen:", err)
    }
    // Create a gRPC server object
    s := grpc.NewServer()
    // Attach the Greeter service to the server
    helloworldpb.RegisterGreeterServer(s, &server{})
    // Serve gRPC server
    log.Println("Serving gRPC on 0.0.0.0:8080")
    go func() {
        log.Fatalln(s.Serve(lis))
    }()

    // Create a client connection to the gRPC server we just started
    // This is where the gRPC-Gateway proxies the requests
    conn, err := grpc.DialContext(
        context.Background(),
        "0.0.0.0:8080",
        grpc.WithBlock(),
        grpc.WithTransportCredentials(insecure.NewCredentials()),
    )
    if err != nil {
        log.Fatalln("Failed to dial server:", err)
    }

    gwmux := runtime.NewServeMux()
    // Register Greeter
    err = helloworldpb.RegisterGreeterHandler(context.Background(), gwmux, conn)
    if err != nil {
        log.Fatalln("Failed to register gateway:", err)
    }

    gwServer := &http.Server{
        Addr:    ":8090",
        Handler: gwmux,
    }

    log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
    log.Fatalln(gwServer.ListenAndServe())

}

