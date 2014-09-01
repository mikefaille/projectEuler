package main

func getSubtotal(max int, element int) int {
	subsum := 0
	var subtotal int
	//m_5 := []int(0)

	for i := 0; i < ((max-1)/element); i++ {
		subsum = subsum + element
		subtotal += subsum
		println("subtotal no" , i+1 , "=" , subtotal , "\n cause we add" ,  subsum)

	}

	return subtotal

}

func main() {

	max := 1000
	sum_3 :=getSubtotal(max,3)


	sum_5:= getSubtotal(max,5)

	sum_15  := getSubtotal(max,15) // inclusion exclusion principle https://fr.wikipedia.org/wiki/Principe_d'inclusion-exclusion


	total := sum_3 + sum_5 - sum_15 // total of unique numbers

	print(total)

}


