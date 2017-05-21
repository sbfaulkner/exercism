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

// Name returns the robot's random name (assigning it if necessary).
func (robot *Robot) Name() string {
	if robot.name == "" {
		robot.name = randomName()
	}
	return robot.name
}

// Reset clears the robot's settings.
func (robot *Robot) Reset() {
	robot.name = ""
}

var names = map[string]bool{}

func randomName() (name string) {
	rand.Seed(time.Now().UnixNano())

	for {
		name = fmt.Sprintf("%c%c%03d", 'A'+rune(rand.Intn(26)), 'A'+rune(rand.Intn(26)), rand.Intn(1000))
		if !names[name] {
			names[name] = true
			break
		}
	}

	return name
}
