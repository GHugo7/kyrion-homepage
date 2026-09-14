package main

type Service struct {
	Titre       string `json:"titre"`
	Description string `json:"description"`
	Categories  string `json:"categories"`
	Url         string `json:"url"`
	Online      bool   `json:"online"`
}

type Category struct {
	Nom      string    `json:"nom"`
	Services []Service `json:"services"`
}

type NPMProxyHost struct {
	DomainNames []string `json:"domain_names"`
	ForwardHost string   `json:"forward_host"`
	Enabled     bool     `json:"enabled"`
}

var tokenResp struct {
	Token string `json:"token"`
}

type Override struct {
	Category string
	URL      string
	Enabled  bool
}
