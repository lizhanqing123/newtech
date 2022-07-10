
package main

import (
    "context"
    "fmt"
    gtrace "github.com/moxiaomomo/grpc-jaeger"
    "google.golang.org/grpc"
    pb "grpc2/proto"
    "os"
    "time"
)


const(
 address = ":8080"
 defaultName = "word123"
)

func main () {

      dialOpts := []grpc.DialOption{grpc.WithInsecure(),grpc.WithBlock()}

      tracer, _, err := gtrace.NewJaegerTracer("testCli", "192.168.129.130:6831")
      if err != nil {
         fmt.Printf("new tracer err: %+v\n", err)
         os.Exit(-1)
      }

      if tracer != nil {
         dialOpts = append(dialOpts, gtrace.DialOption(tracer))
      }

     conn, err2 := grpc.Dial(address, dialOpts...)
     if err2 != nil {
      fmt.Printf("did not connect: %v", err2)
     }
     defer conn.Close()
     c := pb.NewGreeterClient(conn)

     name := defaultName
     if len(os.Args) > 1 {
      name = os.Args[1]
     }
     ctx, cancel := context.WithTimeout(context.Background(), time.Second)
     defer cancel()

     r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
     if err != nil {
     fmt.Printf("could not greet: %v", err)
     }
     fmt.Printf("Greeting: %s", r.GetMessage())

     select{}
}

