package pages

type ID string

// PageChangeMsg is used to change the current page
type PageChangeMsg struct {
	ID ID
}
