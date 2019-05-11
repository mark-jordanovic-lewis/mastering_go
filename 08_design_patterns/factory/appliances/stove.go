package appliances

// Concrete type of Appliance
type Stove struct {
	typeName string
}

func (self *Stove) Start() {
	self.typeName = "Stove"
}

func (self *Stove) GetPurpose() string {
	return "I am a " + self.typeName + ", I heat things up."
}
