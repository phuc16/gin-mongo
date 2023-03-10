package model

type Code struct {
	Success      int `yaml:"success"`
	BadRequest   int `yaml:"bad_request"`
	Unauthorized int `yaml:"unauthorized"`
	Forbidden    int `yaml:"forbidden"`
	NotFound     int `yaml:"not_found"`
	Conflict     int `yaml:"conflict"`
	ServerError  int `yaml:"server_error"`
}

type ResponseCode struct {
	ResCode Code `yaml:"response_code"`
}
