package performancetest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func PerformanceScript() {

	jsonFile, err := os.Open("src/test.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened json file")
	defer jsonFile.Close()
	var jsonData JSON
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &jsonData)

	var array = jsonData.Results[0].Times

	var suma = 0
	var max, min = array[0], array[0]
	var length = len(array)
	for i := 0; i < length; i++ {
		suma += array[i]
		if max < array[i] {
			max = array[i]
		}
		if min > array[i] {
			min = array[i]
		}
	}
	var medium = suma / length

	fmt.Println(fmt.Sprintf("%s%d", "Min: ", min))
	fmt.Println(fmt.Sprintf("%s%d", "Med: ", medium))
	fmt.Println(fmt.Sprintf("%s%d", "Max: ", max))
}

type JSON struct {
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Timestamp     time.Time `json:"timestamp"`
	CollectionId  string    `json:"collection_id"`
	FolderId      int       `json:"folder_id"`
	EnvironmentId string    `json:"environment_id"`
	TotalPass     int       `json:"totalPass"`
	TotalFail     int       `json:"totalFail"`
	Results       []struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		Time         int    `json:"time"`
		ResponseCode struct {
			Code int    `json:"code"`
			Name string `json:"name"`
		} `json:"responseCode"`
		Tests struct {
		} `json:"tests"`
		TestPassFailCounts struct {
		} `json:"testPassFailCounts"`
		Times    []int `json:"times"`
		AllTests []struct {
		} `json:"allTests"`
	} `json:"results"`
	Count      int `json:"count"`
	TotalTime  int `json:"totalTime"`
	Collection struct {
		Requests []struct {
			Id     string `json:"id"`
			Method string `json:"method"`
		} `json:"requests"`
	} `json:"collection"`
}
