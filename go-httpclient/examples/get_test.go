package examples

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/ken5scal/go-httpclient/gohttp"
)

func TestGetEndpoints(t *testing.T) {
	mock1 := gohttp.Mock{
		Method: http.MethodGet,
		Url:    "https://api.github.com",
		Error:  errors.New("timeout getting github endpoints"),
	}
	mock2 := gohttp.Mock{
		Method:             http.MethodGet,
		Url:                "https://api.github.com",
		ResponseStatusCode: http.StatusOK,
		ResponseBody:       `{"current_user_url": 123}`,
		Error:              errors.New("timeout getting github endpoints"),
	}
	mock3 := gohttp.Mock{
		Method:             http.MethodGet,
		Url:                "https://api.github.com",
		ResponseStatusCode: http.StatusOK,
		ResponseBody:       `{"current_user_url": "https://api.github.com/user"}`,
		Error:              nil,
	}

	var mock3want Endpoints
	if err := json.Unmarshal([]byte(mock3.ResponseBody), &mock3want); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		mock    gohttp.Mock
		want    *Endpoints
	}{
		{
			name:    "TestErrorFetchingFromGithub",
			mock:    mock1,
			want:    nil,
		},
		{
			name:    "TestErrorUnmarshalResponseBody",
			mock:    mock2,
			want:    nil,
		},
		{
			name:    "TestNoError",
			mock:    mock3,
			want:    &mock3want,
		},
	}

	gohttp.StartMockServer()

	for _, tt := range tests {
		gohttp.AddMock(tt.mock)
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEndpoints()
			if err != nil && tt.mock.Error != nil {
				if got != nil {
					t.Errorf("no endpoints expected, wantErr %v", tt.mock.Error)
				}

				if !errors.Is(err, tt.mock.Error) {
					t.Errorf("GetEndpoints() error = %v, wantErr %v", err, tt.mock.Error)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEndpoints() = %v, want %v", got, tt.want)
			}
		})
	}
}