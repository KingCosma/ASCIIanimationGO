package main

//import the necessary packages
import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func animate(filename string) {
	fmt.Print("\033[H\033[2J")         // clear the screen before we start
	data, err := os.ReadFile(filename) //reads the "index.txt" file
	if err != nil {                    // check for any error
		log.Fatal(err)
	}
	frames := strings.Split(string(data), "/endf") //split the string into frames
	fmt.Println("")
	for i := range frames { //loop over the frames
		fmt.Println(frames[i])             //print the frame
		time.Sleep(500 * time.Millisecond) //wait for 0.5 second
		fmt.Print("\033[H\033[2J")         //clear the screen for the next frame
	}
	fmt.Println(frames[len(frames)-1]) //print the last frame to retain it

}

func main() {
	animate("index.txt")
}
