package main

import (
	"fmt"
	"log"

	"github.com/otabek1800/Portfolio-Api-Gateway/api"
	"github.com/otabek1800/Portfolio-Api-Gateway/api/handler"
	pb "github.com/otabek1800/Portfolio-Api-Gateway/genprotos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	UserConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":8085"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer UserConn.Close()

	UpaymentConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":8086"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer UpaymentConn.Close()
	skill := pb.NewSkillServiceClient(UserConn)
	pro := pb.NewProjectServiceClient(UserConn)
	edu := pb.NewEducationServiceClient(UserConn)
	exp := pb.NewExperienceServiceClient(UserConn)
	us := pb.NewUserServiceClient(UserConn)


	h := handler.NewHandler(skill, pro, edu, exp, us)
	r := api.NewGin(h)

	fmt.Println("Server started on port:8080")
	err = r.Run()
	if err != nil {
		log.Fatal("Error while Run: ", err.Error())
	}
}
