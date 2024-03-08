package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type response struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

type errorInfo struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func BodyParser(r *http.Request, body interface{}) error {
	return json.NewDecoder(r.Body).Decode(&body)
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	bytes, _ := json.MarshalIndent(data, "", "  ")

	w.Header().Set("Content-Type", "Application/json")
	w.Write(bytes)
}

func HandleBadRequestErrWithMessage(w http.ResponseWriter, err error, message string) error {
	if err == nil {
		return nil
	}

	log.Println(message+" ", err)
	w.WriteHeader(http.StatusBadRequest)
	writeJSON(w, response{Error: true,
		Data: errorInfo{
			Status:  http.StatusBadRequest,
			Message: message + ": " + err.Error(),
		}})
	return err
}

func HandleGrpcErrWithMessage(w http.ResponseWriter, err error, args ...interface{}) error {
	if err == nil {
		return nil
	}
	st, ok := status.FromError(err)
	if !ok || st.Code() == codes.Internal {
		// logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		writeJSON(w, response{Error: true,
			Data: errorInfo{
				Status:  http.StatusInternalServerError,
				Message: st.Message(),
			}})
		return err

	} else if st.Code() == codes.NotFound {
		// logger.Error(err)
		w.WriteHeader(http.StatusNotFound)
		writeJSON(w, response{Error: true,
			Data: errorInfo{
				Status:  http.StatusNotFound,
				Message: st.Message(),
			}})
		return err

	} else if st.Code() == codes.Unavailable {
		// logger.Error(err)
		w.WriteHeader(http.StatusBadGateway)
		writeJSON(w, response{Error: true,
			Data: errorInfo{
				Status:  http.StatusBadGateway,
				Message: st.Message(),
			}})
		return err

	} else if st.Code() == codes.AlreadyExists {
		// logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, response{Error: true,
			Data: errorInfo{
				Status:  http.StatusBadRequest,
				Message: st.Message(),
			}})
		return err

	} else if st.Code() == codes.InvalidArgument {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, response{Error: true,
			Data: errorInfo{
				Status:  http.StatusBadRequest,
				Message: st.Message(),
			}})
		return err

	} else if st.Code() == codes.DataLoss {
		// logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, response{Error: true,
			Data: errorInfo{
				Status:  http.StatusBadRequest,
				Message: st.Message(),
			}})
		return err

	} else if st.Code() == codes.PermissionDenied {
		// logger.Error(err)
		w.WriteHeader(http.StatusForbidden)
		writeJSON(w, response{Error: true,
			Data: errorInfo{
				Status:  http.StatusForbidden,
				Message: st.Message(),
			}})
		return err

	} else if strings.Contains("User blocked in user service", st.Message()) {
		// logger.Error(err)
		w.WriteHeader(http.StatusForbidden)
		writeJSON(w, response{Error: true,
			Data: errorInfo{
				Status:  http.StatusForbidden,
				Message: st.Message(),
			}})
		return err
	} else if st.Code() == codes.Unauthenticated {
		// logger.Error(err)
		w.WriteHeader(http.StatusUnauthorized)
		writeJSON(w, response{Error: true,
			Data: errorInfo{
				Status:  http.StatusUnauthorized,
				Message: st.Message(),
			}})
		return err

	}
	// logger.Error(err)
	w.WriteHeader(http.StatusInternalServerError)
	writeJSON(w, response{Error: true,
		Data: errorInfo{
			Status:  http.StatusInternalServerError,
			Message: st.Message(),
		}})

	return err
}

func WriteJSONWithSuccess(w http.ResponseWriter, data interface{}) {
	data = response{
		Error: false,
		Data:  data,
	}
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("WriteJSONWithSuccess err:", err)
	}

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func HandleInternalWithMessage(w http.ResponseWriter, err error, message string) error {
	if err == nil {
		return nil
	}

	log.Println(message+" ", err)
	w.WriteHeader(http.StatusInternalServerError)
	writeJSON(w, response{Error: true,
		Data: errorInfo{
			Status:  http.StatusInternalServerError,
			Message: message + " " + err.Error(),
		}})
	return err
}

func HandleBadRequestResponse(w http.ResponseWriter, message string) {
	log.Println(message)
	w.WriteHeader(http.StatusBadRequest)
	writeJSON(w, response{Error: true,
		Data: errorInfo{
			Status:  http.StatusBadRequest,
			Message: message,
		}})
}
