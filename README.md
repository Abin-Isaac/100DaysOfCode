
## Learning Source :- https://golangbot.com/

## First Day -
Had the first ever interaction with Go- Language of Google // Printed on the console along with basic maths operations.

## 2nd day - Learnt something new about new as 

1) Every Go file starts with a package name. Here we have used package main
2) func main comes under package main 
3) In the Go file other packages will be imported by import keyword
4) We have imported fmt package which will be used inside main fucntion to print o/p to console.
5) 2 ways to write import "fmt" or import ("fmt")
6) main is a special function as always and the program execution starts from there. Braces are the same as it was in C
7) For Calling a function we need to use package.fucntion() syntax 

Run the code here for first and second day  https://play.golang.org/p/igXSORua9we

## 3rd day - Learnt integer varaibles

A)Declaration synatx is like - var age int

B)Two places to declare variable 1) After import "fmt" 2) Inside main before usage as always

C)We reassigned that varaible in the progarm and printed the value...dont know if it is immutable or not ?

D) Default value is 0 for int 

Run the code here for the third day https://play.golang.org/p/7GXCJViKu5W

## 4th day - Variables with initial values 

1) Declaration synatx is like - var age = value
2) No requirement to assign type as Go will automatically infer the type.
3) var age   = 'a' is printing 97 which is ASCII value of a so it may be taking the int type by default ??

Run the code here :- https://play.golang.org/p/4uEAlOFi3OB

## 5th day - Multiple Variables in one line

There are two ways to decalre multiple varaibeles in one line depdning on the type 
1) Same data type - var width, height int = 100, 50
2) Different data type - var 	(a=10; b='c';c=8.9)

Note : If you try assigning more than one character then you will get "more than one character in rune literal"..??

Run the code here :- https://play.golang.org/p/lWshD_L6NQU

## 6th day - Go with short hand...No var keyword required

1) := is the sign of short hand in variable initialization.
2) We can use in single or multi varaible initialization in one line like ---> name, age := "naveen", 29
3) Compiler will be angry if we leave one varaible uninitialized...you will get an error 
4) And there should be at least one new varaible on the left of := for initialization 

Run the code here :- https://play.golang.org/p/YfdJ8X6I8jv

