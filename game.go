package main

import (
	"bufio"
	"fmt"
	"github.com/uptonm/blackjack/model"
	"os"
	"strings"
)

// CheckForBlackjack is a helper method, consuming the player and dealers hand and returning the result
// of HasBlackjack
func CheckForBlackjack(playerHand, dealerHand *model.Hand) (bool, bool) {
	return playerHand.HasBlackjack(), dealerHand.HasBlackjack()
}

// ShowBlackjackResults is a helper method, consuming the player and dealer hands, printing the results
// if either have Blackjack
func ShowBlackjackResults(playerHand, dealerHand *model.Hand) {
	playerHas, dealerHas := playerHand.HasBlackjack(), dealerHand.HasBlackjack()

	if playerHas && dealerHas {
		fmt.Println("\nBoth players have blackjack! Draw!")
	} else if playerHas {
		fmt.Println("\nYou have blackjack! You win!")
	} else if dealerHas {
		fmt.Println("\nDealer has blackjack! Dealer wins!")
	}
}

// PromptForHit accepts a *bufio.Reader and repeatedly prompts the user to Hit or Stay until they input
// a valid choice
func PromptForHit(reader *bufio.Reader) string {
	fmt.Printf("\n\nPlease Choose [H]it/[S]tay: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.Trim(choice, "\n")

	for choice != "H" && choice != "h" && choice != "Hit" && choice != "hit" &&
		choice != "S" && choice != "s" && choice != "Stay" && choice != "stay" {
		fmt.Printf("\n\nPlease Choose [H]it/[S]tay: ")
		choice, _ = reader.ReadString('\n')
		choice = strings.Trim(choice, "\n")
	}

	return choice
}

// PromptForRematch accepts a *bufio.Reader and repeatedly prompts the user to Rematch or Exit until they input
// a valid choice
func PromptForRematch(reader *bufio.Reader) string {
	fmt.Printf("\n\nPlay Again [Y]es/[N]o: ")
	again, _ := reader.ReadString('\n')
	again = strings.Trim(again, "\n")

	for again != "Y" && again != "y" && again != "Yes" && again != "yes" &&
		again != "N" && again != "n" && again != "No" && again != "no" {
		fmt.Printf("\n\nPlay Again [Y]es/[N]o: ")
		again, _ := reader.ReadString('\n')
		again = strings.Trim(again, "\n")
	}

	return again
}

// DisplayGameState is a helper method that accepts both the playerHand and the dealerHands, as well
// as a boolean stating whether the game is over. If so it shows the dealer's full hand
func DisplayGameState(playerHand, dealerHand *model.Hand, gameOver bool) {
	fmt.Println("Dealer's Hand:")
	dealerHand.Display(gameOver)
	fmt.Printf("\n")
	fmt.Println("Your Hand:")
	playerHand.Display(gameOver)
}

// PlayBlackjack is the method containing the entire Blackjack game loop
func PlayBlackjack() {
	playing := true

	for playing {
		reader := bufio.NewReader(os.Stdin)
		deck := model.InitDeck()
		deck.Shuffle()

		playerHand := model.InitHand(false)

		dealerHand := model.InitHand(true)

		for i := 0; i < 2; i++ {
			playerCard, _ := deck.Deal()
			playerHand.AddCard(*playerCard)

			dealerCard, _ := deck.Deal()
			dealerHand.AddCard(*dealerCard)
		}

		DisplayGameState(playerHand, dealerHand, false)

		gameOver := false
		for !gameOver {
			playerBj, dealerBj := CheckForBlackjack(playerHand, dealerHand)

			if playerBj || dealerBj {
				gameOver = true
				fmt.Print("\033[H\033[2J") // Clear the console between hands
				fmt.Printf("Final Results: \n\n")
				DisplayGameState(playerHand, dealerHand, gameOver)
				ShowBlackjackResults(playerHand, dealerHand)
				continue
			}

			choice := PromptForHit(reader)

			if choice == "H" || choice == "h" || choice == "Hit" || choice == "hit" {
				playerCard, _ := deck.Deal()
				playerHand.AddCard(*playerCard)

				fmt.Print("\033[H\033[2J") // Clear the console between hands
				DisplayGameState(playerHand, dealerHand, gameOver)

				if playerHand.GetValue() > 21 {
					fmt.Println("\nBUST! You have lost. Better luck next time")
					gameOver = true
				}
			} else {
				gameOver = true
				playerVal := playerHand.GetValue()
				dealerVal := dealerHand.GetValue()

				fmt.Print("\033[H\033[2J") // Clear the console between hands
				fmt.Printf("Final Results: \n\n")
				DisplayGameState(playerHand, dealerHand, gameOver)

				if playerVal > dealerVal {
					fmt.Println("You Win!")
				} else if playerVal == dealerVal {
					fmt.Println("Draw!")
				} else {
					fmt.Println("Dealer Wins! Better luck next time.")
				}
			}
		}

		again := PromptForRematch(reader)

		if again == "N" || again == "n" || again == "No" || again == "no" {
			fmt.Println("Thanks for playing.")
			gameOver = true
			playing = false
		} else {
			fmt.Print("\033[H\033[2J") // Clear the console between games
			gameOver = true
		}
	}
}