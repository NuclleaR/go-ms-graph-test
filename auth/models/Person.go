package models

type Person struct {
	Birthday          string `json:"birthday"`
	CompanyName       string `json:"companyName"`
	Department        string `json:"department"`
	DisplayName       string `json:"displayName"`
	GivenName         string `json:"givenName"`
	Id                string `json:"id"`
	ImAddress         string `json:"imAddress"`
	IsFavorite        bool   `json:"isFavorite"`
	JobTitle          string `json:"jobTitle"`
	OfficeLocation    string `json:"officeLocation"`
	PersonNotes       string `json:"personNotes"`
	Profession        string `json:"profession"`
	Surname           string `json:"surname"`
	UserPrincipalName string `json:"userPrincipalName"`
	YomiCompany       string `json:"yomiCompany"`
	PersonType        struct {
		Class    string `json:"class"`
		Subclass string `json:"subclass"`
	} `json:"personType"`
	//PostalAddresses []struct {
	//	OdataType string `json:"@odata.type"`
	//} `json:"postalAddresses"`
	//Websites          []struct {
	//	OdataType string `json:"@odata.type"`
	//} `json:"websites"`
	//Phones []struct {
	//	OdataType string `json:"@odata.type"`
	//} `json:"phones"`
}
