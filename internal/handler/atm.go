package handler

import (
	"net/http"

	"github.com/End-rey/VTBMoreTech5Backend/internal/dto"
	"github.com/gin-gonic/gin"
)

type allAtmsResponce struct {
	offices []dto.OfficeDTO
} 

// @Summary     all atms
// @Description Show all atms
// @ID          allAtms
// @Tags  	    atm
// @Accept      json
// @Produce     json
// @Success     200 {object} allAtmsResponce
// @Failure     500 {object} response
// @Router      /api/atms [get]
func (h *Handler) getAllAtms(c *gin.Context) {
	atms, err := h.services.Office.GetAll()
	if err != nil {
		h.l.Error(err, "http - api - atms - getAllAtms")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}
	
	c.JSON(http.StatusOK, allAtmsResponce{atms})
}