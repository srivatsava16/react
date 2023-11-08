package models

type Employee struct {
	Id        string
	FirstName string
	LastName  string
	FullName  string
	Phone     string
	Email     string
	City      string
	//Img       []byte
}

type Response struct {
	Status string
	Code   string
	Error  string
	Data   interface{}
}
type Manualresponse struct {
	Hit    string
	Input  Employee
	Manual Response
}
type ExtJson struct {
	HitName          string
	HitMatch         bool
	Input            Employee
	ManualResponse   Response
	ExternalResponse Response
}
