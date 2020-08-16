package gomda

import "testing"

type a struct {
	FieldOne string
	FieldTwo int
}

type b struct {
	FieldOne bool
	FieldTwo a
}

func TestEquals_ShouldReturnTrueIfIntValuesAreSame(t *testing.T) {
	n1 := 1
	n2 := 1
	if !Equals(n1, n2) {
		t.Fail()
	}
}

func TestEquals_ShouldReturnTrueIfFloatValuesAreSame(t *testing.T) {
	f1 := 1.1
	f2 := 1.1
	if !Equals(f1, f2) {
		t.Fail()
	}
}

func TestEquals_ShouldReturnTrueIfComplexValuesAreSame(t *testing.T) {
	c1 := complex64(1)
	c2 := complex64(1)
	if !Equals(c1, c2) {
		t.Fail()
	}
}

func TestEquals_ShouldReturnTrueIfBooleanValuesAreSame(t *testing.T) {
	if !Equals(true, true) {
		t.Fail()
	}

	if !Equals(false, false) {
		t.Fail()
	}
}

func TestEquals_ShouldReturnTrueIfStringValuesAreSame(t *testing.T) {
	s1 := "hello"
	s2 := "hello"
	if !Equals(s1, s2) {
		t.Fail()
	}
}

func TestEquals_ShouldReturnTrueIfNoNestedStructValuesAreSame(t *testing.T) {
	a1 := a{"1", 1}
	a2 := a{"1", 1}

	if !Equals(a1, a2) {
		t.Fail()
	}
}

func TestEqual_ShouldReturnTrueIfNestedStructValuesAreSame(t *testing.T) {
	b1 := b{false, a{"hello", 1}}
	b2 := b{false, a{"hello", 1}}

	if !Equals(b1, b2) {
		t.Fail()
	}
}

func TestEquals_ShouldReturnTrueIfArrayValuesAreSame(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}

	if !Equals(a, b) {
		t.Fail()
	}
}

func TestEquals_ShouldReturnFalseIfArrayValuesAreNotSame(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 1, 3, 4, 5}
	c := [2]int{1, 2}

	if Equals(a, b) || Equals(a, c) {
		t.Fail()
	}
}

func TestEquals_ShouldReturnTrueIfSliceValuesAreSame(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	if !Equals(a, b) {
		t.Fail()
	}
}

func TestEquals_ShouldReturnFalseIfSliceValuesAreNotSame(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 1, 3}
	c := []int{1}

	if Equals(a, b) || Equals(a, c) {
		t.Fail()
	}
}

func TestEquals_ShouldReturnTrueIfSliceValuesWithStructInThemAreSame(t *testing.T) {
	c := []a{{"", 1}}
	d := []a{{"", 1}}

	if !Equals(c, d) {
		t.Fail()
	}
}

func TestEquals_ShouldReturnTrueIfMapValuesAreSame(t *testing.T) {
	m1 := make(map[string]int)
	m1["a"] = 1
	m1["b"] = 2

	m2 := make(map[string]int)
	m2["b"] = 2
	m2["a"] = 1

	if !Equals(m1, m2) {
		t.Fail()
	}
}

func TestEquals_ShouldReturnFalseIfMapValuesAreDifferent(t *testing.T) {
	m1 := make(map[string]int)
	m1["a"] = 1
	m1["b"] = 2

	m2 := make(map[string]int)
	m2["B"] = 2
	m2["a"] = 1

	if Equals(m1, m2) {
		t.Fail()
	}
}

func TestEquals_ShouldReturnTrueIfPointerValuesAreSame(t *testing.T) {
	p1 := 1
	p2 := 1

	if !Equals(&p1, &p2) {
		t.Fail()
	}
}

func TestEquals_ShouldReturnTrueIfValuesAreNil(t *testing.T) {
	if !Equals(nil, nil) {
		t.Fail()
	}
}

func TestEquals_ShouldReturnFalseIfFirstParamIsNil(t *testing.T) {
	if Equals(1, nil) {
		t.Fail()
	}
}

func TestEquals_ShouldReturnFalseIfSecondParamIsNil(t *testing.T) {
	if Equals(nil, 1) {
		t.Fail()
	}
}
