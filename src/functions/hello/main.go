package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type HelloRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type HelloResponse struct {
	Message string      `json:"message"`
	Event   interface{} `json:"event"`
}

func HandleHelloRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received event body: %s", event.Body)

	var req HelloRequest
	err := json.Unmarshal([]byte(event.Body), &req)
	if err != nil {
		log.Printf("Error unmarshalling request: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, err
	}

	response := HelloResponse{
		Message: fmt.Sprintf("Hello %s, you are %d years old!", req.Name, req.Age),
		Event:   event,
	}

	responseBody, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshalling response: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       `{"error": "internal server error"}`,
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		Body:            string(responseBody),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(HandleHelloRequest)
}
