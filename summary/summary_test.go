package summary

import (
	"covid/wongnai"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestClassifyByGroupShouldBeSuccess(t *testing.T) {
	t.Run("Test classify by age group should be right", func(t *testing.T) {
		expected := &SummaryCovidByGroup{
			AgeGroup: make(map[string]int),
			Province: make(map[string]int),
		}
		expected.AgeGroup["0-30"] = 2
		expected.AgeGroup["31-60"] = 2
		expected.AgeGroup["61+"] = 1
		expected.AgeGroup["N/A"] = 1
		expected.Province["N/A"] = 6

		wongnaiCase := &wongnai.WongnaiCovidCaseSummary{
			Data: []wongnai.WongnaiCovidCaseSummaryData{
				{
					Age: 0,
				},
				{
					Age: 30,
				},
				{
					Age: 31,
				},
				{
					Age: 60,
				},
				{
					Age: 61,
				},
				{
					Age: -1,
				},
			},
		}
		actual := classifyByGroup(wongnaiCase)
		assert.Equal(t, expected, actual)
	})

	t.Run("test classify by province group should be right", func(t *testing.T) {
		expected := &SummaryCovidByGroup{
			AgeGroup: make(map[string]int),
			Province: make(map[string]int),
		}
		expected.Province["Bangkok"] = 2
		expected.Province["Nan"] = 2
		expected.Province["Surin"] = 1
		expected.Province["N/A"] = 1
		expected.AgeGroup["0-30"] = 6

		wongnaiCase := &wongnai.WongnaiCovidCaseSummary{
			Data: []wongnai.WongnaiCovidCaseSummaryData{
				{
					Province: "Bangkok",
				},
				{
					Province: "Bangkok",
				},
				{
					Province: "Nan",
				},
				{
					Province: "Nan",
				},
				{
					Province: "Surin",
				},
				{
					Province: "",
				},
			},
		}
		actual := classifyByGroup(wongnaiCase)
		assert.Equal(t, expected, actual)
	})
}

type httpClienterMock struct{}

func (c *httpClienterMock) Do(req *http.Request) (*http.Response, error) {
	resBody, _ := json.Marshal(&wongnai.WongnaiCovidCaseSummary{
		Data: []wongnai.WongnaiCovidCaseSummaryData{
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

func TestGetSummaryShouldBeSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Test get summary should be get status 200", func(t *testing.T) {
		wongnaier := wongnai.New()
		summaryer := New(&wongnaier)

		h := &httpClienterMock{}

		rr := httptest.NewRecorder()
		router := gin.Default()

		router.GET("/covid/summary", summaryer.GetSummary(h))

		req, err := http.NewRequest(http.MethodGet, "/covid/summary", nil)

		router.ServeHTTP(rr, req)

		assert.NoError(t, err)
		assert.Equal(t, 200, rr.Code)
		result := rr.Result()
		body, err := ioutil.ReadAll(result.Body)
		assert.NoError(t, err)

		expected := &SummaryCovidByGroup{
			Province: make(map[string]int),
			AgeGroup: make(map[string]int),
		}

		expected.AgeGroup["31-60"] = 1
		expected.Province["Phang Nga"] = 1

		actual := &SummaryCovidByGroup{}

		err = json.Unmarshal([]byte(body), &actual)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}
