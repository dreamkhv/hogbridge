package request

type SendgridMessage struct {
	Personalizations []Personalization `json:"personalizations" validate:"required"`
	From             EmailAddress      `json:"from" validate:"required"`
	Subject          string            `json:"subject" validate:"required"`
	Content          []Content         `json:"content" validate:"required"`
}

type Personalization struct {
	To []EmailAddress `json:"to" validate:"required"`
}

type EmailAddress struct {
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name" validate:"required"`
}

type Content struct {
	Type  string `json:"type" validate:"required"`
	Value string `json:"value" validate:"required"`
}
