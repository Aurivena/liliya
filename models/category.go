package models

type CategoryNames struct {
	Names []string `json:"categories"`
}
type Category struct {
	Id   string `json:"category_id"`
	Name string `json:"name"`
}
