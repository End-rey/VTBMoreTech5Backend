package handler

import (
	"net/http"

	"github.com/End-rey/VTBMoreTech5Backend/internal/dto"
	"github.com/gin-gonic/gin"
)

type allOfficesResponce struct {
	offices []dto.OfficeDTO
} 

// @Summary     all offices
// @Description Show all offices
// @ID          allOffices
// @Tags  	    office
// @Accept      json
// @Produce     json
// @Success     200 {object} allOfficesResponce
// @Failure     500 {object} response
// @Router      /api/offices [get]
func (h *Handler) getAllOffices(c *gin.Context) {
	offices, err := h.services.Office.GetAll()
	if err != nil {
		h.l.Error(err, "http - api - offices - getAllOffices")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}
	
	c.JSON(http.StatusOK, allOfficesResponce{offices})
}
