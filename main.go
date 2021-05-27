package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type CovidDataSet struct {
	Date       string `json:"date"`
	State      string `json:"state"`
	Confirmed  string `json:"confirmed"`
	Deceased   string `json:"deceased"`
	Recovered  string `json:"recovered"`
	State_Name string `json:"state_name"`
}

func main() {
	file, err := os.Open("./data/covid_dataset_state_level.csv")
	if err != nil {
		fmt.Println((err))
	}

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	// fmt.Println(records)
	var covid_data CovidDataSet
	var coviddatasets []CovidDataSet
	for _, rec := range records {
		covid_data.Date = rec[0]
		covid_data.State = rec[1]
		covid_data.Confirmed = rec[2]
		covid_data.Deceased = rec[3]
		covid_data.Recovered = rec[4]
		covid_data.State_Name = rec[5]
		coviddatasets = append(coviddatasets, covid_data)
	}
	json_data, err := json.Marshal(coviddatasets)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Println(string(json_data))

	// Creating JSON Output from .csv
	json_file, err := os.Create("output.json")
	if err != nil {
		fmt.Println(err)
	}
	defer json_file.Close()
	json_file.Write(json_data)
	json_file.Close()
}
