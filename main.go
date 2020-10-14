package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/inancgumus/screen"
)

type dieState struct {
	id     int32
	sides  int32
	value  int32
	locked bool
}

type diceStates []*dieState

func (d *diceStates) RollUnlocked() {
	// while allLocked == false do this.  Check at end of loop to see if all are locked.
	allLocked := false
	for allLocked == false {
		fmt.Println("Rolling unlocked dice...\n")

		for _, die := range *d {
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			if die.locked == false {
				die.value = r1.Int31n(die.sides) + 1
			}
		}
		d.ShowDice()

	}
}

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

func SetInitConditions() diceStates {
	fmt.Println("Setting initial die states...")

	// Get number of dice from user
	fmt.Println("How many dice would you like to roll? >> 7")
	var numDiceVal int32 = 7 // numDicePointer := fmt.Scanln(&numDiceVal)
	numDicePointer := &numDiceVal

	// Assume 6 sided die for now
	var numSidesVal int32 = 6
	// numSidesPointer := &numSidesVal

	// Make a container for all of the dice states (as pointers)
	var diceSlice diceStates = make([]*dieState, *numDicePointer)

	// Set each dice state's initial values, and assign to the slice as a pointer
	for i, _ := range diceSlice {
		var initialValues dieState = dieState{rune(i), numSidesVal, 0, false}
		diceSlice[i] = &initialValues
	}
	return diceSlice
}

func (d *diceStates) Lock() {

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
			break
		}
		switch char {
		case 49: // Key "1"
			if (*d)[0].locked == true {
				(*d)[0].locked = false
				d.ShowDice()
			} else {
				(*d)[0].locked = true
				d.ShowDice()
			}
		case 50: // Key "2"
			if (*d)[1].locked == true {
				(*d)[1].locked = false
				d.ShowDice()
			} else {
				(*d)[1].locked = true
				d.ShowDice()
			}
		case 51: // Key "3"
			if (*d)[2].locked == true {
				(*d)[2].locked = false
				d.ShowDice()
			} else {
				(*d)[2].locked = true
				d.ShowDice()
			}
		case 52: // Key "4"
			if (*d)[3].locked == true {
				(*d)[3].locked = false
				d.ShowDice()
			} else {
				(*d)[3].locked = true
				d.ShowDice()
			}
		case 53: // Key "5"
			if (*d)[4].locked == true {
				(*d)[4].locked = false
				d.ShowDice()
			} else {
				(*d)[4].locked = true
				d.ShowDice()
			}
		case 54: // Key "6"
			if (*d)[5].locked == true {
				(*d)[5].locked = false
				d.ShowDice()
			} else {
				(*d)[5].locked = true
				d.ShowDice()
			}
		case 55: // Key "7"
			if (*d)[6].locked == true {
				(*d)[6].locked = false
				d.ShowDice()
			} else {
				(*d)[6].locked = true
				d.ShowDice()
			}
		case 56: // Key "8"
			if (*d)[7].locked == true {
				(*d)[7].locked = false
				d.ShowDice()
			} else {
				(*d)[7].locked = true
				d.ShowDice()
			}
		case 57: // Key "9"
			if (*d)[8].locked == true {
				(*d)[8].locked = false
				d.ShowDice()
			} else {
				(*d)[8].locked = true
				d.ShowDice()
			}
		case 48: // Key "0"
			if (*d)[9].locked == true {
				(*d)[9].locked = false
				d.ShowDice()
			} else {
				(*d)[9].locked = true
				d.ShowDice()
			}
		default:
			// Do nothing
		}
	}

}

func (d *diceStates) ShowDice() {
	screen.Clear()
	fmt.Println("\nPress the number of the die to toggle locked / unlocked state.  Press Esc to lock in.")
	fmt.Println("\n{Dice Number, Value, Lock State}  ")
	/* fmt.Printf("\n{Dice Number, Value, Lock State}  You pressed: rune %q, key %X\r\n", char, key) */
	for _, die := range *d {
		if die.locked == false {
			fmt.Printf("Die %d: %d\tUnlocked\n", die.id, die.value)
		} else {
			fmt.Printf("Die %d: %d\tLocked\n", die.id, die.value)
		}
	}
}

func (d *diceStates) CheckAllLocked() bool {
	for i := 1; i < len((*d)); i++ {
		// If the index you are checking is not the same as index zero, they can not all be the same.  If you get to the end, you haven't returned yet, so they all must be the same, and you should return true.
		if (*d)[i] != (*d)[0] {
			return false
		}
	}
	return true

}

func main() {

	// Initalize the dice
	dice := SetInitConditions()

	// Roll dice
	dice.RollUnlocked()

	// Ask user which they would like to keep / are they done rolling
	dice.Lock()

	// Repeat until no dice or done rolling, then reset dice.
	// Check to see if all dice are locked in.  If they are, the turn is over, and we should flag this somehow.  If not, main should re-roll unlocked dice.
	d.CheckAllLocked() // Returns true if all dice are the same.  This line should only evaluate if index 0 is locked, because the default is unlocked (false)
	// Add in error checking so that you can not unlock dice that were previously locked

}
