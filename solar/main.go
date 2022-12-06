package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type SolarInfo struct {
	x int
	y int
	radius int
}

// func solarSize() {
// 	var a [][]int
// 	for i := 30; i < 30; i++ {
// 		for j := 30; j < 30; j++ {
// 			fmt.Print(a[i][j])
// 		}
// 	}
// }

func planetSize(i, j, radius int) {	
	for i := -radius; i <= radius; i++ {
		fmt.Println()
		for j := -radius; j <= radius; j++ {
			if(i*i + j*j <= radius * radius) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
	}
}

func main() {
	
	var xx, yy int

	circleReader := bufio.NewReader(os.Stdin)
	fmt.Println("Sun, Earth, Moon")
	fmt.Println("날짜를 입력하세요. ex. 1월 1일")
	strCircle, err := circleReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	xx, err = strconv.Atoi(strCircle[:1])
	if err != nil {
		log.Fatal(err)
	}
	yy, err = strconv.Atoi(strCircle[5:6])
	if err != nil {
		log.Fatal(err)
	}
	

	var moon SolarInfo = SolarInfo{xx, yy, 0}
	var earth SolarInfo = SolarInfo{xx+3, yy, 1}
	var Sun SolarInfo = SolarInfo{xx+7, yy, 2}

	planetSize(moon.x, moon.y, moon.radius)
	planetSize(earth.x, earth.y, earth.radius)
	planetSize(Sun.x, Sun.y, Sun.radius)
	
}

// 1월 1일을 입력했을 때, y좌표는 1이고 x좌표만이 계속 변동되면 될 것 같다. 