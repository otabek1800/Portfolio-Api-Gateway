package handler

import pb "github.com/Javokhdev/Portfolio-Api-Gateway/genprotos"

type Handler struct {
	Skill            pb.SkillServiceClient
	Project			 pb.ProjectServiceClient
	Education	     pb.EducationServiceClient
	Experience       pb.ExperienceServiceClient
	Users            pb.UserServiceClient
}

func NewHandler(sl pb.SkillServiceClient, ps pb.ProjectServiceClient,
	edu pb.EducationServiceClient, ex pb.ExperienceServiceClient, us pb.UserServiceClient) *Handler {
	return &Handler{sl, ps, edu, ex, us}
}
