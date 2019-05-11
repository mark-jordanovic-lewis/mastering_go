package appliances

// Concrete type of Appliance
type Fridge struct {
	typeName string
}

func (self *Fridge) Start() {
	self.typeName = "Fridge"
}

func (self *Fridge) GetPurpose() string {
	return "I am a " + self.typeName + ", I cool things down."
}