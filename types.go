package main

import "fmt"

type course struct {	
	id	int

	credits credits

	title string
	id	int
	classNum int
	teacher	string
	longDesc	string
	cap	int
}

type student struct
	major major
	minor	minor
	credits credits
}

type major struct {
	title string
}

type minor struct {
	title string
}

// type credits int 

type report {
	courses []course
	id int
}

type studentReport {
	student_id int
}

type Reporter interface {
	Generate() (StudentReport, error)
}

type CourseGetter interface {
	GetCourse(id int) (*Course, error)
	GetCourses() (*[]Course, error)
	GetCoursePrerequisites(Course) ([]Course, error)
	FilterCourses([]Course) []Course
}

type StudentGetter interface {
	GetStudent(student_id int)	Student
	GetStudentReport(student_id int)  (*studentReport, error)
	GetStudentMajor(student_id int) *Major
	GetStudentMinor(student_id int) *Minor
}

type StudentController interface {
	Create() Student
	Delete() *Student
}
