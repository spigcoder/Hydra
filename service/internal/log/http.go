package log

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/spigcoder/Hydra/pkg/errs"
)

type HTTPHandler struct {
	log *Log
}

func NewHTTPHandler(log *Log) *HTTPHandler {
	return &HTTPHandler{
		log: log,
	}
}

func (h *HTTPHandler) Produce(ctx *gin.Context) {
	var req ProduceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errs.Wrap(errs.BadRequest, "ctx.ShouldBindJson").WithErr(err).Response(ctx)
		return
	}

	offset, _ := h.log.Append(req.Record)

	ctx.JSON(http.StatusOK, ProduceResponse{
		Offset: offset,
	})
}

func (h *HTTPHandler) Consume(ctx *gin.Context) {
	var req ConsumeRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		errs.Wrap(errs.BadRequest, "ctx.ShouldBindJson").WithErr(err).Response(ctx)
		return
	}

	record, err := h.log.Read(req.Offset)
	if err != nil {
		errs.Wrap(errs.NotFound, "req.offset > database max offset").WithErr(err).Response(ctx)
		return
	}

	ctx.JSON(http.StatusOK, ConsumeResponse{
		Record: record,
	})
}
