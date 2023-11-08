package models

type Student struct {
	Id        string
	Firstname string
	Lastname  string
	Username  string
	Password  string
	Address   string
	Place     string
}
type Response struct {
	Status string
	Code   int
	Error  string
	Data   interface{}
}
type Manualresponse struct {
	Hit    string
	Input  Student
	Manual Response
}
type ExtJson struct {
	HitName          string
	HitMatch         bool
	Input            Student
	ManualResponse   Response
	ExternalResponse Response
}
