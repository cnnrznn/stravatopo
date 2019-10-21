package main

import (
    "context"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/cnnrznn/stravatropo"
)

type Request struct {
    PingURL string `json:"pingURL"`
    GetURL string `json:"getURL"`
}

type Response struct {
    PingLats []int `json:"pingLats"`
    GetLats []int `json:"getLats"`
}

func handler(ctx context.Context, req Request) (Response, error) {
    return Response{ []int{0, 1, 2}, []int{3, 4, 5} }, nil
}

func main() {
    lambda.Start(handler)
}

