package menu

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type Week int
type LunOrDin int

const (
	MON Week = iota + 2
	TUE
	WED
	THU
	FRI
)

const (
	LUNCH LunOrDin = iota
	DINNER
)

type Menu map[Week][]string

type MenuTable struct {
	Table map[LunOrDin]Menu
}

var mem *MenuTable

func init() {
	mem = &MenuTable{
		Table: make(map[LunOrDin]Menu),
	}
	mem.Table[LUNCH] = make(Menu)
	mem.Table[DINNER] = make(Menu)
}

func Start() {
	f, _ := excelize.OpenFile("diet.xlsx")
	sheetName := f.GetSheetList()[0]

	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
		return
	}

	lunchOrDinner := LUNCH
	for i, row := range rows {
		if i < 4 {
			continue
		}
		if i > 23 {
			break
		}
		for j, colCell := range row {
			if j < 2 || j >= 7 {
				if colCell == "석  식" {
					lunchOrDinner = DINNER
				}
				continue
			}
			mem.Table[lunchOrDinner][Week(j)] = append(mem.Table[lunchOrDinner][Week(j)], colCell)
		}
	}
	fmt.Println(mem.Table[DINNER][TUE])
}
