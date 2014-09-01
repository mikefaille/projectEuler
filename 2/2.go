package main

type NumbersMem struct {
	lastNumber int
	currentNumber int
}

func operFibonacci(number *NumbersMem) NumbersMem {
	//println("before ", number.currentNumber)
	newNumber := NumbersMem{number.currentNumber, number.currentNumber + number.lastNumber}
	//println("after ", newNumber.currentNumber)
	return newNumber
}

func main(){
	var sumEvenFibo int
	// init Fibonacci
	var number NumbersMem
	number = NumbersMem{1, 2}

	for number.currentNumber<4000000  {
		println("new loop")
		println("even ? ", number.currentNumber)
		if  (number.currentNumber % 2) == 0 {
			sumEvenFibo += number.currentNumber
			println(sumEvenFibo)
		}
		number = operFibonacci(&number )

		println("if < 4000000 ? loop it! ", number.currentNumber)
	}

	println(sumEvenFibo)

}
