package main

import (
    "encoding/json"
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/lambda"
    "github.com/cnnrznn/stravatopo/message"
    "os"
)

func main() {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    creds, err := sess.Config.Credentials.Get()
    if err != nil {
        panic("Cannot get credentials")
    }
    fmt.Println("AWS Credentials:", creds)

    client := lambda.New(sess, &aws.Config{Region: aws.String("us-west-2")})

    //result, err := client.ListFunctions(nil)
    //if err != nil {
    //    fmt.Println("Cannot list functions", err)
    //    os.Exit(0)
    //}

    //fmt.Println("Functions:")

    //for _, f := range result.Functions {
    //    fmt.Println("Name:        " + aws.StringValue(f.FunctionName))
    //    fmt.Println("Description: " + aws.StringValue(f.Description))
    //    fmt.Println("")
    //}

    request := message.Request{PingURL: "strava.com", GetURL: "strava.com"}
    var response message.Response

    payload, err := json.Marshal(request)
    if err != nil {
        fmt.Println("Error marshalling request")
        os.Exit(0)
    }

    result, err := client.Invoke(&lambda.InvokeInput{FunctionName: aws.String("message"),
                                                     Payload: payload})
    if err != nil {
        fmt.Println("Error calling message")
        os.Exit(0)
    }

    err = json.Unmarshal(result.Payload, &response)
    if err != nil {
        fmt.Println("Error unmarshalling result")
        os.Exit(0)
    }

    if response.StatusCode != 200 {
        fmt.Println("Error response code:", response.StatusCode)
        os.Exit(0)
    }

    // If the result is failure, we got an error
    if response.Body.Result == "failure" {
        fmt.Println("Failed to get items")
        os.Exit(0)
    }

    // Print out items
    if len(resp.Body.Data) > 0 {
        for i := range resp.Body.Data {
            fmt.Println(resp.Body.Data[i].Item)
        }
    } else {
        fmt.Println("There were no items")
    }
}

