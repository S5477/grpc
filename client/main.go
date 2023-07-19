package main

import (
	"context"
	"encoding/json"
	"net/http"

	pb "github.com/S5477/grpc/client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// grpc server address
const address = "localhost:8081"

type Res struct {
	data Data
}

type Data struct {
	Msg string `json:"msg"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	err := http.ListenAndServe("localhost:3000", mux)

	if err != nil {
		panic(err)
	}

}

func hello(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	c := pb.NewServerClient(conn)

	a, err := c.Hello(
		context.Background(),
		&pb.HelloRequest{
			Name: "S5477",
		},
	)

	data := setRes(a.Msg)

	res, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func setRes(msg string) Res {
	return Res{
		Data{
			Msg: msg,
		},
	}
}
