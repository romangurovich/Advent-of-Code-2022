package main

import (
  common "Advent-of-Code-2022/Common"
  "fmt"
  "strconv"
)

func main() {
  data := common.GetData(8)

  rows := make([][]int, len(data))
  cols := make([][]int, len(data))
  trees := make(map[Point]bool, 0)

  for r, entry := range data {
    for c, char := range entry {
      height, _ := strconv.Atoi(string(char))
      if rows[r] != nil {
        rows[r] = append(rows[r], height)
      } else {
        rows[r] = []int{height}
      }
      if cols[c] != nil {
        cols[c] = append(cols[c], height)
      } else {
        cols[c] = []int{height}
      }
    }
  }

  for i, _ := range rows[0] {
    point := Point{X: i, Y: 0}
    trees[point] = true
  }

  for i, _ := range rows[len(rows)-1] {
    point := Point{X: i, Y: len(rows) - 1}
    trees[point] = true
  }

  for r, row := range rows {
    tallestLeft := row[0]
    tallestRight := row[len(row)-1]

    for i := 1; i < len(row)-1; i++ {
      point := Point{X: i, Y: r}

      if row[i] > tallestLeft {
        tallestLeft = row[i]
        trees[point] = true
      }
    }

    for i := len(row) - 2; i > 0; i-- {
      point := Point{X: i, Y: r}

      if row[i] > tallestRight {
        tallestRight = row[i]
        trees[point] = true
      }
    }
  }

  for i, _ := range cols[0] {
    point := Point{X: 0, Y: i}
    trees[point] = true
  }

  for i, _ := range cols[len(cols)-1] {
    point := Point{X: len(cols) - 1, Y: i}
    trees[point] = true
  }

  for c, col := range cols {
    tallestTop := col[0]
    tallestBottom := col[len(col)-1]

    for i := 1; i < len(col)-1; i++ {
      point := Point{X: c, Y: i}

      if col[i] > tallestTop {
        tallestTop = col[i]
        trees[point] = true
      }
    }

    for i := len(col) - 2; i > 0; i-- {
      point := Point{X: c, Y: i}

      if col[i] > tallestBottom {
        tallestBottom = col[i]
        trees[point] = true
      }
    }
  }

  fmt.Println(len(trees))
}

type Point struct {
  X int
  Y int
}
