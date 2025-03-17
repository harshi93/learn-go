package main

/* with recursion */
func fibWRec(position int, seq []int) int {
	if seq == nil {
		// we are initializing the sequence with the first two numbers
		seq = []int{0, 1}
	}

	for len(seq) <= position {
		seq = append(seq, seq[len(seq)-1]+seq[len(seq)-2])
		return fibWRec(position, seq)
	}

	// returns the last number in the sequence
	return seq[len(seq)-1]
}

/* without recursion */
func fibWoRec(position int) int {
	// we are initializing the sequence with the first two numbers
	seq := []int{0, 1}

	for len(seq) <= position {
		seq = append(seq, seq[len(seq)-1]+seq[len(seq)-2])
	}

	// returns the last number in the sequence
	return seq[len(seq)-1]
}
