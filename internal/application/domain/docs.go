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
	ID      string
	Name    string
	Mime    string
	File    bool
	Public  bool
	Created string
	Grant   []string
}

type Doc struct {
	Mime string
	Data []byte
}

func NewDoc(mime string, data []byte) *Doc {
	return &Doc{Mime: mime, Data: data}
}

type DocFilters struct {
	Login   string
	Name    string
	Mime    string
	File    *bool
	Public  *bool
	Created string
}
