package main

/*
You are given two classes, Person and Student, where Person is the base class and Student is the derived class. 
Completed code for Person and a declaration for Student are provided for you in the editor. 
Observe that Student inherits all the properties of Person.

Complete the Student class by writing the following:

A Student class constructor, which has 4 parameters:
A string, firstName.
A string, lastName.
An integer, idNumber.
An integer array (or vector) of test scores, scores.
A char calculate() method that calculates a Student object's 
average and returns the grade character representative of their calculated average:

Letter | Average
O		90 <= a <100
E		80 <= a <90
A		70 <= a < 80
P		55 <= a < 70
D		40 <= a < 55
T		a < 40
*/

// ToDo Do this in Golang

'use strict';

var _input = '';
var _index = 0;
process.stdin.on('data', (data) => { _input += data; });
process.stdin.on('end', () => {
    _input = _input.split(new RegExp('[ \n]+'));
    main();    
});
function read() { return _input[_index++]; }

/**** Ignore above this line. ****/

class Person {
    constructor(firstName, lastName, identification) {
        this.firstName = firstName;
        this.lastName = lastName;
        this.idNumber = identification;
    }
    
    printPerson() {
        console.log(
            "Name: " + this.lastName + ", " + this.firstName 
            + "\nID: " + this.idNumber
        )
    }
}

class Student extends Person {
    /*	
    *   Class Constructor
    *   
    *   @param firstName - A string denoting the Person's first name.
    *   @param lastName - A string denoting the Person's last name.
    *   @param id - An integer denoting the Person's ID number.
    *   @param scores - An array of integers denoting the Person's test scores.
    */
    // Write your constructor here
    constructor(firstName, lastName, id, scores){
        super()
        this.firstName = firstName
        this.lastName = lastName
        this.idNumber = id
        this.scores = scores
    }
    /*	
    *   Method Name: calculate
    *   @return A character denoting the grade.
    */
    // Write your method here
    calculate(){
        const total = this.scores.reduce((a,b) => a+b,0 )
        const average = total/this.scores.length
        if(average < 40){
            return 'T'
        }
        if(average < 55){
            return 'D'
        }
        if(average < 70){
            return 'P'
        }
        if(average < 80){
            return 'A'
        }
        if(average < 90){
            return 'E'
        }
        return 'O'
    }
}

function main() {
    let firstName = read()
    let lastName = read()
    let id = +read()
    let numScores = +read()
    let testScores = new Array(numScores)
    
    for (var i = 0; i < numScores; i++) {
        testScores[i] = +read()  
    }

    let s = new Student(firstName, lastName, id, testScores)
    s.printPerson()
    s.calculate()
    console.log('Grade: ' + s.calculate())
}