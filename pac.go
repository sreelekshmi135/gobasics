package main

import(
	"fmt"
	"math/rand"

)
func main(){
	
	rand.Seed(42)
	answers := []string{
		"It is certain",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
	
	}
	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])
}