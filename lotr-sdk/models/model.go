package models

type Character struct {
	Id      string `json:"_id"`
	Birth   string `json:"birth"`
	Height  string `json:"height"`
	Race    string `json:"race"`
	Gender  string `json:"gender"`
	Spouse  string `json:"spouse"`
	Death   string `json:"death"`
	Realm   string `json:"realm"`
	Hair    string `json:"hair"`
	Name    string `json:"name"`
	WikiUrl string `json:"wikiUrl"`
}

type Book struct {
	Id   string `json:"_id"`
	Name string `json:"name"`
}

type Chapter struct {
	Id          string `json:"_id"`
	ChapterName string `json:"chapterName"`
}

type Movie struct {
	Id                         string `json:"_id"`
	Name                       string `json:"name"`
	RunTimeInMinutes           int    `json:"runtimeInMinutes"`
	BudgetInMillions           int    `json:"budgetInMillions"`
	BoxOfficeRevenueInMillions int    `json:"boxOfficeRevenueInMillions"`
	AcademyAwardNominations    int    `json:"academyAwardNominations"`
	AcademyAwardWins           int    `json:"academyAwardWins"`
	RottenTomatoesScore        int    `json:"rottenTomatoesScore"`
}

type Quote struct {
	Id        string `json:"_id"`
	Dialog    string `json:"dialog"`
	Movie     string `json:"movie"`
	Character string `json:"character"`
	AId       string `json:"id"`
}
