package model

type ProductPropertyValueInput struct {
	Code  string `json:"Code"`
	Value string `json:"Value"`
}

type ProductFillable struct {
	Name        string                      `json:"Name"`
	Description string                      `json:"Description"`
	ImageSrc    string                      `json:"ImageSrc"`
	Price       float64                     `json:"Price"`
	CategoryID  uint                        `json:"CategoryID"`
	Properties  []ProductPropertyValueInput `json:"Properties"`
}
