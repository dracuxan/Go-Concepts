package closure

func activateCard() func(int) int {
  amount := 100 // the amount after activating car is set to be 100

  debitFunc := func (debitAmount int) int {
    amount -= debitAmount
    /* the ammount var gets stored in var along with the function and gets 
    used even if it is not declared inside the function by the property 
    called closure */
    return amount
  }

  return debitFunc
}

