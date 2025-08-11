package models

type Results struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
}
