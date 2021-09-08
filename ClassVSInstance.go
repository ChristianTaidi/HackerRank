package main
/*
Write a Person class with an instance variable, age, and a constructor that takes an integer, initialAge, as a parameter.
The constructor must assign initialAge to age after confirming the argument passed as initialAge is not negative; 
if a negative argument is passed as initialAge, the constructor should set age to 0 and print Age is not valid, setting age to 0.. 
In addition, you must write the following instance methods:

yearPasses() should increase the age instance variable by 1.
amIOld() should perform the following conditional actions:
If age < 13, print You are young..
If age >= 13 and age < 18, print You are a teenager..
Otherwise, print You are old..
To help you learn by example and complete this challenge, much of the code is provided for you, but you'll be writing everything in the future. The code that creates each instance of your Person class is in the main method. Don't worry if you don't understand it all quite yet!
*/

import "fmt"

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	if initialAge < 0 {
        fmt.Print("Age is not valid, setting age to 0.\n")
        initialAge = 0
    }
    p.age =  initialAge
    return p
}

func (p person) amIOld() {
	if p.age < 13 {
        fmt.Print("You are young.\n")
        return
    }
    if p.age < 18 {
        fmt.Print("You are a teenager.\n")
        return
    }
    fmt.Print("You are old.\n")
}

func (p person) yearPasses() person {
	p.age += 1
    return p
}

func main() {
    var T,age int

    fmt.Scan(&T)

    for i := 0; i < T; i++ {
        fmt.Scan(&age)
        p := person{age: age}
        p = p.NewPerson(age)
        p.amIOld()
        for j := 0; j < 3; j++ {
            p = p.yearPasses()
        }
        p.amIOld()
        fmt.Println()
    }
}
