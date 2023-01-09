package models

// POST /api/images/:id
type RunImageResponseSuccess struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  int    `json:"result"`
	ID      int    `json:"id"`
}

type RunImageResponseError struct {
	Jsonrpc string `json:"jsonrpc"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	ID int `json:"id"`
}

type RunImageResponseNotification struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
}

// GET /api/images

type GetImagesSuccess struct {
	Images []string `json:"images"`
}
