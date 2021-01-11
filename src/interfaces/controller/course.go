package controller

import (
	"course/src/application/dto"
	"course/src/application/service"
	"course/src/interfaces/util"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"net/http"
)

// CourseController 课程控制器
type CourseController struct {
	Service *service.CourseService `inject:"-"`
}

func NewCourseController() *CourseController {
	return &CourseController{}
}

func (c *CourseController) SimpleInfo(ctx *gin.Context) goft.Json {
	return c.Service.
		GetSimpleCourse(
			util.Exec(
				ctx.ShouldBindUri,
				&dto.SimpleCourseReq{}).Unwrap().(*dto.SimpleCourseReq))
}

func (c *CourseController) Build(goft *goft.Goft) {
	goft.Handle(http.MethodGet, "courses/:id", c.SimpleInfo)
}

func (c *CourseController) Name() string {
	return "CourseController"
}
