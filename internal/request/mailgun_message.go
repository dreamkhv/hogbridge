package request

import "mime/multipart"

type MailgunMessage struct {
	From       string                `form:"from" validate:"required,email"`
	To         string                `form:"to" validate:"required,email"`
	Subject    string                `form:"subject" validate:"required"`
	Text       string                `form:"text" validate:"required_without=Html"`
	Html       string                `form:"html" validate:"required_without=Text"`
	Attachment *multipart.FileHeader `form:"attachment"`
}
