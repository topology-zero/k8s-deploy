package svc

import (
	"net/http/httptest"
	"time"

	"k8s-deploy/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// ServiceContext 实现 context.Context 接口
type ServiceContext struct {
	GinContext *gin.Context
	Log        *logrus.Entry
}

func NewServiceContext(c *gin.Context) *ServiceContext {
	traceId, _ := c.Get(util.TrafficKey)
	return &ServiceContext{
		GinContext: c,
		Log:        logrus.WithField(util.TraceId, traceId.(string)),
	}
}

func NewServiceContextWithoutGin() *ServiceContext {
	gin.SetMode(gin.ReleaseMode)
	ginCtx, _ := gin.CreateTestContext(httptest.NewRecorder())
	return &ServiceContext{
		GinContext: ginCtx,
		Log:        logrus.WithField(util.TraceId, uuid.New().String()),
	}
}

func (c *ServiceContext) Deadline() (deadline time.Time, ok bool) {
	return c.GinContext.Deadline()
}

func (c *ServiceContext) Done() <-chan struct{} {
	return c.GinContext.Done()
}

func (c *ServiceContext) Err() error {
	return c.GinContext.Err()
}

func (c *ServiceContext) Value(key any) any {
	return c.GinContext.Value(key)
}
