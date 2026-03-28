package views

import (
	"fmt"
	"html"
	"io"
	"strconv"

	"github.com/buemura/gox"
)

func Layout(title string, children gox.Component) gox.Component {
	return gox.ComponentFunc(func(w io.Writer) error {
		var err error
		_ = err
		_, err = io.WriteString(w, "\n  <!DOCTYPE html>\n  <html lang=\"en\">\n    <head>\n      <meta charset=\"UTF-8\" />\n      <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\" />\n      <title>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", title)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</title>\n      <script src=\"https://cdn.tailwindcss.com\"></script>\n    </head>\n    <body class=\"bg-gray-100 min-h-screen\">\n      <div class=\"max-w-xl mx-auto px-5 py-10\">\n        ")
		if err != nil {
			return err
		}
		if children != nil {
			err = children.Render(w)
			if err != nil {
				return err
			}
		}
		_, err = io.WriteString(w, "\n      </div>\n    </body>\n  </html>\n")
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
		_, err = io.WriteString(w, "\n  <li class=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", todoItemClass(done))))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\">\n    <form method=\"POST\" action=\"/toggle\">\n      <input type=\"hidden\" name=\"id\" value=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", strconv.Itoa(id))))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\" />\n      ")
		if err != nil {
			return err
		}
		if done {
			_, err = io.WriteString(w, "\n        <button type=\"submit\" class=\"text-lg\">&#x2705;</button>\n      ")
			if err != nil {
				return err
			}
		} else {
			_, err = io.WriteString(w, "\n        <button type=\"submit\" class=\"text-lg\">&#x2B1C;</button>\n      ")
			if err != nil {
				return err
			}
		}
		_, err = io.WriteString(w, "\n    </form>\n    <span class=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", todoTextClass(done))))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", text)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</span>\n    <form method=\"POST\" action=\"/delete\">\n      <input type=\"hidden\" name=\"id\" value=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", strconv.Itoa(id))))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\" />\n      <button type=\"submit\" class=\"px-2.5 py-1 border border-red-500 text-red-500 rounded hover:bg-red-50 text-sm\">&#x2716;</button>\n    </form>\n  </li>\n")
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
			_, err = io.WriteString(w, "\n    <h1 class=\"text-2xl font-bold text-gray-800 mb-5\">Todo App</h1>\n    <form class=\"flex gap-2 mb-6\" method=\"POST\" action=\"/add\">\n      <input type=\"text\" name=\"text\" placeholder=\"What needs to be done?\" required class=\"flex-1 px-3 py-2 border border-gray-300 rounded-md text-base focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent\" />\n      <button type=\"submit\" class=\"px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700 text-base font-medium cursor-pointer\">Add</button>\n    </form>\n    ")
			if err != nil {
				return err
			}
			if len(todos) == 0 {
				_, err = io.WriteString(w, "\n      <p class=\"text-center text-gray-400 py-10\">No todos yet. Add one above!</p>\n    ")
				if err != nil {
					return err
				}
			} else {
				_, err = io.WriteString(w, "\n      <ul class=\"space-y-2\">\n        ")
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
				_, err = io.WriteString(w, "\n      </ul>\n    ")
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
