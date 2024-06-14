package handler

import (
	pb "github.com/Javokhdev/Portfolio-Api-Gateway/genprotos"

	"github.com/gin-gonic/gin"
)

// Create 			Experience handles the creation of a new Porfolio
// @Summary 		Create Porfolio
// @Description 	Create page
// @Tags 			Experience
// @Accept  		json
// @Produce  		json 
// @Param   		Create  body    pb.Experience  true   "Create"
// @Success 		200   {string}  string  	"Create Successful"
// @Failure 		401   {string}  string  	"Error while Created"
// @Router 			/experience/create [post]
func (h *Handler) CreateExperience(ctx *gin.Context) {
	arr := pb.Experience{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.Experience.CreateExperience(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// UpdateExperience 	handles the creation of a new Experience
// @Summary 		Update Experience
// @Description 	Update page
// @Tags 			Experience
// @Accept  		json
// @Produce  		json
// @Param     		id path string true "Experience ID"
// @Param   		Update  body    pb.Experience  true   "Update"
// @Success 		200   {string}  string      "Update Successful"
// @Failure 		401   {string}  string      "Error while created"
// @Router 			/experience/update/{id} [put]
func (h *Handler) UpdateExperience(ctx *gin.Context) {
	arr := pb.Experience{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.Experience.UpdateExperience(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// DeleteExperience 	handles the creation of a new Experience
// @Summary 		Delete Experience
// @Description 	Delete page
// @Tags 			Experience
// @Accept  		json
// @Produce  		json
// @Param     		id     path    string   true  "Experience ID"
// @Success			200  {string}  string  "Delete Successful"
// @Failure 		401  {string}  string  "Error while Deleted"
// @Router 			/experience/delete/{id} [delete]
func (h *Handler) DeleteExperience(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.Experience.DeleteExperience(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// GetAllExperience 	handles the creation of a new Experience
// @Summary 		GetAll Experience
// @Description 	GetAll page
// @Tags 			Experience
// @Accept  		json
// @Produce  		json
// @Param 			query  query   pb.Experience true  "Query parameter"
// @Success 		200  {object}  pb.GetAllExperiences  	"GetAll Successful"
// @Failure 		401  {string}  string  				"Error while GetAlld"
// @Router 			/experience/getall [get]
func (h *Handler) GetAllExperience(ctx *gin.Context) {
	Experience := &pb.Experience{}
	// restoran.Address = ctx.Param("restoran")
	// Experience.Name = ctx.Param("name")

	res, err := h.Experience.GetAllExperience(ctx, Experience)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByIdExperience 	handles the creation of a new Experience
// @Summary 		GetById Experience
// @Description 	GetById page
// @Tags 			Experience
// @Accept  		json
// @Produce  		json
// @Param     		id    path    string  true  "Experience ID"
// @Success 		200 {object}  pb.Experience   "GetById Successful"
// @Failure 		401 {string}  string 		"Error while GetByIdd"
// @Router 			/experience/getbyid/{id} [get]
func (h *Handler) GetByIdExperience(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.Experience.GetByIdExperience(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}
