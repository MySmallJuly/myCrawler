package engine

import (
	"crawler/fetcher"
	"fmt"
	"log"
)

func Run(seeds ...Requset)  {
	var requests []Requset
	for _, r := range seeds{
		requests = append(requests,r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s",r.Url)

		body, err := fetcher.Fetch(r.Url)
		if err != nil{
			log.Print("Fetcher: error    fetching url %s : %v " ,r.Url, err)
			continue
		}
		parserResult := r.ParserFunc(body)
		requests = append(requests,parserResult.Requsets...)
		//fmt.Println(requests)
		for _, item := range parserResult.Items {
			fmt.Printf("Got item %s\n",item)

		}
	}

}
