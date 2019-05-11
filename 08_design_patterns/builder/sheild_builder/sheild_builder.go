package sheild_builder

import "strings"

type sheild struct {
	front bool
	back bool
	right bool
	left bool
}

type sheildBuilder struct {
	code string
}

func NewSheildBuilder() *sheildBuilder {
	return &sheildBuilder{}
}

// not sure why the builder is returned and mutated...
func (self *sheildBuilder) RaiseFront() *sheildBuilder  {
	self.code +="F"
	return self
}

func (self *sheildBuilder) RaiseBack() *sheildBuilder  {
	self.code +="B"
	return self
}

func (self *sheildBuilder) RaiseRight() *sheildBuilder  {
	self.code +="R"
	return self
}

func (self *sheildBuilder) RaiseLeft() *sheildBuilder  {
	self.code +="L"
	return self
}

// could have been implemented using bitmaps
func (self *sheildBuilder) Build() *sheild {
	code := self.code
	return &sheild{
		front: strings.Contains(code, "F"),
		back: strings.Contains(code, "B"),
		right: strings.Contains(code, "R"),
		left: strings.Contains(code, "L"),
	}
}
