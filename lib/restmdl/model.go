package restmdl

type RequestMetaData struct {
	UserId *string

	Ip        *string
	UserAgent *string
}

type BaseRS struct {
	Version string `json:"version"`
	Context string `json:"context"`

	Data  interface{} `json:"data"`
	Error error       `json:"error"`
}
