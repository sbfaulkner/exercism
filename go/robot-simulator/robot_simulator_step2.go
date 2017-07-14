package robot

// Action is used to communicate the robot's intention to the room
type Action Command

// StartRobot pipes actions to the room and closes the channel when done
func StartRobot(cmd chan Command, act chan Action) {
	for {
		if c, ok := <-cmd; ok {
			act <- Action(c)
		} else {
			break
		}
	}

	close(act)
}

// Room simulates the motion of a robot within a room
func Room(extent Rect, robot Step2Robot, act chan Action, rep chan Step2Robot) {
	for {
		if a, ok := <-act; ok {
			switch a {
			case 'A':
				robot.advance(extent)
			case 'L':
				robot.left()
			case 'R':
				robot.right()
			}
		} else {
			rep <- robot
			break
		}
	}
}

// advance the robot within the bounds provided
func (robot *Step2Robot) advance(extent Rect) {
	switch robot.Dir {
	case N:
		if robot.Northing < extent.Max.Northing {
			robot.Northing++
		}
	case S:
		if robot.Northing > extent.Min.Northing {
			robot.Northing--
		}
	case E:
		if robot.Easting < extent.Max.Easting {
			robot.Easting++
		}
	case W:
		if robot.Easting > extent.Min.Easting {
			robot.Easting--
		}
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
