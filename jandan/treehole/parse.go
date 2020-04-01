package treehole

import (
	"encoding/json"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/iochen/jandanTreeholeRSS/jandan/common/network"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var ErrNotTreehole = errors.New("not a treehole")

type neighbourAPI struct {
	Data Neighbour `json:"data"`
}

func ParseFromHtml(r io.Reader) (*Treehole, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return &Treehole{}, err
	}

	// check if it is a treehole
	title := doc.Find("div#content h1.title").Text()
	title = strings.TrimSpace(title)
	if !strings.HasPrefix(title, "树洞") {
		return &Treehole{}, ErrNotTreehole
	}

	idResult := regexp.MustCompile(`No\.(\d+)`).FindAllStringSubmatch(title, -1)
	id, err := strconv.Atoi(idResult[0][1])
	if err != nil {
		return &Treehole{}, err
	}

	jsScripts := doc.Find("script").Text()
	apiResult := regexp.MustCompile(`/api/comment/neighbor/\d+/\d+`).FindString(jsScripts)
	reader, err := network.HttpGetWithUA("https://jandan.net" + apiResult)
	if err != nil {
		return &Treehole{}, err
	}
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return &Treehole{}, err
	}
	neighbour := &neighbourAPI{}
	if err := json.Unmarshal(b, neighbour); err != nil {
		return &Treehole{}, err
	}

	var content string
	doc.Find("div.comment-topic p").Each(func(i int, selection *goquery.Selection) {
		content += selection.Text() + "\n"
	})

	th := &Treehole{
		ID:        ID(id),
		Neighbour: neighbour.Data,
		Author:    doc.Find("div.comment-topic b").Text(),
		Content:   trimUesless(content),
		Comments:  GetComments(ID(id)),
	}

	return th, nil
}

func GetFromURL(url string) (*Treehole, error) {
	body, err := network.HttpGetWithUA(url)
	if err != nil {
		return nil, err
	}

	return ParseFromHtml(body)
}

func GetFromID(id ID) (*Treehole, error) {
	return GetFromURL("https://jandan.net/t/" + strconv.Itoa(int(id)))
}

func trimUesless(str string) string {
	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, "</a>", "")
	return regexp.MustCompile(`<a href="#tucao-\d+" data-id="\d+" class="tucao-link">`).ReplaceAllString(str, "")
}
