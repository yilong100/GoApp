package models

type User struct {
	ID               string `json: id`
	Name             string `json: "name"`
	Age              int    `json: "age"`
	DreamPlaceToLive string `json: "dreamplacetolive`
}

var Users = []User{
	{ID: "1", Name: "James Wu", Age: 22, DreamPlaceToLive: "Toronto"},
	{ID: "2", Name: "Damian Lillard", Age: 33, DreamPlaceToLive: "Miami"},
	{ID: "3", Name: "Emma Mackey", Age: 27, DreamPlaceToLive: "Paris"},
}

type ResponseObject struct {
	Message string `json: message`
}
