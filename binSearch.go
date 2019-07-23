package binSearch

import(
  //"fmt"
  "errors"
)

type Note struct{
  address     int
  value       interface{}
  rightBranch *Note
  leftBranch  *Note
}

type Cache struct{
  name  string
  root  Note
}

func New(name string) *Cache {
  root := Note{
    address:     -1,
    value:       nil,
    rightBranch: nil,
    leftBranch:  nil,
	}

	cache := Cache{
		name:         name,
		root:         root,
	}

	return &cache
}

func (n *Note) findValue(address int) (interface{}, error)  {
  if n.address == -1 {
    return nil, errors.New("Address not found")
  }

  if n.address == address {
    return n.value, nil
  }

  if n.address > address {
    return n.rightBranch.findValue(address)
  }

  if n.address < address {
    return n.leftBranch.findValue(address)
  }
  return nil, nil
}

func (c *Cache) Find(address int) (interface{}, error)  {
  return c.root.findValue(address)
}


func (n *Note) insertValue(address int, value interface{}) error  {
  tn := Note{
		address:     -1,
		value:       nil,
		rightBranch: nil,
		leftBranch:  nil,
	}

  if n.address == -1 {
    n.address = address
		n.value = value
		n.rightBranch = &tn
		n.leftBranch = &tn

    return nil
  }

  if n.address == address {
    n.value = value
    return nil
  }

  if n.address > address {
    return n.rightBranch.insertValue(address, value)
  }

  if n.address < address {
    return n.leftBranch.insertValue(address, value)
  }
  return nil
}

func (c *Cache) Insert(address int, value interface{}) error  {
  return c.root.insertValue(address, value)
}
