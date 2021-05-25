package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerWithHealthyShouldBeSuccess(t *testing.T) {
	t.Run("Test new request to /healthy should be success", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/healthy", nil)
		assert.NoError(t, err)

		srv := setUpServer()
		srv.ServeHTTP(rr, req)
		assert.Equal(t, 200, rr.Code)
		result := rr.Result()
		body, err := ioutil.ReadAll(result.Body)
		assert.NoError(t, err)
		actual := &health{}
		json.Unmarshal([]byte(body), &actual)
		assert.Equal(t, "OK", actual.Status)
	})
}
