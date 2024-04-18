package constant

var WhitelistPictureExtension = map[string]struct{}{
	".png":  {},
	".jpg":  {},
	".jpeg": {},
	".webp": {},
}

var AcceptedImageContentType = map[string]struct{}{
	"image/png":  {},
	"image/jpg":  {},
	"image/jpeg": {},
	"image/webp": {},
}
