package controllers


import (

	"fmt"
	"context"

	model "backend_community_grpc/models"
	validation "backend_community_grpc/validations"
	repository "backend_community_grpc/repositories"
	proto "backend_community_grpc/proto/boilerplates"

)


type JoinRequest struct {}


func (*JoinRequest) UserJoin(

	ctx context.Context,
	join_request *proto.JoinRequest,

) (*proto.JoinResponse, error) {

	// Github User Validation
	valid, err := validation.ValidateGitHubUsername(join_request.GithubUsername)
	if err != nil {
		return &proto.JoinResponse {

			JoinResponse: "Server Error",
		}, nil
	}

	if !valid {
		return &proto.JoinResponse {

			JoinResponse: "Github Username invalid",
		}, nil
	}

	// request model
	request := model.JoinModel{

		UsernameGithub: join_request.GithubUsername,
	}
	
	
	message := ""

	// call repo to add join request to database
	if repo := repository.AddJoinRequest(request); repo {

		message = fmt.Sprintf(

			"Uraa... Permintaan '%v' untuk bergabung telah dikirim",
			request.UsernameGithub,
		)

	} else {
		
		message = fmt.Sprintf(
			"Hufft... Permintaan '%v' untuk bergabung tidak terkirim, coba lagi !.",
			request.UsernameGithub,

		)

	}


	return &proto.JoinResponse {

		JoinResponse: message,
	}, nil

}