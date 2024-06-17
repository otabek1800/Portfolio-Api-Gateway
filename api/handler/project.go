package handler

import (
	pb "github.com/Javokhdev/Portfolio-Api-Gateway/genprotos"

	"github.com/gin-gonic/gin"
)

// Create 			Project handles the creation of a new Project
// @Summary 		Create Project
// @Description 	Create page
// @Tags 			Project
// @Accept  		json
// @Produce  		json
// @Param   		Create  body    pb.Project  true   "Create"
// @Success 		200   {string}  string  	"Create Successful"
// @Failure 		401   {string}  string  	"Error while Created"
// @Router 			/project/create [post]
func (h *Handler) CreateProject(ctx *gin.Context) {
	arr := pb.Project{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.Project.CreateProject(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// UpdateProject 	handles the creation of a new Project
// @Summary 		Update Project
// @Description 	Update page
// @Tags 			Project
// @Accept  		json
// @Produce  		json
// @Param     		id path string true "Project ID"
// @Param   		Update  body    pb.Project  true   "Update"
// @Success 		200   {string}  string      "Update Successful"
// @Failure 		401   {string}  string      "Error while created"
// @Router 			/project/update/{id} [put]
func (h *Handler) UpdateProject(ctx *gin.Context) {
	arr := pb.Project{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.Project.UpdateProject(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// DeleteProject 	handles the creation of a new Project
// @Summary 		Delete Project
// @Description 	Delete page
// @Tags 			Project
// @Accept  		json
// @Produce  		json
// @Param     		id     path    string   true  "Project ID"
// @Success			200  {string}  string  "Delete Successful"
// @Failure 		401  {string}  string  "Error while Deleted"
// @Router 			/project/delete/{id} [delete]
func (h *Handler) DeleteProject(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.Project.DeleteProject(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// GetAllProject 	handles the creation of a new Project
// @Summary 		GetAll Project
// @Description 	GetAll page
// @Tags 			Project
// @Accept  		json
// @Produce  		json
// @Param 			query  query   pb.Project true  "Query parameter"
// @Success 		200  {object}  pb.GetAllProjects  	"GetAll Successful"
// @Failure 		401  {string}  string  				"Error while GetAlld"
// @Router 			/project/getall [get]
func (h *Handler) GetAllProject(ctx *gin.Context) {
	Project := &pb.Project{}
	// restoran.Address = ctx.Param("restoran")
	// Project.Name = ctx.Param("name")

	res, err := h.Project.GetAllProject(ctx, Project)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByIdProject 	handles the creation of a new Project
// @Summary 		GetById Project
// @Description 	GetById page
// @Tags 			Project
// @Accept  		json
// @Produce  		json
// @Param     		id    path    string  true  "Project ID"
// @Success 		200 {object}  pb.Project   "GetById Successful"
// @Failure 		401 {string}  string 		"Error while GetByIdd"
// @Router 			/Project/getbyid/{id} [get]
func (h *Handler) GetByIdProject(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.Project.GetByIdProject(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}
