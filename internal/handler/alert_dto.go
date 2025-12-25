package handler

type CreateAlertReq struct {
	Symbol      string  `json:"symbol" validate:"required,oneof='BTCUSDT ETHUSDT DOGEUSDT BNBUSDT'"`
	TargetPrice float64 `json:"target_price" validate:"required"`
}

type JSONResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}


