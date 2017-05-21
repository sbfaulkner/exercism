package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

const testVersion = 1

// Robot type represents a robot's factory settings
type Robot struct {
	name string
}

func (robot *Robot) Name() string {
	if robot.name == "" {
		rand.Seed(time.Now().UnixNano())
		robot.name = fmt.Sprintf("%c%c%03d", 'A'+rune(rand.Intn(26)), 'A'+rune(rand.Intn(26)), rand.Intn(1000))
	}
	return robot.name
}

func (robot *Robot) Reset() {
	robot.name = ""
}
