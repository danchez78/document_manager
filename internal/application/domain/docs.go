package domain

type DocInput struct {
	Data   []byte
	Name   string
	File   bool
	Public bool
	Mime   string
	Grant  []string
}

type DocInfo struct {
	ID     string
	Name   string
	File   bool
	Public bool
	Mime   string
	Grant  []string
}
