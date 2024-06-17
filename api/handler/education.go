package handler

import (
	pb "github.com/otabek1800/Portfolio-Api-Gateway/genprotos"

	"github.com/gin-gonic/gin"
)

// Create 			Education handles the creation of a new Porfolio
// @Summary 		Create Porfolio
// @Description 	Create page
// @Tags 			Education
// @Accept  		json
// @Produce  		json
// @Param   		Create  body    pb.Education  true   "Create"
// @Success 		200   {string}  string  	"Create Successful"
// @Failure 		401   {string}  string  	"Error while Created"
// @Router 			/education/create [post]
func (h *Handler) CreateEducation(ctx *gin.Context) {
	arr := pb.Education{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.Education.CreateEducation(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success>>>>>>>>")
}

// UpdateEducation 	handles the creation of a new Education
// @Summary 		Update Education
// @Description 	Update page
// @Tags 			Education
// @Accept  		json
// @Produce  		json
// @Param     		id path string true "Education ID"
// @Param   		Update  body    pb.Education  true   "Update"
// @Success 		200   {string}  string      "Update Successful"
// @Failure 		401   {string}  string      "Error while created"
// @Router 			/education/update/{id} [put]
func (h *Handler) UpdateEducation(ctx *gin.Context) {
	arr := pb.Education{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	_, err = h.Education.UpdateEducation(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success>>>>>>>")
}

// DeleteEducation 	handles the creation of a new Education
// @Summary 		Delete Education
// @Description 	Delete page
// @Tags 			Education
// @Accept  		json
// @Produce  		json
// @Param     		id     path    string   true  "Education ID"
// @Success			200  {string}  string  "Delete Successful"
// @Failure 		401  {string}  string  "Error while Deleted"
// @Router 			/education/delete/{id} [delete]
func (h *Handler) DeleteEducation(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	_, err := h.Education.DeleteEducation(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, "Success>>>>>>>>")
}

// GetAllEducation 	handles the creation of a new Education
// @Summary 		GetAll Education
// @Description 	GetAll page
// @Tags 			Education
// @Accept  		json
// @Produce  		json
// @Param 			query  query   pb.Education true  "Query parameter"
// @Success 		200  {object}  pb.GetAllEducations  	"GetAll Successful"
// @Failure 		401  {string}  string  				"Error while GetAlld"
// @Router 			/education/getall [get]
func (h *Handler) GetAllEducation(ctx *gin.Context) {
	Education := &pb.Education{}
	// restoran.Address = ctx.Param("restoran")
	// Education.Name = ctx.Param("name")

	res, err := h.Education.GetAllEducation(ctx, Education)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByIdEducation 	handles the creation of a new Education
// @Summary 		GetById Education
// @Description 	GetById page
// @Tags 			Education
// @Accept  		json
// @Produce  		json
// @Param     		id    path    string  true  "Education ID"
// @Success 		200 {object}  pb.Education   "GetById Successful"
// @Failure 		401 {string}  string 		"Error while GetByIdd"
// @Router 			/education/getbyid/{id} [get]
func (h *Handler) GetByIdEducation(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}
	res, err := h.Education.GetByIdEducation(ctx, &id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}
