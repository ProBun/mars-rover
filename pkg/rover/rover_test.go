package rover

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ProBun/mars-rover/pkg/planet"
)

func TestRover_ProcessMoveCommands(t *testing.T) {
	mars := planet.NewMars(2, 2)

	tests := []struct {
		name            string
		rover           Rover
		command         string
		expectedX       int
		expectedY       int
		expectedHeading int
		err             error
	}{
		{
			"basic movement",
			Rover{
				XPos:    0,
				YPos:    0,
				heading: 0,
				planet:  &mars,
			},
			"FRFRF",
			0,
			1,
			2,
			nil,
		},
		{
			"getting out of range",
			Rover{
				XPos:    0,
				YPos:    0,
				heading: 0,
				planet:  &mars,
			},
			"FF",
			2,
			0,
			0,
			lostError,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.rover.ProcessMoveCommands(test.command)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expectedX, test.rover.XPos)
			assert.Equal(t, test.expectedY, test.rover.YPos)
			assert.Equal(t, test.expectedHeading, test.rover.heading)
		})

	}
}
