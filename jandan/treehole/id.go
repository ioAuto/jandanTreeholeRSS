package treehole

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/iochen/jandanTreeholeRSS/jandan/common/network"
	"strconv"
)

func LatestID() ID {
	r, err := network.HttpGetWithUA("https://jandan.net/treehole")
	if err != nil {
		return 0
	}

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return 0
	}

	i, err := strconv.Atoi(doc.Find("span.righttext").First().Text())
	if err != nil {
		return 0
	}
	return ID(i)
}
