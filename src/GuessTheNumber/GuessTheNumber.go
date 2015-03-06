package main 

import (
	"fmt"
)

func main() {
	var (
		usernum int
		topbord int = 101 
		lowbord int = 1
		guess	int = (topbord + lowbord)/2
		)
	
	fmt.Println("Input number from 1 to 100")
	
	for  {
		_ , err := fmt.Scanln(&usernum)
		if err!=nil {
			fmt.Println(err,"wrong input try again!")
			var clearer string
			fmt.Scanln(&clearer)
			continue
			}
		if (usernum<1 || usernum>100) {
			fmt.Println("From 1 to 100 plz")
			continue
			} 
		break
		}
	
	for i := 0; i < 7; i++ {
		fmt.Printf("Do you guess number %d?\n",guess)
		fmt.Println("Answer \"yes\", \"bigger\" if your number bigger,\n or \"less\" if your number less")
		var answer string
		fmt.Scanln(&answer)
		if  answer == "bigger"{
				lowbord = guess
		} else if answer == "less" {
				topbord = guess
		} else if answer == "yes" {
				fmt.Printf("I find answer on %d try\n",i+1)
				break
		} else {
				fmt.Println("Wrong input, try again")
				i--;
		}
		if usernum == guess {
		fmt.Println("You guess number %d and try to trick me:(",usernum)
			break
		}
		guess = (topbord + lowbord)/2
	}
	
	
}