package generic

type Numbers interface {
  int | int32 | float32 | float64
}

func sumNumbers[T Numbers](nums []T) T { /* being the generic datatype that takes all values defined int the interface */
  var result T = 0
  for i := range nums {
    result += nums[i]
  }
  return result
}
