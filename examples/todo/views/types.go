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

func todoCheckStyle(done bool) string {
	if done {
		return "background:#4f46e5;border-color:#4f46e5"
	}
	return "border-color:#d1d5db;background:transparent"
}

func todoActiveCount(todos []Todo) int {
	count := 0
	for _, t := range todos {
		if !t.Done {
			count++
		}
	}
	return count
}

func todoAllDone(todos []Todo) bool {
	if len(todos) == 0 {
		return false
	}
	for _, t := range todos {
		if !t.Done {
			return false
		}
	}
	return true
}
