package form

type Avatar struct {
	UID uint `json:"uid" form:"uid" binding:"required,min=1"`
}
