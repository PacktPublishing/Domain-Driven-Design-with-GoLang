package chapter2

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler interface {
	IsUserPaymentActive(ctx context.Context, userID string) bool
}

type UserActiveResponse struct {
	IsActive bool
}

func router(u UserHandler) {
	m := mux.NewRouter()
	m.HandleFunc("/user/{userID}/payment/active", func(writer http.ResponseWriter, request *http.Request) {
		// check auth, etc

		uID := mux.Vars(request)["userID"]
		if uID == "" {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		isActive := u.IsUserPaymentActive(request.Context(), uID)

		b, err := json.Marshal(UserActiveResponse{IsActive: isActive})
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, _ = writer.Write(b)
	}).Methods(http.MethodGet)
}
