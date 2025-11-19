package provisionclient

type ContactsRole struct {
	ID  PVID   `json:"id,omitempty"`
	Tag string `json:"tag,omitempty"`
}

type ContactsContact struct {
	ID         PVID           `json:"id,omitempty"`
	Name       string         `json:"name,omitempty"`
	Slug       string         `json:"slug,omitempty"`
	Type       string         `json:"string,omitempty"`
	ParentID   PVID           `json:"parent_id,omitempty"`
	CategoryID PVID           `json:"category_id,omitempty"`
	Date       string         `json:"date,omitempty"`
	Modified   string         `json:"modified,omitempty"`
	Attrs      map[string]any `json:"attrs,omitempty"`
}
