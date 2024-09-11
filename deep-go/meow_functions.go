package deep_go

// MeowTimes is a sample function that returns a string of "meow" repeated n times
// For example, MeowTimes(3) returns "meowmeowmeow"
func MeowTimes(n int) string {
	var meows string
	for i := 0; i < n; i++ {
		meows += "meow"
	}
	return meows
}
