package models

// Profile represents user profile information
type Profile struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Bio       string `json:"bio"`
	Limit     int    `json:"limit"`
	IsPremium bool   `json:"is_premium"`
}
