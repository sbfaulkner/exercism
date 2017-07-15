package robot

import "fmt"

// Action is used to communicate the robot's intention to the room
type Action Command

// StartRobot pipes actions to the room and closes the channel when done
func StartRobot(cmd chan Command, act chan Action) {
	for {
		select {
		case c := <-cmd:
			act <- Action(c)
		}
	}
}

// Room simulates the motion of a robot within a room
func Room(extent Rect, robot Step2Robot, act chan Action, rep chan Step2Robot) {
	for {
		select {
		case a := <-act:
			switch a {
			case 0:
				rep <- robot
			case 'A':
				robot.advance(func(p Pos) bool {
					return p.Easting >= extent.Min.Easting &&
						p.Northing >= extent.Min.Northing &&
						p.Easting <= extent.Max.Easting &&
						p.Northing <= extent.Max.Northing
				})
			case 'L':
				robot.left()
			case 'R':
				robot.right()
			}
		}
	}
}

func (robot Step2Robot) String() string {
	return fmt.Sprintf("{%s %v}", robot.Dir, robot.Pos)
}

// advance the robot within the bounds provided
func (robot *Step2Robot) advance(valid func(p Pos) bool) {
	dest := robot.Pos

	switch robot.Dir {
	case N:
		dest.Northing++
	case S:
		dest.Northing--
	case E:
		dest.Easting++
	case W:
		dest.Easting--
	}

	if valid(dest) {
		robot.Pos = dest
	}
}

// left turns the robot left
func (robot *Step2Robot) left() {
	robot.Dir = (robot.Dir + 3) % 4
}

// right turns the robot right
func (robot *Step2Robot) right() {
	robot.Dir = (robot.Dir + 1) % 4
}
