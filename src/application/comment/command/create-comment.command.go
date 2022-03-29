package command

type CommandCreateComment struct {
	Title         string `json:"Title" binding:"required,min=5"`
	Content       string `json:"Content" binding:"required,min=5"`
	PublicationId int    `json:"PublicationId" binding:"required"`
	UserId        int    `json:"UserId"`
}
