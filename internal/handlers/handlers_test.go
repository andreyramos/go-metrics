package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andreyramos/go-metrics/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestPostHandler(t *testing.T) {
	type want struct {
		statusCode int
	}
	tests := []struct {
		name    string
		request string
		want    want
	}{
		{
			name: "test #1",
			want: want{
				statusCode: 200,
			},
			request: "/update/gauge/Alloc/31.55",
		},
		{
			name: "test #2",
			want: want{
				statusCode: 200,
			},
			request: "/update/counter/PollCount/1",
		},
		{
			name: "test #3",
			want: want{
				statusCode: 400,
			},
			request: "/update/gauge/Alloc/xxx",
		},
		{
			name: "test #4",
			want: want{

				statusCode: 501,
			},
			request: "/update/alarm/smoke/1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, tt.request, nil)
			w := httptest.NewRecorder()

			cr := chi.NewRouter()
			cr.Post("/update/{type}/{name}/{value}", Post(storage.NewMemStorage()))
			cr.ServeHTTP(w, r)
			res := w.Result()
			defer res.Body.Close()
			assert.Equal(t, tt.want.statusCode, res.StatusCode)

		})
	}
}
