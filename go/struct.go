package main

type Service struct {
	Titre       string `json:"titre"`
	Description string `json:"description"`
	Categories  string `json:"categories"`
}

type Category struct {
	Nom      string    `json:"nom"`
	Services []Service `json:"services"`
}
