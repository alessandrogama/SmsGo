package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

// Declare variables
var (
	accountSid string
	authToken  string
	fromPhone  string
	toPhone    string
	client     *twilio.RestClient
)

// Function for send menssages
func SendMessage(msg string) {

	params := &openapi.CreateMessageParams{}
	params.SetTo(toPhone)
	params.SetFrom(fromPhone)
	params.SetBody(msg)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}

// Loading variables for .env and client Twilio
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("error loading .env:%s\n", err.Error())
		os.Exit(1)
	}
	accountSid = os.Getenv("ACCOUNT_SID")
	authToken = os.Getenv("AUTH_TOKEN")
	fromPhone = os.Getenv("FROM_PHONE")
	toPhone = os.Getenv("TO_PHONE")

	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

}

// Call function send message
func main() {

	msg := fmt.Sprintf(os.Getenv("MSG"), "GAMA")
	SendMessage(msg)
}
