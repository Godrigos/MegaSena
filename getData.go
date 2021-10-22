package main

import (
	"log"
	"net/http"
)

func getData() *http.Response {
	link := "http://loterias.caixa.gov.br/wps/portal/loterias/landing/megasena/!ut/p/a1/04_Sj9CPykssy0xPLMnMz0vMAfGjzOLNDH0MPAzcDbwMPI0sDBxNXAOMwrzCjA0sjIEKIoEKnN0dPUzMfQwMDEwsjAw8XZw8XMwtfQ0MPM2I02-AAzgaENIfrh-FqsQ9wNnUwNHfxcnSwBgIDUyhCvA5EawAjxsKckMjDDI9FQE-F4ca/dl5/d5/L2dBISEvZ0FBIS9nQSEh/pw/Z7_HGK818G0K8DBC0QPVN93KQ10G1/res/id=historicoHTML/c=cacheLevelPage/=/"
	c := http.Client{}
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		panic(err)
	}
	req.AddCookie(&http.Cookie{Name: "security", Value: "true", Path: "/"})

	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	} else if resp.StatusCode != 200 {
		log.Fatal(resp.Status)
	}
	return resp
}
