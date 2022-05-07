package cmd

import (
	"context"
	"log"
	"net/http"

	openapiclient "github.com/LightWebService/CommunicationClient"
	"github.com/spf13/cobra"
)

var (
	registerEmail    string
	registerPassword string
	registerNickName string
)

var registerCommand = &cobra.Command{
	Use:   "register",
	Short: "Register to LWS Service",
	Long:  "Required Parameters: --email, --nickname, --password",
	Run: func(cmd *cobra.Command, args []string) {
		// Create Request
		registerRequest := *openapiclient.NewRegisterRequest()
		registerRequest.UserEmail = &registerEmail
		registerRequest.UserNickName = &registerNickName
		registerRequest.UserPassword = &registerPassword

		// Setup API Location
		configuration := *openapiclient.NewConfiguration()
		apiClient := *openapiclient.NewAPIClient(&configuration)

		response, err := apiClient.AccountApi.ApiAccountPost(context.Background()).RegisterRequest(registerRequest).Execute()

		if err != nil {
			handleRegistrationError(&registerRequest, response, err)
		} else {
			log.Println("Successfully Registered to Service! Please login.")
		}
	},
}

func handleRegistrationError(requestBody *openapiclient.RegisterRequest, response *http.Response, err error) {
	defer response.Body.Close()

	switch response.StatusCode {
	case 400:
		log.Fatalf("Please verify email address is in valid format.")
	case 409:
		log.Fatalf("User email with %s already exists! Please use another email.", *requestBody.UserEmail)
	default:
		log.Fatalf("Unknown Error occurred! : %s\n", response.Status)
	}
}

func init() {
	registerCommand.Flags().StringVarP(&registerEmail, "email", "E", "", "Registration Email.")
	registerCommand.Flags().StringVarP(&registerPassword, "password", "P", "", "Registration Password")
	registerCommand.Flags().StringVarP(&registerNickName, "nickname", "N", "", "Registration Nick Name")

	registerCommand.MarkFlagRequired("email")
	registerCommand.MarkFlagRequired("password")
	registerCommand.MarkFlagRequired("nickname")
}
