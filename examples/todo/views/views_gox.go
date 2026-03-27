package views

import (
	"fmt"
	"github.com/buemura/gox"
	"html"
	"io"
	"strconv"
)

func Layout(title string, children gox.Component) gox.Component {
	return gox.ComponentFunc(func(w io.Writer) error {
		var err error
		_ = err
		_, err = io.WriteString(w, "\n  <!DOCTYPE html>\n  ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<html")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " lang=\"en\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n    ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<head")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n      ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<meta")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " charset=\"UTF-8\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " />")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n      ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<meta")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " name=\"viewport\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " content=\"width=device-width, initial-scale=1.0\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " />")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n      ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<title")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", title)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</title>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n      ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<script")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " src=\"https://cdn.tailwindcss.com\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</script>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n    ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</head>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n    ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<body")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"bg-gray-100 min-h-screen\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n      ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<div")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"max-w-xl mx-auto px-5 py-10\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n        ")
		if err != nil {
			return err
		}
		if children != nil {
			err = children.Render(w)
			if err != nil {
				return err
			}
		}
		_, err = io.WriteString(w, "\n      ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n    ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</body>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n  ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</html>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n")
		if err != nil {
			return err
		}
		return nil
	})
}

func TodoItem(id int, text string, done bool) gox.Component {
	return gox.ComponentFunc(func(w io.Writer) error {
		var err error
		_ = err
		_, err = io.WriteString(w, "\n  ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<li")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", todoItemClass(done))))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n    ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<form")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " method=\"POST\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " action=\"/toggle\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n      ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<input")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " type=\"hidden\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " name=\"id\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " value=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", strconv.Itoa(id))))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " />")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n      ")
		if err != nil {
			return err
		}
		if done {
			_, err = io.WriteString(w, "\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<button")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " type=\"submit\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"text-lg\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "&#x2705;")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</button>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n      ")
			if err != nil {
				return err
			}
		} else {
			_, err = io.WriteString(w, "\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<button")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " type=\"submit\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"text-lg\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "&#x2B1C;")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</button>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n      ")
			if err != nil {
				return err
			}
		}
		_, err = io.WriteString(w, "\n    ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</form>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n    ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<span")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", todoTextClass(done))))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", text)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</span>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n    ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<form")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " method=\"POST\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " action=\"/delete\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n      ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<input")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " type=\"hidden\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " name=\"id\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " value=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", strconv.Itoa(id))))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " />")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n      ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<button")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " type=\"submit\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"px-2.5 py-1 border border-red-500 text-red-500 rounded hover:bg-red-50 text-sm\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "&#x2716;")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</button>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n    ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</form>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n  ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</li>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n")
		if err != nil {
			return err
		}
		return nil
	})
}

func TodoList(todos []Todo) gox.Component {
	return gox.ComponentFunc(func(w io.Writer) error {
		var err error
		_ = err
		_, err = io.WriteString(w, "\n  ")
		if err != nil {
			return err
		}
		err = Layout("Gox Todo App", gox.ComponentFunc(func(w io.Writer) error {
			var err error
			_ = err
			_, err = io.WriteString(w, "\n    ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<h1")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"text-2xl font-bold text-gray-800 mb-5\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "Todo App")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</h1>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n    ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<form")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"flex gap-2 mb-6\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " method=\"POST\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " action=\"/add\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n      ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<input")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " type=\"text\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " name=\"text\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " placeholder=\"What needs to be done?\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " required")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"flex-1 px-3 py-2 border border-gray-300 rounded-md text-base focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " />")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n      ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<button")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " type=\"submit\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700 text-base font-medium cursor-pointer\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "Add")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</button>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n    ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</form>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n    ")
			if err != nil {
				return err
			}
			if len(todos) == 0 {
				_, err = io.WriteString(w, "\n      ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "<p")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " class=\"text-center text-gray-400 py-10\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, ">")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "No todos yet. Add one above!")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "</p>")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n    ")
				if err != nil {
					return err
				}
			} else {
				_, err = io.WriteString(w, "\n      ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "<ul")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " class=\"space-y-2\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, ">")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n        ")
				if err != nil {
					return err
				}
				for _, todo := range todos {
					_, err = io.WriteString(w, "\n          ")
					if err != nil {
						return err
					}
					err = TodoItem(todo.ID, todo.Text, todo.Done).Render(w)
					if err != nil {
						return err
					}
					_, err = io.WriteString(w, "\n        ")
					if err != nil {
						return err
					}
				}
				_, err = io.WriteString(w, "\n      ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "</ul>")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n    ")
				if err != nil {
					return err
				}
			}
			_, err = io.WriteString(w, "\n  ")
			if err != nil {
				return err
			}
			return nil
		})).Render(w)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n")
		if err != nil {
			return err
		}
		return nil
	})
}
