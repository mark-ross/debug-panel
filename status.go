package debug_panel

// Status is an enumerated value that describes
// the state of the progress flag.
type Status string

const(
	// Failure is when some error is reported.
	// Note, this doesn't display the error, only reports
	// that an error occurred in the flag.
	Failure Status = "Failure"

	// NotStarted, as the Name suggests, should be used
	// when the code portion has not yet been reached
	NotStarted Status = "Not Started"

	// InProgress describes when the designated portion
	// of code is still calculating which should be
	// useful for long running calculations
	InProgress Status = "In Progress"

	// Complete should be used when the calculation is
	// completed -- or, in the case of long-running
	// computations, after a piece of the code has been
	// reached.
	Complete Status = "Complete"

	// Waiting should be use if you have waiting Go Routines
	// waiting for work to be given to them.
	Waiting Status = "Waiting"
)

