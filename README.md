# 대전대학교 HRC 학식 메뉴 알림 카카오봇 스킬 서버

-   해결해야 할 과제 :
-   엑셀 파일을 어떻게 Go 데이터로 저장하고 json 으로 표현 할지...

```
http.HandleFunc("/today", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		itemA := skillserver.ICItem{Title: "1", Description: "백미밥"}
		itemB := skillserver.ICItem{Title: "2", Description: "배추김치"}
		profile := skillserver.ICProfile{Title: "10월 15일 점심"}
		itemCard := skillserver.ItemCardType{
			Profile: profile,
			ItemList: []skillserver.ICItem{
				itemA,
				itemB,
			},
		}
		carousel := skillserver.CarouselType{
			Type: "itemCard",
			Items: []skillserver.ItemCardType{
				itemCard,
				itemCard,
			},
		}
		res := skillserver.SkillResponseType{
			Version: "2.0",
			Template: skillserver.TemplateType{
				Outputs: []skillserver.OutputsType{
					{
						Carousel: carousel,
					},
				},
			},
		}
		json.NewEncoder(rw).Encode(res)
	})

	fmt.Println("Listen on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
```

```
개선 전 코드

type Week int

const (
	MON Week = iota + 2
	TUE
	WED
	THU
	FRI
)

type MenuList struct {
	Menu map[Week][]string
}

type WeeksMenu struct {
	Lunch  *MenuList
	Dinner *MenuList
}

var mem *WeeksMenu

func initMem() {
	mem = &WeeksMenu{
		Lunch: &MenuList{
			Menu: make(map[Week][]string),
		},
		Dinner: &MenuList{
			Menu: make(map[Week][]string),
		},
	}
}

func main() {
	initMem()
	f, _ := excelize.OpenFile("diet.xlsx")
	sheetName := f.GetSheetList()[0]

	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
		return
	}

	lunchOrDinner := false
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
					lunchOrDinner = true
				}
				continue
			}
			if lunchOrDinner {
				mem.Dinner.Menu[Week(j)] = append(mem.Dinner.Menu[Week(j)], colCell)
			} else {
				mem.Lunch.Menu[Week(j)] = append(mem.Lunch.Menu[Week(j)], colCell)
			}
		}
	}
	fmt.Println(mem.Lunch.Menu)
}
```
