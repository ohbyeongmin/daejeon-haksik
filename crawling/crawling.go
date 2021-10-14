package crawling

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type DaejeonHRCUrl struct {
	listURL     string
	infoURL     string
	downloadURL string
	bbsId       string
	mi          string
	nttSn       string
	fileKey     string
}

var dURL *DaejeonHRCUrl

func init() {
	dURL = &DaejeonHRCUrl{
		listURL:     "https://www.dju.ac.kr/hrc/na/ntt/selectNttList.do",
		infoURL:     "https://www.dju.ac.kr/hrc/na/ntt/selectNttInfo.do",
		downloadURL: "https://www.dju.ac.kr/common/nttFileDownload.do",
		bbsId:       "2126",
		mi:          "4597",
	}
	dURL.setNttSn()
	dURL.setFileKey()
}

func (d *DaejeonHRCUrl) setNttSn() {
	url := d.getDietListURL()
	res, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()

	doc.Find(".nttInfoBtn").Each(func(i int, s *goquery.Selection) {
		_, ok := s.Attr("href")
		if ok {
			data, _ := s.Attr("data-id")
			d.nttSn = data
		}
	})
}

func (d *DaejeonHRCUrl) setFileKey() {
	var downloadLink string
	url := dURL.getDietInfoURL()
	res, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()

	doc.Find(".file").Children().Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			downloadLink = href
		}
	})

	d.fileKey = strings.Split(strings.Split(strings.Split(downloadLink, "?")[1], "&")[0], "=")[1]
}

func (d DaejeonHRCUrl) getDietListURL() string {
	return fmt.Sprintf("%s?mi=%s&bbsId=%s", d.listURL, d.mi, d.bbsId)
}

func (d DaejeonHRCUrl) getDietInfoURL() string {
	return fmt.Sprintf("%s?nttSn=%s&mi=%s&bbsId=%s", d.infoURL, d.nttSn, d.mi, d.bbsId)
}

func (d DaejeonHRCUrl) getDownloadURL() string {
	return fmt.Sprintf("%s?fileKey=%s&nttSn=%s&bbsId=%s", d.downloadURL, d.fileKey, d.nttSn, d.bbsId)
}

func DownloadDietFile() {
	downloadLink := dURL.getDownloadURL()
	downRes, err := http.Get(downloadLink)
	if err != nil {
		log.Panic(err)
	}

	download, err := ioutil.ReadAll(downRes.Body)
	if err != nil {
		log.Panic(err)
	}
	file, err := os.Create("./" + "diet.xlsx")

	if err != nil {
		log.Panic(err)
	}
	file.Write(download)
}
