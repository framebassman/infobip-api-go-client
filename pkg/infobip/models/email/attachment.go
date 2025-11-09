package email

// Attachment defines model for Attachment.
type Attachment struct {
	// ContentType Content type of the attachment.
	ContentType *string `json:"contentType,omitempty"`

	// FileName Name of the attachment.
	FileName *string `json:"fileName,omitempty"`

	// Id Id of the attachment.
	Id string `json:"id"`

	// Size Size of the attachment in bytes.
	Size *int32 `json:"size,omitempty"`

	// Url URL of the attachment.
	Url *string `json:"url,omitempty"`
}

// AttachmentList defines model for AttachmentList.
type AttachmentList = []Attachment
