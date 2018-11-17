package faker

// MimeType will generate a random mime file type
func (f *Faker) MimeType() string {
	return f.getRandValue([]string{"file", "mime_type"})
}

// Extension will generate a random file extension
func (f *Faker) Extension() string {
	return f.getRandValue([]string{"file", "extension"})
}
