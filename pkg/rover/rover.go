package rover

import (
	"fmt"

	"github.com/ProBun/mars-rover/pkg/planet"
)

var lostError = fmt.Errorf("LOST")

type Rover struct {
	XPos    int
	YPos    int
	heading int // * 90 for heading in degrees
	Lost    bool
	planet  planet.Planet
}

func NewRover(planet planet.Planet) Rover {
	r := Rover{heading: 0, planet: planet}
	r.CurrentPosition()
	return r
}

func (r *Rover) ProcessMoveCommands(moveCommands string) error {
	for _, command := range moveCommands {
		switch command {
		case 'L', 'R':
			if err := r.RotateDrone(command); err != nil {
				return err
			}
		case 'F':
			if err := r.MoveForward(); err != nil {
				return err
			}
		default:
			r.Lost = true
			return fmt.Errorf("unrecognized command, drone is lost")
		}
	}

	r.CurrentPosition()

	return nil
}

func (r *Rover) RotateDrone(direction rune) error {
	switch direction {
	case 'L':
		r.heading--
	case 'R':
		r.heading++
	}

	if r.heading > 3 {
		r.heading = 0
	}

	if r.heading < 0 {
		r.heading = 3
	}

	return nil
}

func (r *Rover) MoveForward() error {
	switch r.heading {
	case 0:
		r.XPos++
	case 1:
		r.YPos++
	case 2:
		r.XPos--
	case 3:
		r.XPos--
	}

	if r.XPos > r.planet.X()-1 || r.XPos < 0 {
		r.CurrentPosition()
		r.Lost = true
		return lostError
	}

	if r.YPos > r.planet.Y()-1 || r.YPos < 0 {
		r.CurrentPosition()
		r.Lost = true
		return lostError
	}

	return nil
}

func (r *Rover) CurrentPosition() {
	fmt.Print(fmt.Sprintf("(%d,%d,%s) ", r.XPos, r.YPos, processHeading(r.heading)))
}

func processHeading(heading int) string {
	var stringHeading string
	switch heading {
	case 0:
		stringHeading = "N"
	case 1:
		stringHeading = "E"
	case 2:
		stringHeading = "S"
	case 3:
		stringHeading = "W"
	}

	return stringHeading
}
