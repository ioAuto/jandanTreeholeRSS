package treehole

import (
	"encoding/json"
	"github.com/iochen/jandanTreeholeRSS/jandan/common/network"
	"strconv"
	"time"
)

type sTime time.Time

type commentJson struct {
	Comments []struct {
		ID      int    `json:"comment_ID"`
		Date    sTime  `json:"comment_date_int"`
		Author  string `json:"comment_author"`
		Content string `json:"comment_content"`
	} `json:"tucao"`
}

func (st *sTime) UnmarshalJSON(input []byte) error {
	i, err := strconv.ParseInt(string(input), 10, 64)
	if err != nil {
		return err
	}
	*st = sTime(time.Unix(i, 0))
	return nil
}

func GetComments(id ID) (*[]Comment, error) {
	bytes, err := network.HttpGetWithUA("https://jandan.net/tucao/all/" + strconv.Itoa(int(id)))
	if err != nil {
		return nil, err
	}

	var cmtJson commentJson
	if err := json.Unmarshal(bytes, &cmtJson); err != nil {
		return nil, err
	}
	cmts := make([]Comment, len(cmtJson.Comments))
	for k, v := range cmtJson.Comments {
		cmts[k] = Comment{
			ID:      v.ID,
			Date:    time.Time(v.Date),
			Author:  v.Author,
			Content: trimUesless(v.Content),
		}
	}

	return &cmts, err
}
