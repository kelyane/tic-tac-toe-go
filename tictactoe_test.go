package tictactoe

import (
  . "github.com/smartystreets/goconvey/convey"
  "testing"
)

func TestField(t *testing.T){
  Convey("Field", t, func() {
    field := NewField()
    So(field.isEmpty(),ShouldBeTrue)
  })
}
