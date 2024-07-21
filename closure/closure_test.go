package closure

import "testing"

// Functions for testing the activateCard Function.

func TestActivateCard1(t *testing.T) {
  expected := 98 // the final result
  activate := activateCard() // getting the debitFunc
  result := activate(2) // deducting the ammount and storing it into result

  if result != expected {
    t.Error("Expected 98 got", result) // result should be equal to 98
  }
}

func TestActivateCard2(t *testing.T) {
  expected := 2 // the final result
  activate := activateCard() // getting the debitFunc
  result := activate(98) // deducting the ammount and storing it into result

  if result != expected {
    t.Error("Expected 2 got", result) // result should be equal to 2
  }
}

