package dto

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

type DocFilters struct {
	Login   string
	Name    string
	Mime    string
	File    *bool
	Public  *bool
	Created string
}
