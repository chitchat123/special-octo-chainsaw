package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func Test_primeHandler(t *testing.T) {
	tests := []struct {
		name       string
		req        []interface{}
		want       []bool
		wantError  string
		wantStatus int
	}{
		{"1", []any{1, 2}, []bool{false, true}, "", http.StatusOK},
		{"2", []any{1, 2, 3, 4, 5, 6, 6}, []bool{false, true, true, false, true, false, false}, "", http.StatusOK},
		{"3", []any{1, 2, "a"}, nil, "the given input is invalid. Element on index 2 is not a number", http.StatusBadRequest},
		{"4", []any{1, 2.11}, nil, "the given input is invalid. Element on index 1 is not a number", http.StatusBadRequest},
		{"5", []any{1, 2, []int{1, 2}}, nil, "the given input is invalid. Element on index 2 is not a number", http.StatusBadRequest},
	}
	// Create a new router instance
	r := gin.New()

	l := log.New(os.Stdout, "TestPrimeHandler: ", log.LstdFlags)

	r.POST("/", primeHandler(l))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.req)
			if err != nil {
				t.Errorf("Failed to create a new HTTP request: %v", err)
			}
			request, err := http.NewRequest("POST", "/", bytes.NewBuffer(body))
			if err != nil {
				return
			}
			response := httptest.NewRecorder()
			r.ServeHTTP(response, request)
			if tt.want != nil {
				var output []interface{}
				err = json.Unmarshal(response.Body.Bytes(), &output)
				if err != nil {
					t.Errorf("Failed to unmarshal the response body: %v", err)
				}

				tmp := output
				convertedSlice := make([]bool, len(tmp))
				for i, v := range tmp {
					convertedSlice[i] = v.(bool)
				}

				if !reflect.DeepEqual(tt.want, convertedSlice) || tt.wantStatus != response.Code {
					t.Errorf("toInt() gotRes = %#v, want %#v", convertedSlice, tt.want)
				}
			} else {
				var output gin.H
				err = json.Unmarshal(response.Body.Bytes(), &output)
				if err != nil {
					t.Errorf("Failed to unmarshal the response body: %v", err)
				}
				if output["error"] != tt.wantError {
					t.Errorf("toInt() gotRes = %#v, want %#v", output["error"], tt.wantError)
				}
			}
		})
	}
}
