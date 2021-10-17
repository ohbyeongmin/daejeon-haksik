package cli

import (
	"flag"
	"os"

	"github.com/ohbyeongmin/daejeon-haksik/crawling"
	"github.com/ohbyeongmin/daejeon-haksik/menu"
	skillserver "github.com/ohbyeongmin/daejeon-haksik/skill-server"
)

func Start() {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	var server bool
	var download bool

	fs.BoolVar(&server, "s", false, "server start")
	fs.BoolVar(&download, "d", false, "file download")

	fs.Parse(os.Args[1:])

	if len(os.Args) < 2 {
		fs.Usage()
		os.Exit(1)
	}

	if server {
		menu.InitMenu()
		skillserver.ServerStart()
	} else if download {
		crawling.DownloadDietFile()
	}
}
