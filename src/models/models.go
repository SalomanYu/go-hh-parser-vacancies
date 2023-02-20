package models


type City struct {
	HH_ID 		int
	EDWICA_ID 	int
	Name		string
}

type Vacancy struct {
	Id              int
	Title           string
	ProfessionId    int
	CityId			int
	DateUpdate		string
	Url             string
	Experience		string
	SalaryFrom      float64
	SalaryTo        float64
	ProfAreas		string // Slice joined by |
	Skills          string
	Specializations string
}

type Currency struct {
	Code	string
	Abbr	string
	Name	string
	Rate	float64
}

type Salary struct {
	From	float64
	To		float64
	Currency string
}

type Profession struct {
	Id			int
	Name		string
	OtherNames 	[]string
	Level 		int
	ParentId 	int
}