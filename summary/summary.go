package summary

import (
	"covid/interfaces"
	"covid/wongnai"
	"fmt"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type SummaryCovidByGroup struct {
	Province map[string]int
	AgeGroup map[string]int
}

type Summaryer interface {
	GetSummary(interfaces.HttpClienter) gin.HandlerFunc
}

type summaryer struct {
	wongnaier *wongnai.Wongnaier
	logger    *logrus.Entry
}

func New(wongnaier *wongnai.Wongnaier) Summaryer {
	return &summaryer{
		wongnaier: wongnaier,
	}
}

func (s *summaryer) GetSummary(httpClient interfaces.HttpClienter) gin.HandlerFunc {
	return func(c *gin.Context) {
		wongnaiSummary, err := s.wongnaier.GetWongnaiCovidCase(httpClient)
		if err != nil {
			s.logger.WithError(err)
			return
		}
		result := classifyByGroup(wongnaiSummary)
		sortProvince(result)
		c.JSON(200, result)
	}
}

func classifyByGroup(wongnaiSummary *wongnai.WongnaiCovidCaseSummary) *SummaryCovidByGroup {
	r := &SummaryCovidByGroup{}
	r.AgeGroup = make(map[string]int)
	r.Province = make(map[string]int)
	for _, v := range wongnaiSummary.Data {
		switch age := v.Age; {
		case age >= 0 && age <= 30:
			r.AgeGroup["0-30"] = r.AgeGroup["0-30"] + 1

		case age >= 31 && age <= 60:
			r.AgeGroup["31-60"] = r.AgeGroup["31-60"] + 1

		case age >= 61:
			r.AgeGroup["61+"] = r.AgeGroup["61+"] + 1

		default:
			r.AgeGroup["N/A"] = r.AgeGroup["N/A"] + 1
		}

		if v.Province == "" {
			r.Province["N/A"] = r.Province["N/A"] + 1
		} else {
			r.Province[v.Province] = r.Province[v.Province] + 1
		}
	}
	return r
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func sortProvince(summaryCovidGroup *SummaryCovidByGroup) {
	p := make(PairList, len(summaryCovidGroup.Province))

	i := 0

	for k, v := range summaryCovidGroup.Province {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)
	fmt.Printf("%v\n", p)

	summaryCovidGroup.Province = make(map[string]int)

	for _, v := range p {
		summaryCovidGroup.Province[v.Key] = v.Value
	}
	// fmt.Printf("%v", summaryCovidGroup.Province)
}
