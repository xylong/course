package ctrl

import (
	course2 "course/src/domain/model/course"
	"course/src/test/lib"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"net/http"
)

func init() {
	RegisterController(&Course{})
}

type Course struct {

}

func (c *Course) Name() string {
	return "Course"
}

func (c *Course) Store(ctx *gin.Context) goft.Json {
	course:=course2.New()
	goft.Error(ctx.ShouldBindJSON(course))
	goft.Error(lib.DB.Create(course).Error)
	return gin.H{"msg":course.ID}
}

func (c *Course) Build(goft *goft.Goft) {
	goft.Handle(http.MethodPost,"courses",c.Store)
}