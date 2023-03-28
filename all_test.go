package go_warmup

import (
	concurrency "go_warmup/Concurrency"
	conditionals "go_warmup/Conditionals"
	primitives "go_warmup/Variables"
	"reflect"
	"testing"
)

func Test_primitive_strings(t *testing.T) {
	a, b, c := primitives.Strings()
	if reflect.TypeOf(a).Kind() != reflect.String {
		t.Fatal("Variable a is not a string!")
	}
	if reflect.TypeOf(b).Kind() != reflect.String {
		t.Fatal("Variable b is not a string!")
	}
	if reflect.TypeOf(c).Kind() != reflect.String {
		t.Fatal("Variable c is not a string!")
	}
}

func Test_primitive_bool(t *testing.T) {
	if primitives.Boolean(false, false) {
		t.Fatal("Or returned True to two false inputs")
	}

	if !primitives.Boolean(true, false) {
		t.Fatal("Or returned False to one false and one true input")
	}

	if !primitives.Boolean(false, true) {
		t.Fatal("Or returned False to one true and one false input")
	}

	if !primitives.Boolean(true, true) {
		t.Fatal("Or returned False to two true inputs")
	}
}

func Test_conditionals_switch(t *testing.T) {
	x := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	y := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for i := 0; i < 10; i++ {
		test := conditionals.Switch_statement(x[i])
		if y[i] != test {
			t.Fatal(y[i] + " does not equal '" + test + "'")
		}
	}
}

func Test_buffered_channel(t *testing.T) {
	for i := 1; i <= 10; i++ {
		bufferedChannel := concurrency.Buffered_channel(i)
		output := make(chan int, 1)
		go func(limit int, output chan int) {
			count := 0
			for {
				breakout := false
				select {
				case bufferedChannel <- string(rune(i)):
					count++
				default:
					breakout = true
				}
				if breakout {
					break
				}
			}
			output <- count
		}(i, output)

		count := <-output

		if count != i {
			t.Fatal("Count does not equal buffered channel limit of " + string(rune(i)))
		}

	}
}
