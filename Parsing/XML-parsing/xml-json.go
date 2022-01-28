package main

import (
	"encoding/xml"
	"fmt"
)

//A Person has a list of Addresses. Since we
//know that address is a direct descendant,
//we can use the > keyword to parse
//each <address> from <addresses> into an
//embedded struct on our Person struct.

type Person struct {
	Name      string `xml:"name"`
	Addresses []struct {
		Street string `xml:"street"`
		City   string `xml:"city"`
		Type   string `xml:"type,attr"`
	} `xml:"addresses>address"`
}

var Santosh Person
var document = []byte(
	`
	<person>
	<name>Santosh Suryawanshi</name>
	<addresses>
	<address type="secondary">
	<street>321 MG Road</street>
	<city>Pune</city>
	</address>
	<address type="primary">
	<street>St Paul Street </street>
	<city>Mumbai</city>
	</address>
	</addresses>
	</person>
	`)

func main() {
	xml.Unmarshal(document, &Santosh)
	fmt.Println(Santosh)
}
