package entity

type ListFavorite struct {
	SharedEntity
	ChooserID string `json:"chooser_id"`
	ListID    string `json:"list_id"`
}

func NewListFavorite(chooserID, listID string) *ListFavorite {
	return &ListFavorite{
		SharedEntity: *NewSharedEntity(),
		ChooserID:    chooserID,
		ListID:       listID,
	}
}

func (lf *ListFavorite) Equals(other *ListFavorite) bool {
	return lf.ChooserID == other.ChooserID && lf.ListID == other.ListID
}
