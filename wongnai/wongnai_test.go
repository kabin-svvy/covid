package wongnai

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type httpClienterMock struct{}

func (c *httpClienterMock) Do(req *http.Request) (*http.Response, error) {
	resBody, _ := json.Marshal(&WongnaiCovidCaseSummary{
		Data: []WongnaiCovidCaseSummaryData{
			{
				Confirmdate:    "2021-05-04",
				Age:            32,
				Genderen:       "Female",
				Nationen:       "Thailand",
				Province:       "Phang Nga",
				Provinceid:     38,
				Provinceen:     "Phang Nga",
				Statquarantine: 17,
			},
		},
	})

	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, string(resBody))
	}))
	return http.Get(ts.URL)
}

func TestGetWongnaiCovidCaseShouldBeSuccess(t *testing.T) {
	t.Run("Test get wongnai covid case should be success", func(t *testing.T) {
		expected := &WongnaiCovidCaseSummary{
			Data: []WongnaiCovidCaseSummaryData{
				{
					Confirmdate:    "2021-05-04",
					Age:            32,
					Genderen:       "Female",
					Nationen:       "Thailand",
					Province:       "Phang Nga",
					Provinceid:     38,
					Provinceen:     "Phang Nga",
					Statquarantine: 17,
				},
			},
		}
		clienter := &httpClienterMock{}
		wongnaier := New()
		actual, err := wongnaier.GetWongnaiCovidCase(clienter)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}
