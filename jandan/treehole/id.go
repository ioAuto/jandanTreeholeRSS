package treehole

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/iochen/jandanTreeholeRSS/jandan/common/network"
	"strconv"
)

func LatestID() ID {
	b, err := network.HttpGetWithUA("https://jandan.net/treehole")
	if err != nil {
		return 0
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
	if err != nil {
		return 0
	}

	i, err := strconv.Atoi(doc.Find("span.righttext").First().Text())
	if err != nil {
		return 0
	}
	return ID(i)
}
