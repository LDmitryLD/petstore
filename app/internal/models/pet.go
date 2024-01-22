package models

type Pet struct {
	ID        int        `json:"id,omitempty"`
	Category  Category   `json:"category,omitempty"`
	Name      string     `json:"name,omitempty"`
	PhotoUrls []string   `json:"photoUrls,omitempty"`
	Tags      []Category `json:"tags,omitempty"`
	Status    string     `json:"status,omitempty"`
}

type Category struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
