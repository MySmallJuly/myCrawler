package main

import (
	"crawler/engine"
	"crawler/zhanai/parser"
)

func main() {
	engine.Run(engine.Requset{
		Url: "https://www.zhenai.com/zhenghun/",
		ParserFunc:parser.ParserCityList,
	})

}



