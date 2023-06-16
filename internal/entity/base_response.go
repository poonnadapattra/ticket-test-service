package entity

type BaseResponse struct {
	Code    int         `json:"code"`
	Massage string      `json:"massage"`
	Data    interface{} `json:"data"`
}

type BaseResponsePagging struct {
	Code    int         `json:"code"`
	Massage string      `json:"massage"`
	Data    interface{} `json:"data"`
	Pagging Pagging     `json:"pagging"`
}

type Pagging struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
}
