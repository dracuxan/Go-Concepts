package generic

import "testing"

func TestSumInt(t *testing.T) {
  nums1 := []int{1, 2, 3, 4, 5}
  expected := 15
  result := sumNumbers(nums1)

  if expected != result {
    t.Errorf("Expected %v got %v.", expected, result)
  }
}


func TestSumInt32(t *testing.T) {
  nums1 := []int32{1, 2, 3, 4, 5}
  var expected int32 = 15
  result := sumNumbers(nums1)

  if expected != result {
    t.Errorf("Expected %v got %v.", expected, result)
  }
}

func TestSumFloat32(t *testing.T) {
  nums1 := []float32{1.1, 2.1, 3.1, 4.1, 5.1}
  var expected float32 = 15.5
  result := sumNumbers(nums1)

  if expected != result {
    t.Errorf("Expected %v got %v.", expected, result)
  }
}

func TestSumFloat64(t *testing.T) {
  nums1 := []float64{1.1, 2.1, 3.1, 4.1, 5.1}
  expected := 15.5
  result := sumNumbers(nums1)

  if expected != result {
    t.Errorf("Expected %v got %v.", expected, result)
  }
}
