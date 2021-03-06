package tictactoe

import (
    "strconv"
)

type Field struct{
  positions [3][3] int
}

type Player int

const (
        X Player = 1
        O Player = 2
)

func NewField() Field {
  return Field{ positions: [3][3]int{{0,0,0},{0,0,0},{0,0,0}}}
}

func (this *Field) IsGameFinished(x int, y int, player Player) bool{
  return (!this.isTurnPossible(x,y) && this.isWinner(player))
}

func (this *Field) Turn(x int, y int, player Player) Field {
  this.positions[x][y] = int(player)
  return *this
}

func (this *Field) isWinner(player Player) bool{
  return (
     (this.positions[0][0] == int(player) && this.positions[0][1] == int(player) && this.positions[0][2] == int(player)) || //First row horizontal
     (this.positions[1][0] == int(player) && this.positions[1][1] == int(player) && this.positions[1][2] == int(player)) || //Second row horizontal
     (this.positions[2][0] == int(player) && this.positions[2][1] == int(player) && this.positions[2][2] == int(player)) || //Third row horizontal
     (this.positions[0][0] == int(player) && this.positions[1][0] == int(player) && this.positions[2][0] == int(player)) || //First column vertical
     (this.positions[0][1] == int(player) && this.positions[1][1] == int(player) && this.positions[2][1] == int(player)) || //Second column vertical
     (this.positions[0][2] == int(player) && this.positions[1][2] == int(player) && this.positions[2][2] == int(player)) || //Third column vertical
     (this.positions[0][0] == int(player) && this.positions[1][1] == int(player) && this.positions[2][2] == int(player)) || //Main diagonal
     (this.positions[0][2] == int(player) && this.positions[1][1] == int(player) && this.positions[2][0] == int(player)) )  //Oposite diagonal
}

func (this *Field) isTurnPossible(x int, y int) bool {
  return this.positions[x][y] == 0
}

func (this *Field) FieldToString() string{
  fieldToStr := ""

  for i := 0; i <= 2; i++ {
    for j := 0; j <= 2; j++ {
      fieldToStr += strconv.Itoa(this.positions[i][j])
      fieldToStr += " | "
    }
    fieldToStr += "\n"
  }
  
  return fieldToStr
}
