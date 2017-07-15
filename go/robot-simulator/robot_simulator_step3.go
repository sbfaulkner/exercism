package robot

import (
	"fmt"
	"strings"
)

type Action3 struct {
	name string
	cmd  Command
}

func (command Command) String() (s string) {
	switch command {
	case 'A':
		s = "Advance"
	case 'L':
		s = "Left"
	case 'R':
		s = "Right"
	}

	return
}

func (action Action3) String() string {
	return fmt.Sprintf("{%v %v}", action.name, action.cmd)
}

// StartRobot3 pipes actions to the room and closes the channel when done
func StartRobot3(name, script string, action chan Action3, log chan string) {
	defer func() {
		action <- Action3{name, 0}
	}()

	if name == "" {
		log <- "robot: without a name"
		return
	}

	for _, c := range script {
		if !strings.ContainsRune("ALR", c) {
			log <- "robot: undefined command in script"
			return
		}

		action <- Action3{name, Command(c)}
	}
}

func (robot Step3Robot) String() string {
	return fmt.Sprintf("{%s %s}", robot.Name, robot.Step2Robot)
}

// Room3 simulates the motion of a robot within a room
func Room3(extent Rect, robots []Step3Robot, action chan Action3, report chan []Step3Robot, log chan string) {
	defer func() {
		report <- robots
	}()

	room := map[RU]map[RU]*Step3Robot{}
	for e := extent.Min.Easting; e <= extent.Max.Easting; e++ {
		room[e] = map[RU]*Step3Robot{}
	}

	robotsByName := map[string]*Step3Robot{}

	for i, r := range robots {
		if robotsByName[r.Name] != nil {
			log <- "robot: duplicate name"
		} else {
			robotsByName[r.Name] = &robots[i]
		}

		if room[r.Pos.Easting][r.Pos.Northing] != nil {
			log <- "robot: placed at the same place"
		} else if r.Pos.Easting < extent.Min.Easting || r.Pos.Easting > extent.Max.Easting || r.Pos.Northing < extent.Min.Easting || r.Pos.Northing > extent.Max.Northing {
			log <- "robot: placed outside of the room"
		} else {
			room[r.Pos.Easting][r.Pos.Northing] = &robots[i]
		}
	}

	running := len(robots)

	for running > 0 {
		select {
		case a := <-action:
			robot := robotsByName[a.name]
			if robot == nil {
				log <- "robot: action from an unknown robot"
				return
			}

			switch a.cmd {
			case 0:
				running--
			case 'A':
				delete(room[robot.Easting], robot.Northing)

				robot.advance(func(p Pos) bool {
					if p.Easting < extent.Min.Easting ||
						p.Northing < extent.Min.Northing ||
						p.Easting > extent.Max.Easting ||
						p.Northing > extent.Max.Northing {
						log <- "robot: attempting to advance into a wall"
						return false
					}

					if room[p.Easting][p.Northing] != nil {
						log <- "robot: attempting to advance into another robot"
						return false
					}

					return true
				})

				room[robot.Easting][robot.Northing] = robot
			case 'L':
				robot.left()
			case 'R':
				robot.right()
			}
		}
	}
}
