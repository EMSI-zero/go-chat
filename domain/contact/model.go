package contact

type Contact struct {
	ID          int64  `json:"-"`
	UserID      int64  `json:"userId"`
	ContactID   int64  `json:"contact_id"`
	ContactName string `json:"contact_name"`
}

type AddContectRequest struct {
	ContactID   int64  `json:"contactId"`
	ContactName string `json:"contactName"`
}

type RemoveContactRequest struct {
	ContactID int64 `json:"contactId"`
}
