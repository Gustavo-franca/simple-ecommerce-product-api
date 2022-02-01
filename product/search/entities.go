package search

type (
	Params struct {
		Description string   `json:"description"`
		Title       string   `json:"title"`
		IDs         []string `json:"ids"`
	}
)
