package entities

type Token struct {
	Authorization string `json:"authorization"`
	ExpiresIn     int64  `json:"expires_in"`
	IsFirstLogin  bool   `json:"is_first_login"`
}
