package models

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Errors interface{} `json:"errors"`
	Data   interface{} `json:"data"`
}

type ResponseWithMeta struct {
	Response
	Meta `json:"meta"`
}

type Meta struct {
	TotalPages      int  `json:"total_pages"`
	CurrentPage     int  `json:"current_page"`
	TotalItems      int  `json:"total_items"`
	HasNextPage     bool `json:"has_next_page"`
	HasPreviousPage bool `json:"has_previous_page"`
}
