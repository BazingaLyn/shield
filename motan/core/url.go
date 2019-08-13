package core

type URL struct {
	Protocol   string            `json:"protocol"`
	Host       string            `json:"host"`
	Port       int               `json:"port"`
	Path       string            `json:"path"`
	Group      string            `json:"group"`
	Parameters map[string]string `json:"parameters"`

	address  string
	identity string
}
