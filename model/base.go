package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Content      string
	PreferGender int
	Creator      User
	GradeWanted  Grade
	CourseWanted Course
}

type User struct {
	gorm.Model
	Name     string
	Password string
	Phone    string
}

type Grade struct {
	gorm.Model
	Num         int
	DisplayText string
}

type Course struct {
	gorm.Model
	Name        string
	Description string
}
