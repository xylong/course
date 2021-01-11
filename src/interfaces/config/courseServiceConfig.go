package config

import (
	"course/src/application/assembler"
	"course/src/application/service"
)

// CourseServiceConfig 课程服务配置
type CourseServiceConfig struct {
}

func NewCourseServiceConfig() *CourseServiceConfig {
	return &CourseServiceConfig{}
}

func (c *CourseServiceConfig) CourseService() *service.CourseService {
	return &service.CourseService{
		Req: &assembler.CourseRequest{},
		Rep: &assembler.CourseResponse{},
	}
}
