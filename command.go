package cmd

type Command struct {
	Name        string
	Description string
	Flags       []Flag
	Execute     func(flags Flags, args []string)
}
