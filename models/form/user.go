package form

type Avatar struct {
	MID uint `json:"mid" form:"mid" binding:"required,min=1"`
}
