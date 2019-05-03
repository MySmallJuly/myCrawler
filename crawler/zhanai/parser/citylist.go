package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	match := re.FindAllSubmatch(contents,-1)
	result := engine.ParserResult{}

	for _, i := range match {
		result.Items = append(result.Items,string(i[2]))
		result.Requsets = append(result.Requsets,engine.Requset{
			Url:string(i[1]),
			ParserFunc:engine.NewParser,
		})
	}
	return result
}
