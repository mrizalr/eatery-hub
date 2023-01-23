package models

import "github.com/gofiber/fiber/v2"

func ResponseOK(data interface{}) *Response {
	return &Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Errors: nil,
		Data:   data,
	}
}

func ResponseCreated(data interface{}) *Response {
	return &Response{
		Code:   fiber.StatusCreated,
		Status: "CREATED",
		Errors: nil,
		Data:   data,
	}
}

func ResponseBadRequest(errors interface{}) *Response {
	return &Response{
		Code:   fiber.StatusBadRequest,
		Status: "BAD REQUEST",
		Errors: errors,
		Data:   nil,
	}
}

func ResponseBadGateway(errors interface{}) *Response {
	return &Response{
		Code:   fiber.StatusBadGateway,
		Status: "BAD GATEWAY",
		Errors: errors,
		Data:   nil,
	}
}

func ResponseUnauthorized(errors interface{}) *Response {
	return &Response{
		Code:   fiber.StatusUnauthorized,
		Status: "UNAUTHORIZED",
		Errors: errors,
		Data:   nil,
	}
}
