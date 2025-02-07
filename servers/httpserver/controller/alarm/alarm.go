package alarm

import (
	"net/http"
	"os"

	"project/library/resource"
	"project/library/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func List(c *gin.Context) {
	reqID := uuid.NewString()
	//todo

	resource.LoggerService.Info("Application started", zap.Int("pid", os.Getpid()))
	c.JSON(http.StatusOK, types.NewOKRestResp("test success", reqID))
}
