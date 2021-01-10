package model

import (
	course2 "course/src/domain/model/course"
	"course/src/infrastructure/GormDao"
	"course/src/test/internal/lib"
	"testing"
)

func addCourse(name string, time int64) error {
	course := course2.New(
		course2.SetName(name),
		course2.SetTime(time),
	)
	return lib.DB.Create(course).Error
}

// TestAddCourse 测试添加课程
func TestAddCourse(t *testing.T) {
	list := []struct {
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

	for i := range list {
		name, time := list[i].name, list[i].time
		if err := addCourse(name, time); err != nil {
			t.Errorf("create %s error: %s\n", name, err.Error())
		}
	}
}

// TestCourseLoad 测试课程查询
func TestCourseLoad(t *testing.T) {
	repo := GormDao.NewCourseRepo(lib.DB)
	course := course2.New(course2.SetRepo(repo))

	arr := []int{0, 1}
	for _, id := range arr {
		course.ID = id
		if err := course.Load(); err != nil {
			if course.ID != 0 {
				t.Error(err)
			} else {
				t.Log(course.ID)
			}
		} else {
			t.Logf("%s", course.Info.Name)
		}
	}
}
