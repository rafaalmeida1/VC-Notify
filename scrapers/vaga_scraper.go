package scrapers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeVagas() ([]map[string]string, error) {
	url := "https://www.empregos.com.br/vagas/desenvolvedor"
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("status code error: " + res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var vagas []map[string]string
	doc.Find(".vaga").Each(func(index int, item *goquery.Selection) {
		titulo := item.Find(".cargo").Text()
		empresa := item.Find(".empresa").Text()
		localizacao := item.Find(".localidade").Text()
		link, _ := item.Find("a").Attr("href")

		vagas = append(vagas, map[string]string{
			"titulo":      strings.TrimSpace(titulo),
			"empresa":     strings.TrimSpace(empresa),
			"localizacao": strings.TrimSpace(localizacao),
			"link":        "https://www.empregos.com.br" + link,
		})
	})

	return vagas, nil
}
