package ConfigTypes

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName   xml.Name `xml:"Person"`
	FirstName string   `xml:"FirstName"`
	LastName  string   `xml:"LastName"`
	Age       uint     `xml:"Age"`
}

type Data struct {
	XMLName    xml.Name `xml:"Data"`
	PersonList []Person `xml:"Person"`
}

func (p Person) String() string {
	return fmt.Sprintf("Person:\n%s\n%s\n%v\n", p.FirstName, p.LastName, p.Age)
}
