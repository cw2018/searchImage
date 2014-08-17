package controllers

import (
	"github.com/revel/revel"
  "io/ioutil"
  "log"
	"net/http"
	"encoding/json"
	"strings"
)

type App struct {
	*revel.Controller
}

type data struct {
	GsearchResultClass string
	Width string `json:"width"`
	Height string `json:"height"`
	ImageId string `json:"imageId"`
	TbWidth string `json:"tbWidth"`
	TbHeight string `json:"tbHeight"`
	UnescapedUrl string `json:"unescapedUrl"`
	Url string `json:"url"`
	VisibleUrl string `json:"visibleUrl"`
	Title string `json:"title"`
	TitleNoFormatting string `json:"titleNoFormatting"`
	OriginalContextUrl string `json:"originalContextUrl"`
	Content string `json:"content"`
	ContentNoFormatting string `json:"contentNoFormatting"`
	TbUrl string `json:"tbUrl"`
}

type results struct {
	Results []data `json:"results"`
}

type Response struct {
	ResponseData results `json:"responseData"`
}

func (c App) SearchImage(searchWord string) revel.Result {
	rsz := "large"
	hl := "ja"
	size := "medium"

	apiurl := "http://ajax.googleapis.com/ajax/services/search/images?"
	apiurl = apiurl + "q="+ searchWord +"&v=1.0&rsz="+ rsz +"&hl="+ hl + "&imgsz="+ size

	res, err := http.Get(apiurl)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	jsonstream := string(robots[:])
	dec := json.NewDecoder(strings.NewReader(jsonstream))
	var d Response
	dec.Decode(&d)
	resboby := d.ResponseData.Results
	return c.Render(resboby)
}
