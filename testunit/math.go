package main

type MyInt int

func (mi *MyInt) Add(i int) {
	*mi = *mi + MyInt(i)
	// go Read(mi)
	// go Write(mi)
}

func Read(m *MyInt) {
	_ = m
}

func Write(m *MyInt) {
	*m = 1
}
