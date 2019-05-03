package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n",contents)
	result := ParserCityList(contents)
	const resultsize = 470
	if len(result.Requsets) != resultsize {
		t.Errorf("result should  hava %d requests, but had %d", resultsize, len(result.Requsets))

		if len(result.Items) != resultsize {
			t.Errorf("result should  hava %d Items, but had %d", resultsize, len(result.Items))
		}

	}
}