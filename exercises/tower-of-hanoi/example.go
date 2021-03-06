package towerofhanoi

// This is an example implementation.

// Solve returns list of moves for solving the Tower of Hanoi puzzle.
//
// 'disks' is the amount of disks.
// 'from' is the id for the source rod.
// 'to' is the id for the destination rod.
// 'via' is the id for the helper rod.
//
// A Move value specifies source and destination rods.
//
// Example
//
// In order to illustrate the problem, consider the constants below
// naming the rods. Your implementation is free to define similar
// constants, but it is not required to do so.
// const (
//      A = 'A'
//      B = 'B'
//      C = 'C'
// )
//
// Solve(1, A, B, C) could return a list with a single move:
//  []Move{{A, B}} // move disk from rod A to rod B
func Solve(disks, from, to, via int) []Move {

	if disks < 1 {
		return nil // refuse less than 1 disk
	}

	if from == to || from == via || to == via {
		return nil // refuse dup rods
	}

	fromTo := Move{from, to}

	if disks == 1 {
		return []Move{fromTo} // single disk on rod FROM
	}

	// freeze largest disk on rod FROM,
	// then move all other disks to rod VIA
	// (using rod TO as helper)
	m1 := Solve(disks-1, from, via, to)

	// move largest disk to rod TO
	m2 := append(m1, fromTo)

	// move all remaining disks from rod VIA to rod TO
	// (using rod FROM as helper)
	return append(m2, Solve(disks-1, via, to, from)...)
}
