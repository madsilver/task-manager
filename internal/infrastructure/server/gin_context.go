package server

import (
	"github.com/madsilver/task-manager/internal/adapter/handler"
	"time"

	"github.com/gin-gonic/gin"
)

type ginContext struct {
	ctx *gin.Context
}

func NewContextHTTP(c *gin.Context) handler.ContextHTTP {
	return &ginContext{ctx: c}
}

func (s *ginContext) Param(name string) string {
	return s.ctx.Param(name)
}

func (s *ginContext) QueryParam(key string) string {
	return s.ctx.Query(key)
}

func (s *ginContext) Path() string {
	return s.ctx.Request.URL.Path
}

func (s *ginContext) AddParam(key, value string) {
	s.ctx.Params = append(s.ctx.Params, gin.Param{Key: key, Value: value})
}

func (s *ginContext) Query(key string) string {
	return s.ctx.Query(key)
}

func (s *ginContext) JSON(code int, i any) error {
	s.ctx.JSON(code, i)
	return nil
}

func (s *ginContext) Bind(i interface{}) error {
	return s.ctx.ShouldBindJSON(i)
}

func (s *ginContext) Data(code int, contentType string, data []byte) error {
	s.ctx.Data(code, contentType, data)
	return nil
}

func (s *ginContext) Abort(code int) {
	s.ctx.AbortWithStatus(code)
}

func (s *ginContext) Deadline() (time.Time, bool) {
	return s.ctx.Request.Context().Deadline()
}

func (s *ginContext) Done() <-chan struct{} {
	return s.ctx.Request.Context().Done()
}

func (s *ginContext) Err() error {
	return s.ctx.Request.Context().Err()
}

func (s *ginContext) Value(key any) any {
	return s.ctx.Value(key)
}
