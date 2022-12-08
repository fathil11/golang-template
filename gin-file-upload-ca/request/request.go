package request

import "mime/multipart"

type Request struct {
	File      *multipart.FileHeader `form:"file" binding:"required"`
	Condition string                `form:"condition" binding:"required,oneof='customer logo'"`
}
