package httpresponse

import (
	"github.com/gin-gonic/gin"
	"github.com/spacetronot-research-team/catalog-service/pkg/trace"
)

func Write(ctx *gin.Context, code int, data interface{}, err error) {
	resBody := gin.H{
		"data":       data,
		"request_id": ctx.GetString(trace.Key),
	}

	if err != nil {
		resBody["error"] = err.Error()
	} else {
		resBody["error"] = nil
	}

	ctx.JSON(code, resBody)
}
