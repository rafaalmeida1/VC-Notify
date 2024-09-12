package scrapers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeConcursosByLocalidade(localidade string) ([]map[string]string, error) {
	url := "https://www.pciconcursos.com.br/concursos"
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

	var concursos []map[string]string
	doc.Find("#concursos .na").Each(func(index int, item *goquery.Selection) {
		localidadeText := item.Find(".cc").Text()
		if strings.TrimSpace(strings.ToUpper(localidadeText)) == strings.ToUpper(localidade) {
			url, _ := item.Find(".ca a").Attr("href")
			descricao := item.Find(".cd").Text()
			if url != "" {
				concursos = append(concursos, map[string]string{
					"url":       url,
					"descricao": descricao,
				})
			}
		}
	})

	if len(concursos) == 0 {
		return nil, errors.New("não foram encontrados concursos para a localidade " + localidade)
	}

	return concursos, nil
}

func ScrapeConcursosByCategoriaAndLocalidade(localidade, area string) ([]map[string]string, error) {
	url := "https://www.pciconcursos.com.br/vagas/" + area
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

	var concursos []map[string]string
	doc.Find("#concursos .na").Each(func(index int, item *goquery.Selection) {
		localidadeText := item.Find(".cc").Text()
		if strings.TrimSpace(strings.ToUpper(localidadeText)) == strings.ToUpper(localidade) {
			url, _ := item.Find(".ca a").Attr("href")
			descricao := item.Find(".cd").Text()
			if url != "" {
				concursos = append(concursos, map[string]string{
					"url":       url,
					"descricao": descricao,
				})
			}
		}
	})

	if len(concursos) == 0 {
		return nil, errors.New("não foram encontrados concursos para a localidade " + localidade + " e área " + area)
	}

	return concursos, nil
}

func ScrapeCargosInConcursos() ([]string, error) {
	url := "https://www.pciconcursos.com.br/cargos"
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

	// tenho ul com classe link-d, dentro pegar cada a dentro de cada LI, e só trazer o texto dentro da tag a

	var cargos []string
	doc.Find(".link-d li a").Each(func(index int, item *goquery.Selection) {
		cargo := item.Text()
		cargos = append(cargos, cargo)
	})

	return cargos, nil
}
