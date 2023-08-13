package configurations

import (

	"os"
	"net"
	"fmt"
	"strconv"
	"github.com/joho/godotenv"

	"google.golang.org/grpc"

	logger "backend_community_grpc/middlewares"
	exception "backend_community_grpc/exceptions"

)





/*
 *  Environment Variable Configuration
 */

type envStruct struct {

	// Network Environment
	Endpoint string
	NetworkProtocol string

	// Database Environment
	KeyJsonBase64  string
	SpreadsheetId  int
	SpreadsheetKey string

}


func EnvironmentVariables() envStruct {

	// load environment variables
	err := godotenv.Load()
	exception.TryCatchError(err)


	// convert string to integer "SPREADSHEET_KEY" environment variable
	spreadsheet_id_conv, err := strconv.Atoi(
		os.Getenv("SPREADSHEET_ID"),
	)

	exception.TryCatchError(err)


	// environment variable into struct
	var environment = envStruct{

		// network environment
		Endpoint  : fmt.Sprintf(

			"%v:%v",
			os.Getenv("HOST"),
			os.Getenv("PORT"),
		),

		NetworkProtocol : os.Getenv("PROTOCOL"),

		// database environment
		KeyJsonBase64  : os.Getenv("KEY_JSON_BASE64"),
		SpreadsheetId  : spreadsheet_id_conv,
		SpreadsheetKey : os.Getenv("SPREADSHEET_KEY"),

	}


	return environment

}





/*
 *  Network Configuration
 */

 func NetworkListen() (net.Listener) {

	listen, err := net.Listen(

		// network protocol
		EnvironmentVariables().NetworkProtocol,

		// endpoint
		EnvironmentVariables().Endpoint,

	)

	exception.TryCatchError(err)

	
	return listen

}





/*
 *  gRPC Configuration
 */

func GrpcHandler() (*grpc.Server) {

	grpc_service := grpc.NewServer(
		grpc.UnaryInterceptor(

			logger.GrpcLogger,
		),
	)


	return grpc_service

}