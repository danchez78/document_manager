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
	Grants  []string
}

func (i *DocInfo) HasAccess(userLogin string) bool {
	if i.Public {
		return true
	}
	for _, grant := range i.Grants {
		if grant == userLogin {
			return true
		}
	}
	return false
}

func (i *DocInfo) CanDelete(userLogin string) bool {
	for _, grant := range i.Grants {
		if grant == userLogin {
			return true
		}
	}
	return false
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
