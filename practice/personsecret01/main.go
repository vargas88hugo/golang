package main

import (
	"fmt"
)

func main() {
	varPerson := person{"Hugo", "Vargas"}
	varSecretAgent := secretAgent{person{"James", "Bond"}, true}
	fmt.Println(varPerson.firstName)
	varPerson.personSpeak()
	fmt.Println(varSecretAgent.firstName)
	varSecretAgent.personSpeak()
	varSecretAgent.secretAgentSpeak()
}

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) personSpeak() {
	fmt.Println(p.firstName, `says, "Wassup!"`)
}

func (s secretAgent) secretAgentSpeak() {
	fmt.Println(s.firstName, "has a license to kill:", s.licenseToKill)
}
