//Данный пакет определяет данные, получаемые из конфига, и то, как они будут выводится
package configTypes

import (
	"encoding/xml"
	"fmt"
)

//Данная структура определяет конкретные данные(теги) конфига
type Person struct {
	XMLName   xml.Name `xml:"Person"`
	FirstName string   `xml:"FirstName"`
	LastName  string   `xml:"LastName"`
	Age       uint     `xml:"Age"`
}

//Данная структура определяет, какие данные(теги заключенные внутри тега "Data", будут сичтаны)
type Data struct {
	XMLName    xml.Name `xml:"Data"`
	PersonList []Person `xml:"Person"`
}

//Выводит данные конфига в заданном формате
func (p Person) String() string {
	return fmt.Sprintf("Person:\n%s\n%s\n%v\n", p.FirstName, p.LastName, p.Age)
}
