package response

type GeneralResponse struct {
	Meta `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"msg"`
}
