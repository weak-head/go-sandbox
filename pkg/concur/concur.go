package concur

import (
	"gobox/pkg/concur/routine"
)

func DoConcur() {
	routine.DoSay(5)
}
