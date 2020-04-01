package jandanTreeholeRSS

import (
	"github.com/gorilla/feeds"
	"github.com/iochen/jandanTreeholeRSS/jandan/treehole"
	"log"
	"strconv"
	"strings"
	"time"
)

func GetFeed(n int) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	feed := &feeds.Feed{
		Title:   "jandan Treehole",
		Link:    &feeds.Link{Href: "https://jandan.net/treehole"},
		Created: time.Now(),
	}

	items := make([]*feeds.Item, n)

	id := treehole.LatestID()
	for i := 0; i < n; i++ {
		log.Println("fetching", id)
		th, err := treehole.GetFromID(id)
		if err != nil {
			return "", err
		}
		items[i] = &feeds.Item{
			Title:   "Treehole #" + strconv.Itoa(int(th.ID)),
			Link:    &feeds.Link{Href: "https://jandan.net/t/" + strconv.Itoa(int(th.ID))},
			Author:  &feeds.Author{Name: th.Author},
			Content: strings.ReplaceAll(th.String(), "\n", "<br />"),
		}
		id = th.Neighbour.Prev
	}

	feed.Items = items

	return feed.ToRss()
}
