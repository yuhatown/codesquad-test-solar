package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	circleReader := bufio.NewReader(os.Stdin)

	fmt.Println("원의 크기는?")
	strCircle, err := circleReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	strCircle = strings.TrimSpace(strCircle)
	intCircle, err := strconv.Atoi(strCircle)
	if err != nil {
		log.Fatal(err)
	}

	intCircle = intCircle / 2

		for i := -intCircle; i <= intCircle; i++ {
			fmt.Println()
			for j := -intCircle; j <= intCircle; j++ {
				if(i*i + j*j == intCircle * intCircle) {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
}

