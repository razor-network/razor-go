package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/PaesslerAG/jsonpath"
	"github.com/avast/retry-go"
	"github.com/gocolly/colly"
	"io/ioutil"
	"net/http"
	"razor/core/types"
	"time"
)

//'https://staging-v2.skalenodes.com/v1/whispering-turais  -X POST -H "Content-Type: application/json"}'
//[]byte(`{"Type":1,"name":"test"}`
func (*UtilsStruct) GetDataFromAPI(dataSourceURLStruct types.DataSourceURL) ([]byte, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	var body []byte
	if dataSourceURLStruct.Type == "GET" {
		err := retry.Do(
			func() error {
				response, err := client.Get(dataSourceURLStruct.URL)
				if err != nil {
					return err
				}
				defer response.Body.Close()
				if response.StatusCode != 200 {
					log.Errorf("API: %s responded with status code %d", dataSourceURLStruct.URL, response.StatusCode)
					return errors.New("unable to reach API")
				}
				body, err = IOInterface.ReadAll(response.Body)
				if err != nil {
					return err
				}
				return nil
			}, retry.Attempts(2), retry.Delay(time.Second*2))
		if err != nil {
			return nil, err
		}
	}
	if dataSourceURLStruct.Type == "POST" {
		postBody, err := json.Marshal(dataSourceURLStruct.Body)
		if err != nil {
			log.Errorf("Error in marshalling body of a POST request URL %s: %v", dataSourceURLStruct.URL, err)
		}
		responseBody := bytes.NewBuffer(postBody)
		err = retry.Do(
			func() error {
				response, err := client.Post(dataSourceURLStruct.URL, dataSourceURLStruct.ContentType, responseBody)
				if err != nil {
					log.Errorf("Error sending POST request URL %s: %v", dataSourceURLStruct.URL, err)
					return err
				}
				defer response.Body.Close()
				if response.StatusCode != 200 {
					log.Errorf("URL: %s responded with status code %d", dataSourceURLStruct.URL, response.StatusCode)
					return errors.New("unable to reach API")
				}
				body, err = ioutil.ReadAll(response.Body)
				if err != nil {
					return err
				}
				return nil
			}, retry.Attempts(2), retry.Delay(time.Second*2))
		if err != nil {
			return nil, err
		}
	}
	return body, nil
}

func (*UtilsStruct) GetDataFromJSON(jsonObject map[string]interface{}, selector string) (interface{}, error) {
	if selector[0] == '[' {
		selector = "$" + selector
	} else {
		selector = "$." + selector
	}
	return jsonpath.Get(selector, jsonObject)
}

func (*UtilsStruct) GetDataFromXHTML(dataSourceURLStruct types.DataSourceURL, selector string) (string, error) {
	c := colly.NewCollector()
	var priceData string
	c.OnXML(selector, func(e *colly.XMLElement) {
		priceData = e.Text
	})
	err := c.Visit(dataSourceURLStruct.URL)
	if err != nil {
		return "", err
	}
	return priceData, nil
}
