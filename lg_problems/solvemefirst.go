func solveMeFirst(a uint32, b uint32) uint32 {
	// Hint: Type return (a+b) below
	if constrain1(a) && constrain1(b) {
		return a + b
	} else {
		return 0
	}

}