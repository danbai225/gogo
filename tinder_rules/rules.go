package tinder_rules

import (
	"fmt"

	"github.com/goccy/go-json"
)

type tinderRules struct {
	Ver  string `json:"ver"`
	Tag  string `json:"tag"`
	Data []data `json:"data"`
}
type data struct {
	Id        int        `json:"id"`
	Power     int        `json:"power"`
	Name      string     `json:"name"`
	Procname  string     `json:"procname"`
	Treatment int        `json:"treatment"`
	Policies  []policies `json:"policies"`
}
type policies struct {
	Montype    int    `json:"montype"`
	ActionType int    `json:"action_type"`
	ResPath    string `json:"res_path"`
}

func NewRules(name string, procedure []string, path [][]string) []byte {
	rules := tinderRules{
		Ver:  "5.0",
		Tag:  "hipsuser",
		Data: []data{},
	}
	for i, v := range procedure {
		arr := make([]policies, 0)
		for _, p := range path[i] {
			arr = append(arr, policies{
				Montype:    1,
				ActionType: 15,
				ResPath:    p,
			})
		}
		rules.Data = append(rules.Data, data{
			Id:        99,
			Power:     0,
			Name:      fmt.Sprintf("gogo%d", i),
			Procname:  v,
			Treatment: 3,
			Policies:  arr,
		})
	}

	marshal, _ := json.Marshal(rules)
	return marshal
}
