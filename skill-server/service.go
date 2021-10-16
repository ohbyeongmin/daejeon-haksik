package skillserver

import (
	"time"

	"github.com/ohbyeongmin/daejeon-haksik/constants"
	"github.com/ohbyeongmin/daejeon-haksik/menu"
)

type SkillServerService interface {
	GetMenu(which constants.LunOrDin, weekday time.Weekday) []string
	Tomorrow(which constants.LunOrDin) []string
	AllWeeks(which constants.LunOrDin) [][]string
}

var HRCService SkillServerService = menu.HRCMenuService{}
