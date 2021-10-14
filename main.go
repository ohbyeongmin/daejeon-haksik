package main

import (
	"encoding/json"
	"net/http"

	"github.com/xuri/excelize/v2"
)

type Menu struct {
	Date string   `json:"date"`
	Day  string   `json:"day"`
	List []string `json:"list"`
}

func TestFunc() Menu {
	f, _ := excelize.OpenFile("diet.xlsx")
	sheetName := f.GetSheetList()[0]
	var m = Menu{}

	a := []string{"3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "14"}
	for i, v := range a {
		cell, _ := f.GetCellValue(sheetName, "D"+v)
		if i == 0 {
			m.Date = cell
		} else if i == 1 {
			m.Day = cell
		} else {
			m.List = append(m.List, cell)
		}
	}
	return m
}

func main() {
	m := TestFunc()

	http.HandleFunc("/today", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(m)
	})

	http.ListenAndServe(":3000", nil)
}
