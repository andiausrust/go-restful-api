package model

type Book struct {
	//backticks encodes the value of id into the key ID, here for json
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}
