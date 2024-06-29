//NESTED STRUCTURE
// type car struct {
// 	make       string
// 	model      string
// 	doors      int
// 	mileage    int
// 	frontWheel wheel
// 	backWheel  wheel
// }

// type wheel struct {
// 	radius   int
// 	material string
// }

// The fields of a struct can be accessed using the dot (.) operator.
// myCar := car{}  ->instantiate a struct here, this is an empty struct.
// myCar.frontWheel.radius = 5

// package main

// type messageToSend struct {
// 	message   string
// 	sender    user
// 	recipient user
// }

// type user struct {
// 	name   string
// 	number int
// }





type rect struct {
	width int
	height int
  }
  
  // area has a receiver of (r rect)
  func (r rect) area() int {
	return r.width * r.height
  }
  
  var r = rect{
	width: 5,
	height: 10,
  }
  
  fmt.Println(r.area())
  // prints 50

