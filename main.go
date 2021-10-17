package main

import (
	"github.com/ohbyeongmin/daejeon-haksik/cli"
	_ "github.com/ohbyeongmin/daejeon-haksik/crawling"
)

func main() {
	cli.Start()
}

// 월요일 아침 10시마다 데이터 다운로드 및 초기화
