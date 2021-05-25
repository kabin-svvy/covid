package wongnai

import (
	"covid/interfaces"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	WongnaiCovidCasePath = "https://static.wongnai.com/devinterview/covid-cases.json"
)

type Wongnaier struct {
}

func New() Wongnaier {
	return Wongnaier{}
}

type WongnaiCovidCaseSummary struct {
	Data []WongnaiCovidCaseSummaryData `json:"Data"`
}

type WongnaiCovidCaseSummaryData struct {
	Confirmdate    string `json:"ConfirmDate"`
	No             string `json:"No"`
	Age            int    `json:"Age"`
	Gender         string `json:"Gender"`
	Genderen       string `json:"GenderEn"`
	Nation         string `json:"Nation"`
	Nationen       string `json:"NationEn"`
	Province       string `json:"Province"`
	Provinceid     int    `json:"ProvinceId"`
	District       string `json:"District"`
	Provinceen     string `json:"ProvinceEn"`
	Statquarantine int    `json:"StatQuarantine"`
}

func (w *Wongnaier) GetWongnaiCovidCase(clienter interfaces.HttpClienter) (*WongnaiCovidCaseSummary, error) {
	req, err := http.NewRequest(http.MethodGet, WongnaiCovidCasePath, nil)
	if err != nil {
		return nil, err
	}

	res, err := clienter.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	result := &WongnaiCovidCaseSummary{}

	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
