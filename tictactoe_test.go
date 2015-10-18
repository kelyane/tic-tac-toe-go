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

    player := X
    fieldWinner := NewField()
    fieldWinner = Field{ positions: [3][3]int{{1,1,1},{0,0,0},{0,0,0}}}
    So(fieldWinner.isWinner(player),ShouldBeTrue)
    So(fieldWinner.isTurnPossible(2,2),ShouldBeTrue)
    So(fieldWinner.isTurnPossible(0,0),ShouldBeFalse)
    fieldWinner = Field{ positions: [3][3]int{{0,0,0},{1,1,1},{0,0,0}}}
    So(fieldWinner.isWinner(player),ShouldBeTrue)
    fieldWinner = Field{ positions: [3][3]int{{0,0,0},{0,0,0},{1,1,1}}}
    So(fieldWinner.isWinner(player),ShouldBeTrue)
    fieldWinner = Field{ positions: [3][3]int{{0,0,2},{0,0,2},{0,0,2}}}
    So(fieldWinner.isWinner(player),ShouldBeFalse)
    fieldWinner = Field{ positions: [3][3]int{{1,0,0},{0,1,0},{0,0,1}}}
    So(fieldWinner.isWinner(player),ShouldBeTrue)
    fieldWinner = Field{ positions: [3][3]int{{0,0,1},{0,0,1},{0,0,2}}}
    So(fieldWinner.isWinner(player),ShouldBeFalse)
    fieldWinner = Field{ positions: [3][3]int{{0,1,1},{0,0,1},{1,0,2}}}
    So(fieldWinner.isWinner(player),ShouldBeFalse)
    player = 2;
    fieldWinner = Field{ positions: [3][3]int{{0,0,0},{0,0,0},{1,1,1}}}
    So(fieldWinner.isWinner(player),ShouldBeFalse)
    fieldWinner = Field{ positions: [3][3]int{{2,0,0},{2,0,0},{2,0,0}}}
    So(fieldWinner.isWinner(player),ShouldBeTrue)
    fieldWinner = Field{ positions: [3][3]int{{0,2,0},{0,2,0},{0,2,0}}}
    So(fieldWinner.isWinner(player),ShouldBeTrue)
    fieldWinner = Field{ positions: [3][3]int{{0,0,2},{0,0,2},{0,0,2}}}
    So(fieldWinner.isWinner(player),ShouldBeTrue)
    fieldWinner = Field{ positions: [3][3]int{{0,0,2},{0,2,0},{2,0,0}}}
    So(fieldWinner.isWinner(player),ShouldBeTrue)
    fieldWinner = Field{ positions: [3][3]int{{0,0,0},{0,0,0},{1,2,1}}}
    So(fieldWinner.isWinner(player),ShouldBeFalse)

    So(fieldWinner.printField(),ShouldEqual,"0 | 0 | 0 | \n0 | 0 | 0 | \n1 | 2 | 1 | \n")

    })
}
