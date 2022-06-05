package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seriesmanager-services/dto"
	"seriesmanager-services/helpers"
	"seriesmanager-services/middlewares"
	"seriesmanager-services/models"
	"seriesmanager-services/services"
)

type SeasonController interface {
	Routes(e *gin.Engine)
	PostSeason(ctx *gin.Context)
	GetDistinctBySid(ctx *gin.Context)
	GetInfosBySeasonBySeriesId(ctx *gin.Context)
}

type seasonController struct {
	seasonService services.SeasonService
	jwtHelper     helpers.JwtHelper
}

func NewSeasonController(seasonService services.SeasonService, jwtHelper helpers.JwtHelper) SeasonController {
	return &seasonController{seasonService: seasonService, jwtHelper: jwtHelper}
}

func (s *seasonController) Routes(e *gin.Engine) {
	routes := e.Group("/api/seasons", middlewares.AuthorizeJwt(s.jwtHelper))
	{
		routes.POST("/", s.PostSeason)
		routes.GET("/series/:id", s.GetDistinctBySid)
		routes.GET("/:number/series/:id/infos", s.GetInfosBySeasonBySeriesId)
	}
}

// PostSeason user adds a season
func (s *seasonController) PostSeason(ctx *gin.Context) {
	var seasonDto dto.SeasonCreateDto
	if errDto := ctx.ShouldBind(&seasonDto); errDto != nil {
		response := helpers.NewResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	res := s.seasonService.AddSeason(seasonDto)

	if season, ok := res.(models.Season); ok {
		response := helpers.NewResponse("Saison ajoutée", season)
		ctx.JSON(http.StatusCreated, response)
	} else {
		response := helpers.NewResponse("Informations invalides", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

// GetDistinctBySid gets series seasons by series sid
func (s *seasonController) GetDistinctBySid(ctx *gin.Context) {
	seasons := s.seasonService.GetDistinctBySeriesId(ctx.Param("id"))
	response := helpers.NewResponse("", seasons)
	ctx.JSON(http.StatusOK, response)
}

// GetInfosBySeasonBySeriesId get season user infos
func (s *seasonController) GetInfosBySeasonBySeriesId(ctx *gin.Context) {
	infos := s.seasonService.GetInfosBySeasonBySeriesId(ctx.Param("id"), ctx.Param("number"))
	response := helpers.NewResponse("", infos)
	ctx.JSON(http.StatusOK, response)
}
