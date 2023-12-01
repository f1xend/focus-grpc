package domain

type Rabbit struct {
	Color string
}

type Animal struct {
	Name  string `json:"animal"`
	Color string `json:"color"`
}
