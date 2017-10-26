package stringutil

// Reverse  calls non visiable reverseTwo to show the difference between hidden
// and non hidden functions. Capital letter fucntion names are non hidden.
func Reverse(s string) string {
	return reverseTwo(s)
}
