package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/inancgumus/screen"
)

var allHeld bool = false
var run bool = true
var reinitialize bool = true
var diceSlice diceStates
var dice diceStates
var changeDiceValues bool = true
var numDiceVal int32
var numSidesVal int32

type dieState struct {
	id     int32
	sides  int32
	value  int32
	locked bool
	held   bool
}

type diceStates []*dieState

func (d *diceStates) RollUnlocked() {
	for _, die := range *d {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		if die.held == false && die.locked == false {
			die.value = r1.Int31n(die.sides) + 1
		}

	}
	d.ShowDice()
}

// This might be redundant, but leaving until the end just in case.
func (d *diceStates) RollAll() {
	fmt.Println("Rolling all dice...\n")

	for _, die := range *d {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		die.value = r1.Int31n(die.sides) + 1
	}
	fmt.Println("\nPost-roll values:\n{index, Number of sides, value, locked}\n")
	for _, die := range *d {
		fmt.Println(*die)
	}
}

func SetConditions(changeDiceValues bool) diceStates {
	// If the dice are not being reset, it is the first run through, and we need to create some initial variables.
	// OR the user has chosen to change some of these variables
	if changeDiceValues == true {

		numDicePointer := chooseNumDice()

		// Get number of sides
		screen.Clear()
		fmt.Println("How many sides do these dice have?")
		fmt.Scanln(&numSidesVal)
		/* numSidesPointer := &numSidesVal */

		// Make a container for all of the dice states (as pointers)
		diceSlice = make([]*dieState, *numDicePointer)

	}

	// Set each dice state's initial values, and assign to the slice as a pointer
	for i, _ := range diceSlice {
		var initialValues dieState = dieState{rune(i), numSidesVal, 0, false, false}
		diceSlice[i] = &initialValues
	}
	return diceSlice
}

func chooseNumDice() *int32 {
	// Get number of dice from user
	screen.Clear()
	fmt.Println("\n\n\nHow many dice would you like to roll? (Max 10)")
	fmt.Scanln(&numDiceVal)
	for numDiceVal >= 11 || numDiceVal <= 0 {
		screen.Clear()
		fmt.Println("\n\n\nHow many dice would you like to roll? (Max 10)")
		fmt.Scanln(&numDiceVal)
	}
	/* numDicePointer := &numDiceVal */
	return &numDiceVal
}

func (d *diceStates) ChooseLocked() {
	// Set up keyboard handling
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	// Set the initial screen
	d.ShowDice()

	// Take in key presses, and refresh screen until Esc is pressed
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		d.ShowDice()

		// Handle Key Presses

		if key == keyboard.KeyEsc {
			allHeld = true
			break
		}
		if key == keyboard.KeySpace {
			for _, die := range *d {
				die.held = die.locked
			}
			break
		}
		if char-48 <= rune(len((*d))) {
			switch char {
			case 49: // Key "1"
				if (*d)[0].held == false {
					if (*d)[0].locked == true {
						(*d)[0].locked = false
						d.ShowDice()
					} else {
						(*d)[0].locked = true
						d.ShowDice()
					}
				}
			case 50: // Key "2"
				if (*d)[1].held == false {
					if (*d)[1].locked == true {
						(*d)[1].locked = false
						d.ShowDice()
					} else {
						(*d)[1].locked = true
						d.ShowDice()
					}
				}
			case 51: // Key "3"
				if (*d)[2].held == false {
					if (*d)[2].locked == true {
						(*d)[2].locked = false
						d.ShowDice()
					} else {
						(*d)[2].locked = true
						d.ShowDice()
					}
				}
			case 52: // Key "4"
				if (*d)[3].held == false {
					if (*d)[3].locked == true {
						(*d)[3].locked = false
						d.ShowDice()
					} else {
						(*d)[3].locked = true
						d.ShowDice()
					}
				}
			case 53: // Key "5"
				if (*d)[4].held == false {
					if (*d)[4].locked == true {
						(*d)[4].locked = false
						d.ShowDice()
					} else {
						(*d)[4].locked = true
						d.ShowDice()
					}
				}
			case 54: // Key "6"
				if (*d)[5].held == false {
					if (*d)[5].locked == true {
						(*d)[5].locked = false
						d.ShowDice()
					} else {
						(*d)[5].locked = true
						d.ShowDice()
					}
				}
			case 55: // Key "7"
				if (*d)[6].held == false {
					if (*d)[6].locked == true {
						(*d)[6].locked = false
						d.ShowDice()
					} else {
						(*d)[6].locked = true
						d.ShowDice()
					}
				}
			case 56: // Key "8"
				if (*d)[7].held == false {
					if (*d)[7].locked == true {
						(*d)[7].locked = false
						d.ShowDice()
					} else {
						(*d)[7].locked = true
						d.ShowDice()
					}
				}
			case 57: // Key "9"
				if (*d)[8].held == false {
					if (*d)[8].locked == true {
						(*d)[8].locked = false
						d.ShowDice()
					} else {
						(*d)[8].locked = true
						d.ShowDice()
					}
				}
			case 48: // Key "0"
				if char-40 <= rune(len((*d))) {
					if (*d)[9].held == false {
						if (*d)[9].locked == true {
							(*d)[9].locked = false
							d.ShowDice()
						} else {
							(*d)[9].locked = true
							d.ShowDice()
						}
					}
				}
			}
		}
	}
}

