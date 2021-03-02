package application_service

import "time"

// ApplicationModel ...
type ApplicationModel struct {
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateApplicationModel ...
type CreateApplicationModel struct {
	ID   string `json:"id" swaggerignore:"true"`
	Body string `json:"body"`
}

// ApplicationCreatedModel ...
type ApplicationCreatedModel struct {
	ID string `json:"id"`
}

// UpdateApplicationModel ...
type UpdateApplicationModel struct {
	ID   string `json:"id" swaggerignore:"true"`
	Body string `json:"body"`
}

// ApplicationUpdatedModel ...
type ApplicationUpdatedModel struct {
	ID   string `json:"id"`
	Body string `json:"body"`
}

// DeleteApplicationModel ...
type DeleteApplicationModel struct {
	ID string `json:"id"`
}

// ApplicationListModel ...
type ApplicationListModel struct {
	Count        int                `json:"count"`
	Applications []ApplicationModel `json:"applications"`
}

// ApplicationQueryParamModel ...
type ApplicationQueryParamModel struct {
	Search      string `json:"search"`
	Order       string `json:"order" enums:"id,body,created_at, updated_at"`
	Arrangement string `json:"arrangement" enums:"asc,desc"`
	Offset      int    `json:"offset" default:"0"`
	Limit       int    `json:"limit" default:"10"`
}
