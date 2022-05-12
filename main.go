package main

import(
	"fmt"
	"math/rand"
	"time"
	"os"
)

func clear() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	var token int
	var s string
	var tf int
	var item int
	var item2 int
	var item3 int
	var pitem string
	var pitem1 string
	var pitem2 string

	token = 100

	fmt.Println("Spin or Veiw tokens [s/t]: ")
	fmt.Scanln(&s)
	if s == "y" {
		clear()
		token = token - 10
		tf = 1
	} else if s == "n" {
		
	} else if s == "t"{
		fmt.Println(token)
	} else {
		clear()
		token = token - 1
		tf = 1
	}
	for tf == 1 {
		fmt.Println("Spinning...")
		time.Sleep(2)
		item = rand.Intn(3)
		item2 = rand.Intn(3)
		item3 = rand.Intn(3)
		if item == 0 {
			pitem = "7"
		} else if item == 1 {
			pitem = "<>"
		} else if item == 2 {
			pitem = "="
		} else if item == 3 {
			pitem = "&"
		}
		if item2 == 0 {
			pitem1 = "7"
		} else if item == 1 {
			pitem1 = "<>"
		} else if item == 2 {
			pitem1 = "="
		} else if item == 3 {
			pitem1 = "&"
		}
		if item3 == 0 {
			pitem2 = "7"
		} else if item == 1 {
			pitem2 = "<>"
		} else if item == 2 {
			pitem2 = "="
		} else if item == 3 {
			pitem2 = "&"
		}
		clear()
		fmt.Println("| " + pitem + " | " + pitem1 + " | " + pitem2 + " |")
		if item == 0 && item2 == 0 && item3 == 0 {
			fmt.Println("You earned 100 tokens!")
			token = token + 100
		} else if item == 1 && item2 == 1 && item3 == 1 {
			fmt.Println("You earned 75 tokens!")
			token = token + 75
		} else if item == 2 && item2 == 2 && item3 == 2 {
			fmt.Println("You earned 50 tokens!")
			token = token + 50
		} else if item == 3 && item2 == 3 && item3 == 3 {
			fmt.Println("You earned 25 tokens!")
			token = token + 25
		} else {
			fmt.Println("You dont earn anything")
		}
		var a string
		fmt.Println("Would you like to spin again [y/n]: ")
		fmt.Scanln(&a)
		if a == "y" {
			tf = 1
		}
		if a == "n" {
			tf = 0
		}
	if tf != 1 {
		os.Exit(1)
	}
	if token <= 0 {
		fmt.Println("Out of tokens")
		os.Exit(1)
	}
	} 
}