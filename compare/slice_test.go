package compare_test

import (
	"github.com/gushasha/go-learn/compare"
	"reflect"
	"testing"
)

type StringsTester struct {
	a       []string
	b       []string
	isEqual bool
}

func TestStrings(t *testing.T) {
	tester := StringsTester{
		[]string{"foo", "bar"},
		[]string{"foo", "bar"},
		true,
	}

	if actual := compare.Strings(tester.a, tester.b); actual != tester.isEqual {
		t.Errorf("%#v compare %#v fail, \n expected %v, got %v \n", tester.a, tester.b, tester.isEqual, actual)
	}
}

func TestInts(t *testing.T) {
	slice1 := []int{1, 2}
	slice2 := []int{1}
	isEqual := false

	if actual := compare.Ints(slice1, slice2); actual != isEqual {
		t.Errorf("%#v compare %#v fail, \n expected %v, got %v \n", slice1, slice2, isEqual, actual)
	}
}

func BenchmarkStrings(b *testing.B) {
	slice1 := []string{"foo", "bar", "h", "e", "l", "l", "o"}
	slice2 := []string{"foo", "bar", "h", "e", "l", "l", "oooo"}

	for i := 0; i < b.N; i++ {
		compare.Strings(slice1, slice2)
	}

}

func BenchmarkRelfectDeepEqual(b *testing.B) {
	slice1 := []string{"foo", "bar", "h", "e", "l", "l", "o"}
	slice2 := []string{"foo", "bar", "h", "e", "l", "l", "oooo"}

	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(slice1, slice2)
	}
}
