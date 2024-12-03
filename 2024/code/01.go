package main

import (
  "os"
  "strings"
  "strconv"
  "slices"
  "fmt"
  "math"
)

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func freqMap(arr []int) map[int]int {
  freq := make(map[int]int)
  for _ , num :=  range arr {
      freq[num] = freq[num]+1
  }
  return freq
}

func main() {
  dat, err := os.ReadFile("../input/01.txt")
  check(err)
  lines := strings.Split(string(dat), "\n" )

  left  :=  []int{}
  right := []int{}

  for _, line := range lines {
    if len(line) > 0 {
      fields := strings.Fields(line)
      num_1, err := strconv.Atoi(fields[0])
      check(err)

      left = append(left, num_1)

      num_2, err := strconv.Atoi(fields[1])
      check(err)
      right = append(right, num_2)
    }
  }

  slices.Sort(left)
  slices.Sort(right)

  var part1 float64 = 0
  for i := 0; i < len(left); i++ {
    part1 = math.Abs(float64(left[i] - right[i])) + part1
  }

  fmt.Println("Part 1: ", int(part1))

  freq := freqMap(right)
  var part2 float64 = 0
  for _, num := range left {
    part2 = float64(num) * float64(freq[num]) + part2
  }

  fmt.Println("Part 2: ", int(part2))
}
