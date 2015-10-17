package tictactoe

import (
  . "github.com/smartystreets/goconvey/convey"
  "testing"
)

func TestField(t *testing.T){
  Convey("Field", t, func() {
    field := NewField()
    So(field.isGameFinished(),ShouldBeTrue)
    So(field.Turn(0,0,2),ShouldResemble,Field{ positions: [3][3]int{{2,0,0},{0,0,0},{0,0,0}}})
    So(field.Turn(0,0,2),ShouldNotResemble,Field{ positions: [3][3]int{{0,0,0},{0,0,0},{0,0,0}}})
  })
}
