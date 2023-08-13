package main


import (
	
	exception "backend_community_grpc/exceptions"
	config "backend_community_grpc/configurations"
	controller "backend_community_grpc/controllers"
	proto "backend_community_grpc/proto/boilerplates"

)



func main() {

	
	// service network listen
	network_listen := config.NetworkListen()

	// grpc handler
	grpc_handler := config.GrpcHandler()


	// register gRPC UserService to gRPC server handler
	proto.RegisterUserJoinServiceServer(

		grpc_handler,
		&controller.JoinRequest{},
	)


	// run service
	err := grpc_handler.Serve(network_listen)
	exception.TryCatchError(err)
	

}