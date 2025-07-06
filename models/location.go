package models

type Results struct {
	Config
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
}
