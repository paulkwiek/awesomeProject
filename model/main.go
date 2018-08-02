package model

import ( "fmt"
		"errors"
    )
type jumper interface {
	Jump() string
}
type gopher struct {
	name string
	age int
	isAdult bool
}

type horse struct {
	name string
	weight int
}

func (h horse ) Jump() string {
	if h.weight > 2500 {
		return "I'm too heavy, can't jump..."
	}
	return "I will jump, neigh!!"
}


func (g gopher) Jump() string {

	if g.age < 65 {
		return g.name + " can jump HIGH"
	}
	return g.name + " can still jump"
}

func validateAge(g *gopher) {
	g.isAdult = g.age >= 21
}

func highJump(name string) {
	fmt.Println(name,"can jump HIGH")
}

func lowJump(name string) {
	fmt.Println(name,"can still jump")
}

func getLangs () []string {
langs := []string {"go", "rubby", "javascript"}
fmt.Println(langs[0])
fmt.Println(langs[1])
fmt.Println(langs[2])
return langs
}

func GetList() []jumper {
	phil :=  &gopher{name:"Phil", age: 30}
	noodles :=  &gopher{name:"Noodles", age: 90}
	barbaro := &horse{name:"Barbaro",weight: 2600}
	list := []jumper{phil, noodles, barbaro}
	return list
}

func getGreeting(hour int) (string , error) {
	var message string
	if hour < 7 {
		err := errors.New("Too early for greatings!")
		return message, err
	}
	if hour < 12 {
		message = "Good morning"
	} else if hour < 18 {
		message = "Good Afternoon"
	} else {
		message =  "Good Evening"
	}
	return message, nil
}



