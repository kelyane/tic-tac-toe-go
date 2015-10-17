package tictactoe

type Field struct{
}

func NewField() Field {
  return Field{}
}

func (this *Field) isEmpty() bool{
  return false
}
