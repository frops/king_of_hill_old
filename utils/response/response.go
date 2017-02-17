package response

import (
	"encoding/json"
	"github.com/shagtv/go-api/httperrors"
	"net/http"
)

type JsonResponse struct {
	Data  interface{} `json:"data"`
	Error string `json:"error"`
}

func Send(res http.ResponseWriter, data interface{}, errorText string, responseStatus int) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(responseStatus)
	json.NewEncoder(res).Encode(JsonResponse{Data: data, Error: errorText})
}

func Ok(res http.ResponseWriter, data interface{}) {
	Send(res, data, "", http.StatusOK)
}

func Error(res http.ResponseWriter, err httperrors.IHttperror) {
	Send(res, nil, err.Error(), err.Code())
}
