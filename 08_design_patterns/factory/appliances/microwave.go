package appliances

// Concrete type of Appliance
type Microwave struct {
	typeName string
}

func (self *Microwave) Start() {
	self.typeName = "Microwave"
}

func (self *Microwave) GetPurpose() string {
	return "I am a " + self.typeName + ", I heat things up."
}
