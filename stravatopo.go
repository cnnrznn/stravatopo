package main

import (
    "fmt"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/lambda"
    //"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    client := lambda.New(sess, &aws.Config{Region: aws.String("us-west-2")})

    result, err := client.ListFunctions(nil)
    if err != nil {
        fmt.Println("Cannot list functions", err)
        os.Exit(0)
    }

    fmt.Println("Functions:")

    for _, f := range result.Functions {
        fmt.Println("Name:        " + aws.StringValue(f.FunctionName))
        fmt.Println("Description: " + aws.StringValue(f.Description))
        fmt.Println("")
    }
}

