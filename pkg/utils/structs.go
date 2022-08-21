package utils

type ReqData struct {
	Name   *string `json:"name,omitempty" form:"name,omitempty"`
	Author *string `json:"author,omitempty" form:"author,omitempty"`
}
