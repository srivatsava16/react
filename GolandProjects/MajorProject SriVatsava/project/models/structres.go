//In this file we will write all the structures which are used in our project

package models

type Task struct {
	Id                 int      ` json:"id"`
	Task               string   `json:"task"`
	Type               string   `json:"type"`
	CreatedAt          string   `json:"createdat"`
	AssignedEmployeeId int      `json:"assignedemployeeid"`
	Employees          Employee `gorm:"foreignKey:AssignedEmployeeId;references:Id"`
	Resolved           int      `json:"resolved"`
	ProjectId          int      `json:"projectid"`
	Projects           Project  `gorm:"foreignKey:ProjectId;references:Id"`
}
type Employee struct {
	Id                   int    `gorm:"primary_key" json:"id"`
	FirstName            string `gorm:"column:firstName" json:"firstname"`
	LastName             string `gorm:"column:lastName" json:"lastname"`
	FullName             string `gorm:"column:fullName" json:"fullname"`
	Jobtitle             string `gorm:"column:jobTitle" json:"jobtitle"`
	Jobtype              string `gorm:"column:jobType" json:"jobtype"`
	Phone                string `json:"phone"`
	Email                string `json:"email"`
	Image                string `json:"image"`
	Country              string `json:"country"`
	City                 string `json:"city"`
	Onboardingcompletion int    `gorm:"column:onboardingCompletion"json:"onboardingcompletiion"`
}
type Project struct {
	Id          int    `gorm:"primary_key" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tasks       string `json:"tasks"`
}
