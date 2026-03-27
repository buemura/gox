package views

import (
	"fmt"
	"github.com/buemura/gox"
	"html"
	"io"
)

func Avatar(username string, size string) gox.Component {
	return gox.ComponentFunc(func(w io.Writer) error {
		var err error
		_ = err
		_, err = io.WriteString(w, "\n  ")
		if err != nil {
			return err
		}
		switch size {
		case "lg":
			_, err = io.WriteString(w, "\n      ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "width:80px;height:80px;border-radius:50%;display:flex;align-items:center;justify-content:center;font-weight:700;font-size:32px;color:#fff;background:"+AvatarColor(username))))
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
			_, err = io.WriteString(w, "\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", AvatarInitial(username))))
			if err != nil {
				return err
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
		case "md":
			_, err = io.WriteString(w, "\n      ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "width:48px;height:48px;border-radius:50%;display:flex;align-items:center;justify-content:center;font-weight:700;font-size:18px;color:#fff;flex-shrink:0;background:"+AvatarColor(username))))
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
			_, err = io.WriteString(w, "\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", AvatarInitial(username))))
			if err != nil {
				return err
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
		default:
			_, err = io.WriteString(w, "\n      ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "width:40px;height:40px;border-radius:50%;display:flex;align-items:center;justify-content:center;font-weight:600;font-size:15px;color:#fff;flex-shrink:0;background:"+AvatarColor(username))))
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
			_, err = io.WriteString(w, "\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", AvatarInitial(username))))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n      ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</div>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n  ")
			if err != nil {
				return err
			}
		}
		_, err = io.WriteString(w, "\n")
		if err != nil {
			return err
		}
		return nil
	})
}

func EmptyState(message string) gox.Component {
	return gox.ComponentFunc(func(w io.Writer) error {
		var err error
		_ = err
		_, err = io.WriteString(w, "\n  ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<div")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"flex flex-col items-center justify-center py-16 px-4\"")
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
		_, err = io.WriteString(w, "<svg")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"w-12 h-12 mb-4\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " style=\"color:#2f3336\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " fill=\"none\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " viewBox=\"0 0 24 24\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " stroke=\"currentColor\"")
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
		_, err = io.WriteString(w, "<path")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " stroke-linecap=\"round\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " stroke-linejoin=\"round\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " stroke-width=\"1.5\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " d=\"M7 8h10M7 12h4m1 8l-4-4H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-3l-4 4z\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " />")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n    ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</svg>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n    ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<p")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " style=\"color:#71767b\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"text-base\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", message)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</p>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n  ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
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

func UserCard(user User) gox.Component {
	return gox.ComponentFunc(func(w io.Writer) error {
		var err error
		_ = err
		_, err = io.WriteString(w, "\n  ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<a")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " href=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "/user/"+user.Username)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"flex items-center gap-3 px-4 py-3 transition-colors\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " style=\"text-decoration:none\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " onmouseover=\"this.style.backgroundColor='#16181c'\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " onmouseout=\"this.style.backgroundColor='transparent'\"")
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
		err = Avatar(user.Username, "sm").Render(w)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n    ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<div")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"flex-1 min-w-0\"")
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
		_, err = io.WriteString(w, " class=\"font-bold text-sm truncate\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " style=\"color:#e7e9ea\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", user.DisplayName)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
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
		_, err = io.WriteString(w, " class=\"text-sm truncate\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " style=\"color:#71767b\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "@")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", user.Username)))
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
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n    ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<div")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"px-4 py-1.5 rounded-full text-sm font-bold\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " style=\"background:#e7e9ea;color:#0f1419\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "Follow")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n  ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</a>")
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
