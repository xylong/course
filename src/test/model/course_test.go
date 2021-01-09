package model

import (
	course2 "course/src/domain/model/course"
	"course/src/test/lib"
	"testing"
)

func addCourse(name string,time int64) error  {
	course:=course2.New(
		course2.SetName(name),
		course2.SetTime(time),
		)
	return lib.DB.Create(course).Error
}

func TestAddCourse(t *testing.T) {
	list:=[]struct{
		name string
		time int64
	}{
		{"红楼梦", 56515},
		{"西游记", 123456},
		{"三国演义", 1234},
		{"水浒传", 5456},
		{"哈利·波特", 46454654},
		{"The Lord of the Rings", 0},
	}

	for i:=range list{
		name,time:=list[i].name,list[i].time
		if err:=addCourse(name,time);err!=nil {
			t.Errorf("create %s error: %s\n", name, err.Error())
		}
	}
}