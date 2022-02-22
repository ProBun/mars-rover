package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ProBun/mars-rover/pkg/planet"
	"github.com/ProBun/mars-rover/pkg/rover"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please specify size of the grid, example(5 8): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("There was a problem reading input")
		return
	}

	input = strings.TrimSuffix(input, "\n")

	gridSize := strings.Split(input, " ")

	x, err := strconv.Atoi(gridSize[0])
	if err != nil {
		fmt.Println("X value is not an int")
		return
	}

	y, err := strconv.Atoi(gridSize[1])
	if err != nil {
		fmt.Println("Y value is not an int, error", err)
		return
	}

	mars := planet.NewMars(x, y)

	for {
		fmt.Println("Deploying new drone")
		fmt.Println("Please provide move commands (L - left, R - right, F - forward): ")

		marsRover := rover.NewRover(&mars)

		for {
			moveInput, err := reader.ReadString('\n')
			if err != nil {
				marsRover.Lost = true
				fmt.Println("There was an error reading move commands, drone is lost")
				return
			}

			moveInput = strings.TrimSuffix(moveInput, "\n")

			err = marsRover.ProcessMoveCommands(moveInput)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
		}
	}
}
