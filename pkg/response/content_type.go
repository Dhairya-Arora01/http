package response

// ContentType represents a MIME content type.
type ContentType string

// String represents the underlying string representation of the ContentType.
func (c ContentType) String() string {
	return string(c)
}

const (
	// ContentTypeNil represents that there is no content.
	ContentTypeNil ContentType = ""

	// ContentTypeJSON represents application/json content.
	ContentTypeJSON ContentType = "application/json"

	// ContentTypeXML represents application/xml content.
	ContentTypeXML ContentType = "application/xml"

	// ContentTypeHTML represents text/html content.
	ContentTypeHTML ContentType = "text/html"

	// ContentTypePlain represents text/plain content.
	ContentTypePlain ContentType = "text/plain"

	// ContentTypeFormURLEncoded represents application/x-www-form-urlencoded content.
	ContentTypeFormURLEncoded ContentType = "application/x-www-form-urlencoded"

	// ContentTypeMultipartFormData represents multipart/form-data content.
	ContentTypeMultipartFormData ContentType = "multipart/form-data"

	// ContentTypeJavaScript represents application/javascript content.
	ContentTypeJavaScript ContentType = "application/javascript"

	// ContentTypeCSS represents text/css content.
	ContentTypeCSS ContentType = "text/css"

	// ContentTypeCSV represents text/csv content.
	ContentTypeCSV ContentType = "text/csv"

	// ContentTypeOctetStream represents application/octet-stream content.
	ContentTypeOctetStream ContentType = "application/octet-stream"

	// ContentTypePDF represents application/pdf content.
	ContentTypePDF ContentType = "application/pdf"

	// ContentTypeZIP represents application/zip content.
	ContentTypeZIP ContentType = "application/zip"

	// ContentTypeWebP represents image/webp content.
	ContentTypeWebP ContentType = "image/webp"

	// ContentTypePNG represents image/png content.
	ContentTypePNG ContentType = "image/png"

	// ContentTypeJPEG represents image/jpeg content.
	ContentTypeJPEG ContentType = "image/jpeg"

	// ContentTypeGIF represents image/gif content.
	ContentTypeGIF ContentType = "image/gif"

	// ContentTypeSVG represents image/svg+xml content.
	ContentTypeSVG ContentType = "image/svg+xml"

	// ContentTypeMP4 represents video/mp4 content.
	ContentTypeMP4 ContentType = "video/mp4"

	// ContentTypeWebM represents video/webm content.
	ContentTypeWebM ContentType = "video/webm"

	// ContentTypeMP3 represents audio/mpeg content.
	ContentTypeMP3 ContentType = "audio/mpeg"

	// ContentTypeWAV represents audio/wav content.
	ContentTypeWAV ContentType = "audio/wav"

	// ContentTypeOGG represents audio/ogg content.
	ContentTypeOGG ContentType = "audio/ogg"
)
