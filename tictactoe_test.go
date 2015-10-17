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
    var player = 1;
    fieldWinner := NewField()
    fieldWinner = Field{ positions: [3][3]int{{1,1,1},{0,0,0},{0,0,0}}}
    So(fieldWinner.isWinner(Player(player)),ShouldBeTrue)
    fieldWinner = Field{ positions: [3][3]int{{0,0,0},{1,1,1},{0,0,0}}}
    So(fieldWinner.isWinner(Player(player)),ShouldBeTrue)
    fieldWinner = Field{ positions: [3][3]int{{0,0,0},{0,0,0},{1,1,1}}}
    So(fieldWinner.isWinner(Player(player)),ShouldBeTrue)
    fieldWinner = Field{ positions: [3][3]int{{0,0,2},{0,0,2},{0,0,2}}}
    So(fieldWinner.isWinner(Player(player)),ShouldBeFalse)
    fieldWinner = Field{ positions: [3][3]int{{1,0,0},{0,1,0},{0,0,1}}}
    So(fieldWinner.isWinner(Player(player)),ShouldBeTrue)
    fieldWinner = Field{ positions: [3][3]int{{0,0,1},{0,0,1},{0,0,2}}}
    So(fieldWinner.isWinner(Player(player)),ShouldBeFalse)
    fieldWinner = Field{ positions: [3][3]int{{0,1,1},{0,0,1},{1,0,2}}}
    So(fieldWinner.isWinner(Player(player)),ShouldBeFalse)
    player = 2;
    fieldWinner = Field{ positions: [3][3]int{{0,0,0},{0,0,0},{1,1,1}}}
    So(fieldWinner.isWinner(Player(player)),ShouldBeFalse)
    fieldWinner = Field{ positions: [3][3]int{{2,0,0},{2,0,0},{2,0,0}}}
    So(fieldWinner.isWinner(Player(player)),ShouldBeTrue)
    fieldWinner = Field{ positions: [3][3]int{{0,2,0},{0,2,0},{0,2,0}}}
    So(fieldWinner.isWinner(Player(player)),ShouldBeTrue)
    fieldWinner = Field{ positions: [3][3]int{{0,0,2},{0,0,2},{0,0,2}}}
    So(fieldWinner.isWinner(Player(player)),ShouldBeTrue)
    fieldWinner = Field{ positions: [3][3]int{{0,0,2},{0,2,0},{2,0,0}}}
    So(fieldWinner.isWinner(Player(player)),ShouldBeTrue)
    fieldWinner = Field{ positions: [3][3]int{{0,0,0},{0,0,0},{1,2,1}}}
    So(fieldWinner.isWinner(Player(player)),ShouldBeFalse)
    })
}
