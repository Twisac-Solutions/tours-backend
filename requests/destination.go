package requests

import "mime/multipart"

type CreateDestinationRequest struct {
	Name        string                `form:"name"`
	Description string                `form:"description"`
	Region      string                `form:"region"`
	Country     string                `form:"country"`
	CoverImage  *multipart.FileHeader `form:"coverImage"`
}

type UpdateDestinationRequest struct {
	Name        string                `form:"name"`
	Description string                `form:"description"`
	Region      string                `form:"region"`
	Country     string                `form:"country"`
	CoverImage  *multipart.FileHeader `form:"coverImage"`
}
