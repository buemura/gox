package views

// Todo represents a single todo item.
type Todo struct {
	ID   int
	Text string
	Done bool
}

func todoItemClass(done bool) string {
	base := "flex items-center gap-3 p-3 bg-white rounded-lg shadow-sm"
	if done {
		return base + " opacity-60"
	}
	return base
}

func todoTextClass(done bool) string {
	if done {
		return "flex-1 text-base line-through text-gray-400"
	}
	return "flex-1 text-base text-gray-800"
}
