package shortcut

type Command interface {
	Do()
	Undo()
	Name() string
}
