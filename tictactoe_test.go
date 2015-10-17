package tictactoe

import (
  . "github.com/smartystreets/goconvey/convey"
  "testing"
)

func initial_test(t *testing.T){
  Convey("Field", t, func() {
    field := NewField()
    So(field.isEmpty(),ShouldBeTrue)
  })
}
