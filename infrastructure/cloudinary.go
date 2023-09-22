package infrastructure

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
)

type Cloudinary struct {
	*cloudinary.Cloudinary
}

func NewCloudinary() *Cloudinary {
	cld, _ := credentials()
	return &Cloudinary{
		cld,
	}
}

func credentials() (*cloudinary.Cloudinary, context.Context) {
	// Add your Cloudinary credentials, set configuration parameter
	// Secure=true to return "https" URLs, and create a context
	//===================
	cld, _ := cloudinary.New()
	cld.Config.URL.Secure = true
	ctx := context.Background()
	return cld, ctx
}
