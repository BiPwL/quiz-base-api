package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type cleanTableRequest struct {
	Name string `uri:"name" binding:"required,min=1"`
}

func (server *Server) cleanTable(ctx *gin.Context) {
	var req cleanTableRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	isExist, err := server.store.TableExists(ctx, req.Name)
	if err != nil {
		fmt.Println(isExist, err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if !isExist {
		fmt.Println(isExist, err)
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	err = server.store.CleanTable(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}
