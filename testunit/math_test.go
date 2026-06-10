package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMyIntAdd(t *testing.T) {
	data := []struct {
		title  string
		value  MyInt
		param  int
		should MyInt
	}{
		{title: "A", value: 1, param: 1, should: 2},
		{"B", 2, 1, 3},
		{"C", 9, 1, 10},
	}
	for _, v := range data {
		mi := v.value
		mi.Add(v.param)
		if mi != v.should {
			t.Error("for", v.title, "got", mi, "should got", v.should)
		}
	}

	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		var x MyInt = 1

		Convey("When the integer is incremented", func() {
			x.Add(1)

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}
