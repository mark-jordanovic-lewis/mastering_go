package main
// integer only

import (
  "fmt"
  "math/rand"
  "time"
  "errors"
)

// Node \\
type Node interface {
  Set(n int) error
  Get() (int, error)
}

var ErrInvalidNode = errors.New("Invalid node")

// SLL Nodes \\
type SLLNode struct {
  next *SLLNode
  value int
  SLLNodeMessage string
}

func (self *SLLNode) Set(n int) error {
  if self == nil { return ErrInvalidNode }
  self.value = n
  return nil
}

func (self *SLLNode) Get() (int, error) {
  if self == nil { return -1, ErrInvalidNode }
  return self.value, nil
}

func NewSLLNode() *SLLNode {
  return &SLLNode{ SLLNodeMessage: "Std SLLNode" }
}


// Power Nodes \\
type PowerNode struct {
  next *PowerNode
  value int
  PowerNodeMessage string
}

func (self *PowerNode) Set(n int) error {
  if self == nil { return ErrInvalidNode }
  self.value = n*10
  return nil
}

func (self *PowerNode) Get() (int, error) {
  if self == nil { return -1, ErrInvalidNode }
  return self.value, nil
}

func NewPowerNode() *PowerNode {
  return &PowerNode{ PowerNodeMessage: "Power SSLNode"}
}


// SLL \\
type SLL struct {
  head *SLLNode
  last *SLLNode
}

func (self *SLL) Add(val int) {
  new_node := &SLLNode { value: val }
  if self.head == nil {
    self.head = new_node
  } else if self.last == self.head {
    self.head.next = new_node
  } else if self.last != nil {
    self.last.next = new_node
  }
  self.last = new_node
}

func (self *SLL) String() (out string) {
  out = "SLL["
  for n := self.head; n != nil; n = n.next {
    if n != self.head {
      out += "->"
    }
    val, _ := n.Get()
    out += fmt.Sprintf(" %d ", val)
  }
  out += "]"
  return
}

func main() {
  var node Node
  node = NewSLLNode()
  node.Set(4)
  val, _ := node.Get()
  fmt.Println("node val:", val)

  node = NewPowerNode()
  node.Set(4)
  val, _ = node.Get()
  fmt.Println("node val:", val)

  // checking that concrete type is implementor of the interface of the declared var: <varname>.(*<Type>)
  if _, ok := node.(*PowerNode); ok {
    fmt.Println("The node is a power node")
  }

  list := SLL{}
  list.Add(1)
  list.Add(2)
  list.Add(5)
  list.Add(7)
  fmt.Println(list.String())

  node = createNode()
  switch concretion := node.(type) {
  case *SLLNode:
    fmt.Println("node is of type: ", concretion.SLLNodeMessage)
  case *PowerNode:
    fmt.Println("node is of type: ", concretion.PowerNodeMessage)
  default:
    fmt.Println("unknown node type")
  }

  var nil_node *SLLNode
  if _, err := nil_node.Get(); err != nil { fmt.Println("Got error:", err) }
  if err := nil_node.Set(3); err != nil { fmt.Println("Got error:", err) }
}

func createNode() Node {
  rand.Seed(time.Now().UnixNano())
  flip := rand.Float64()
  var node Node
  if flip < 0.5 {
    node = NewSLLNode()
  } else {
    node = NewPowerNode()
  }
  node.Set(4)
  return node
}
