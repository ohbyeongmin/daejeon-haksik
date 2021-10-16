package skillserver

import (
	"github.com/ohbyeongmin/daejeon-haksik/constants"
	"github.com/ohbyeongmin/daejeon-haksik/menu"
)

type SkillServerService interface {
	Today(which constants.LunOrDin) []string
	Tomorrow(which constants.LunOrDin) []string
	AllWeeks(which constants.LunOrDin) [][]string
}

var HRCService SkillServerService = menu.HRCMenuService{}
