package api

import (
	"github.com/otsbek1800/Portfolio-Api-Gateway/api/handler"
	_"github.com/otabek1800/Portfolio-Api-Gateway/docs"


	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tite Portfolio
// @version 1.0
// @description Portfolio
// @host localhost:8080
// @BasePath /
func NewGin(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
	u := r.Group("/skill")
	u.POST("/create", h.CreateSkill)
	u.PUT("/update/:id", h.UpdateSkill)
	u.DELETE("/delete/:id", h.DeleteSkill)
	u.GET("/getall", h.GetAllSkill)
	u.GET("/getbyid/:id", h.GetByIdSkill)



	p := r.Group("/project")
	p.POST("/create", h.CreateProject)
	p.PUT("/update/:id", h.UpdateProject)
	p.DELETE("/delete/:id", h.DeleteProject)
	p.GET("/getall", h.GetAllProject)
	p.GET("/getbyid/:id", h.GetByIdProject)



	e := r.Group("/education")
	e.POST("/create", h.CreateEducation)
	e.PUT("/update/:id", h.UpdateEducation)
	e.DELETE("/delete/:id", h.DeleteEducation)
	e.GET("/getall", h.GetAllEducation)
	e.GET("/getbyid/:id", h.GetByIdEducation)



	v := r.Group("/experience")
	v.POST("/create", h.CreateExperience)
	v.PUT("/update/:id", h.UpdateExperience)
	v.DELETE("/delete/:id", h.DeleteExperience)
	v.GET("/getall", h.GetAllExperience)
	v.GET("/getbyid/:id", h.GetByIdExperience)


	return r
}
