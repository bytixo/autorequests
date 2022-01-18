package app

import "time"

//  might add more options soon
type Config struct {
	Exclude []string
}
type Request struct {
	Method      string    `json:"method"`
	URL         string    `json:"url"`
	HTTPVersion string    `json:"httpVersion"`
	Headers     []Headers `json:"headers"`
	Cookies     []Cookies `json:"cookies"`
	PostData    PostData  `json:"postData,omitempty"`
}
type Headers struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Cookies struct {
	Name     string    `json:"name"`
	Value    string    `json:"value"`
	Path     string    `json:"path"`
	Domain   string    `json:"domain"`
	Expires  time.Time `json:"expires"`
	HTTPOnly bool      `json:"httpOnly"`
	Secure   bool      `json:"secure"`
	SameSite string    `json:"sameSite,omitempty"`
}
type PostData struct {
	MimeType string `json:"mimeType"`
	Text     string `json:"text"`
}
