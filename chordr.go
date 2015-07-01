package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

var EXIT = 0
var PLAY = 1
var INPUT = 2
var NEXT = 3

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	fmt.Println("0: EXIT; 1: PLAY; 2: INPUT; 3: NEXT")

	if choice := next_chord(); choice == NEXT {
		next_chord()
	} else {
		fmt.Println("Goodbye!")
	}
}

func next_chord() int {

	choice := NEXT
	chord := rand.Intn(4) //len(chords.chord) - 1)

	for choice != EXIT {
		fmt.Print("Enter your choice:")
		fmt.Scanf("%d", &choice)
		switch choice {
		case PLAY:
			play(chord)
		case INPUT:
			input(chord)
		default:
			return choice
		}
	}
	return choice

}

func play(chord int) {
	fmt.Println(Get(chord).name, ":", chord)

	mpg123Path, pathErr := exec.LookPath("mpg123")
	check(pathErr)

	args := []string{"mpg123", Get(chord).soundfile}

	env := os.Environ()
	go func() {
		sigs := make(chan os.Signal, 1)

		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			<-sigs
		}()

		mpg123Err := syscall.Exec(mpg123Path, args, env)
		check(mpg123Err)
	}()

}

func input(chord int) {
	var in_chord int
	_, err := fmt.Scanf("%d", &in_chord)
	check(err)

	if in_chord == chord {
		fmt.Println("Hit!")
	} else {
		fmt.Println("Miss!")
	}
}
