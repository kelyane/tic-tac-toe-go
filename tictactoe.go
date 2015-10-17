package tictactoe

type Field struct{
  positions [3][3] int
}

func NewField() Field {
  return Field{ positions: [3][3]int{{0,0,0},{0,0,0},{0,0,0}}}
}

func (this *Field) isGameFinished() bool{
  return true
}

func (this *Field) Turn(x int, y int, z int) Field {
  this.positions[x][y] = z
  return *this
}