func (d *diceStates) ShowDice() {
	screen.Clear()
	fmt.Println("\nPress the number of the die to toggle locked / unlocked state.\nPress Space to hold the dice that you have locked in.\nPress Esc to finish")
	fmt.Println("\n{Dice Number, Value, Lock State}  ")
	/* fmt.Printf("\n{Dice Number, Value, Lock State}  You pressed: rune %q, key %X\r\n", char, key) */
	for _, die := range *d {
		switch die.held {
		case true:
			fmt.Printf("Die %d: %d\tHeld\n", die.id+1, die.value)
		default:
			switch die.locked {
			case true:
				fmt.Printf("Die %d: %d\tKeep\n", die.id+1, die.value)
			case false:
				fmt.Printf("Die %d: %d\tRe-roll\n", die.id+1, die.value)
			}
		}
	}
}

func (d *diceStates) CheckAllHeld() bool {
	if (*d)[0].held == false {
		return false
	}
	for i := 1; i < len((*d)); i++ {
		// If the index you are checking is not the same as index zero, they can not all be the same.  If you get to the end, you haven't returned yet, so they all must be the same, and you should return true.
		if (*d)[i].held != (*d)[0].held {
			return false
		}
	}
	d.ShowDice()
	return true
}

func main() {
	for run == true {
		var menuChoice int

		if reinitialize == true {
			// Initalize the dice
			dice = SetConditions(true)
			allHeld = false
		} else {
			dice = SetConditions(false)
		}

		// While allHeld is false, continue to re-roll unlocked dice
		for allHeld == false {
			dice.RollUnlocked()

			// Ask user which they would like to keep / are they done rolling
			dice.ChooseLocked()

			/* Check to see if all dice are held.  If they are, the turn is over, and we should flag this somehow.  If not, main should re-roll unlocked dice. */
			allHeld = dice.CheckAllHeld()

		}
		// Ask if you would like to reset dice and take a new turn, or quit.

		fmt.Printf("\nAll dice are held.  Your turn is over.  Would you like to roll a new set of dice, or would you like to end the program?\nEnter:\n( 1 ) to play again with the same number of dice\n( 2 ) to choose a new number of dice\n(Any other entry) Quit.\n")
		fmt.Scanln(&menuChoice)
		switch menuChoice {
		case 1:
			reinitialize = false
			allHeld = false
			run = true
			break
			// Do nothing, run allHeld loop again
		case 2:
			reinitialize = true
			allHeld = false
			run = true
			break
		default:
			screen.Clear()
			fmt.Println("\nThanks for playing!")
			run = false
			break
		}
	}
}
