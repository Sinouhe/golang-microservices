package controllers

import (
	"net/http"
	"strconv"
	"golang-microservices/mvc/services"
	"golang-microservices/mvc/utils"
	"encoding/json"
)

func GetUsers(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "userId must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		handleErrorResponse(resp, apiErr)
		return
	}
	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		handleErrorResponse(resp, apiErr)
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}

func handleErrorResponse (resp http.ResponseWriter, apiErr *utils.ApplicationError) {
	jsonValue, _ := json.Marshal(apiErr)
	resp.WriteHeader(apiErr.StatusCode)
	resp.Write(jsonValue)
}
