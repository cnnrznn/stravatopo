package main

import (
    "context"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/cnnrznn/stravatopo/message"
)

func handler(ctx context.Context, req message.Request) (message.Response, error) {
    return message.Response{ []int{0, 1, 2}, []int{3, 4, 5} }, nil
}

func main() {
    lambda.Start(handler)
}

