package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Two of Clubs")
	for i := 0; i < len(cards); i++ {
		fmt.Println(cards[i])
	}
}

func newCard() string {
	return "Five of Spades"
}
