package main

import (
	"awesomeProject/model"
	"fmt"
	"math"
	"sync"
)
func printString ( n string, wg *sync.WaitGroup) {
	result := 0.0
	for i := 0;i<1000000000;i++ {
		result +=math.Pi * math.Sin(float64(10))
	}
	fmt.Println(n)
	wg.Done()
}
func main() {
var wg sync.WaitGroup

	jumperList := model.GetList()
	wg.Add(len(jumperList))
	for _, jumper := range jumperList {
		go printString(jumper.Jump(),&wg)
	}
wg.Wait()

	/*
langs := getLangs()


	gopher1 :=  gopher{name:"Phil", age: 30}
	gopher2 :=  &gopher{name:"Noodles", age: 90}


	fmt.Println(gopher1.jump());
	fmt.Println(gopher2.jump());
	fmt.Println(gopher2);
	validateAge(gopher2)
	fmt.Println(gopher2);



for _,element := range langs {
	fmt.Println(element)
}

	hourOfDay := time.Now().Hour()
	greeting, err := getGreeting ( hourOfDay )
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(greeting)

*/


}



