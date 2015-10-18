package tictactoe

import (
  . "github.com/smartystreets/goconvey/convey"
  "testing"
)

func TestField(t *testing.T){
  Convey("Field", t, func() {
    field := NewField()
    player := X
    So(field.isGameFinished(0,0,player),ShouldBeFalse)

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

    So(fieldWinner.fieldToString(),ShouldEqual,"0 | 1 | 1 | \n0 | 0 | 1 | \n1 | 0 | 2 | \n")

    player = O
    So(field.Turn(0,0,player),ShouldResemble,Field{ positions: [3][3]int{{2,0,0},{0,0,0},{0,0,0}}})
    So(field.Turn(0,0,player),ShouldNotResemble,Field{ positions: [3][3]int{{0,0,0},{0,0,0},{0,0,0}}})
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

    So(fieldWinner.fieldToString(),ShouldEqual,"0 | 0 | 0 | \n0 | 0 | 0 | \n1 | 2 | 1 | \n")

    })
}

func TestGame(t *testing.T){
  Convey("X is winner", t, func() {
    field := NewField()
    player := X
    So(field.Turn(0,0,player),ShouldResemble,Field{ positions: [3][3]int{{1,0,0},{0,0,0},{0,0,0}}})
    So(field.isGameFinished(0,0,player),ShouldBeFalse)
    player = O
    So(field.Turn(0,1,player),ShouldResemble,Field{ positions: [3][3]int{{1,2,0},{0,0,0},{0,0,0}}})
    So(field.isGameFinished(0,1,player),ShouldBeFalse)
    player = X
    So(field.Turn(2,2,player),ShouldResemble,Field{ positions: [3][3]int{{1,2,0},{0,0,0},{0,0,1}}})
    So(field.isGameFinished(2,2,player),ShouldBeFalse)
    player = O
    So(field.Turn(0,2,player),ShouldResemble,Field{ positions: [3][3]int{{1,2,2},{0,0,0},{0,0,1}}})
    So(field.isGameFinished(2,2,player),ShouldBeFalse)
    player = X
    So(field.Turn(1,1,player),ShouldResemble,Field{ positions: [3][3]int{{1,2,2},{0,1,0},{0,0,1}}})
    So(field.isGameFinished(1,1,player),ShouldBeTrue)
  })

  Convey("O is winner", t, func() {
    field := NewField()

    player := O
    So(field.Turn(0,0,player),ShouldResemble,Field{ positions: [3][3]int{{2,0,0},{0,0,0},{0,0,0}}})
    So(field.isGameFinished(0,0,player),ShouldBeFalse)
    player = X
    So(field.Turn(0,1,player),ShouldResemble,Field{ positions: [3][3]int{{2,1,0},{0,0,0},{0,0,0}}})
    So(field.isGameFinished(0,1,player),ShouldBeFalse)
    player = O
    So(field.Turn(2,2,player),ShouldResemble,Field{ positions: [3][3]int{{2,1,0},{0,0,0},{0,0,2}}})
    So(field.isGameFinished(2,2,player),ShouldBeFalse)
    player = X
    So(field.Turn(0,2,player),ShouldResemble,Field{ positions: [3][3]int{{2,1,1},{0,0,0},{0,0,2}}})
    So(field.isGameFinished(2,2,player),ShouldBeFalse)
    player = O
    So(field.Turn(1,1,player),ShouldResemble,Field{ positions: [3][3]int{{2,1,1},{0,2,0},{0,0,2}}})
    So(field.isGameFinished(1,1,player),ShouldBeTrue)
  })

  Convey("Tie", t, func() {
    field := NewField()
    player := X
    So(field.Turn(0,0,player),ShouldResemble,Field{ positions: [3][3]int{{1,0,0},{0,0,0},{0,0,0}}})
    So(field.isGameFinished(0,0,player),ShouldBeFalse)
    player = O
    So(field.Turn(0,2,player),ShouldResemble,Field{ positions: [3][3]int{{1,0,2},{0,0,0},{0,0,0}}})
    So(field.isGameFinished(0,2,player),ShouldBeFalse)
    player = X
    So(field.Turn(0,1,player),ShouldResemble,Field{ positions: [3][3]int{{1,1,2},{0,0,0},{0,0,0}}})
    So(field.isGameFinished(0,1,player),ShouldBeFalse)
    player = O
    So(field.Turn(1,0,player),ShouldResemble,Field{ positions: [3][3]int{{1,1,2},{2,0,0},{0,0,0}}})
    So(field.isGameFinished(1,0,player),ShouldBeFalse)
    player = X
    So(field.Turn(1,2,player),ShouldResemble,Field{ positions: [3][3]int{{1,1,2},{2,0,1},{0,0,0}}})
    So(field.isGameFinished(1,2,player),ShouldBeFalse)
    player = O
    So(field.Turn(1,1,player),ShouldResemble,Field{ positions: [3][3]int{{1,1,2},{2,2,1},{0,0,0}}})
    So(field.isGameFinished(1,1,player),ShouldBeFalse)
    player = X
    So(field.Turn(2,0,player),ShouldResemble,Field{ positions: [3][3]int{{1,1,2},{2,2,1},{1,0,0}}})
    So(field.isGameFinished(2,0,player),ShouldBeFalse)
    player = O
    So(field.Turn(2,1,player),ShouldResemble,Field{ positions: [3][3]int{{1,1,2},{2,2,1},{1,2,0}}})
    So(field.isGameFinished(2,1,player),ShouldBeFalse)
    player = X
    So(field.Turn(2,2,player),ShouldResemble,Field{ positions: [3][3]int{{1,1,2},{2,2,1},{1,2,1}}})
    So(field.isGameFinished(2,2,player),ShouldBeFalse)
  })
}
