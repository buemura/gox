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
			_, err = io.WriteString(w, "\n      <div style=\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "width:80px;height:80px;border-radius:50%;display:flex;align-items:center;justify-content:center;font-weight:700;font-size:32px;color:#fff;background:"+AvatarColor(username))))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\">\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", AvatarInitial(username))))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n      </div>\n    ")
			if err != nil {
				return err
			}
		case "md":
			_, err = io.WriteString(w, "\n      <div style=\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "width:48px;height:48px;border-radius:50%;display:flex;align-items:center;justify-content:center;font-weight:700;font-size:18px;color:#fff;flex-shrink:0;background:"+AvatarColor(username))))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\">\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", AvatarInitial(username))))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n      </div>\n    ")
			if err != nil {
				return err
			}
		default:
			_, err = io.WriteString(w, "\n      <div style=\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "width:40px;height:40px;border-radius:50%;display:flex;align-items:center;justify-content:center;font-weight:600;font-size:15px;color:#fff;flex-shrink:0;background:"+AvatarColor(username))))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\">\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", AvatarInitial(username))))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n      </div>\n  ")
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
		_, err = io.WriteString(w, "\n  <div class=\"flex flex-col items-center justify-center py-16 px-4\">\n    <svg class=\"w-12 h-12 mb-4\" style=\"color:#2f3336\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\">\n      <path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"1.5\" d=\"M7 8h10M7 12h4m1 8l-4-4H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-3l-4 4z\" />\n    </svg>\n    <p style=\"color:#71767b\" class=\"text-base\">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", message)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</p>\n  </div>\n")
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
		_, err = io.WriteString(w, "\n  <a href=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "/user/"+user.Username)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\" class=\"flex items-center gap-3 px-4 py-3 transition-colors\" style=\"text-decoration:none\" onmouseover=\"this.style.backgroundColor='#16181c'\" onmouseout=\"this.style.backgroundColor='transparent'\">\n    ")
		if err != nil {
			return err
		}
		err = Avatar(user.Username, "sm").Render(w)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n    <div class=\"flex-1 min-w-0\">\n      <div class=\"font-bold text-sm truncate\" style=\"color:#e7e9ea\">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", user.DisplayName)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>\n      <div class=\"text-sm truncate\" style=\"color:#71767b\">@")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", user.Username)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>\n    </div>\n    <div class=\"px-4 py-1.5 rounded-full text-sm font-bold\" style=\"background:#e7e9ea;color:#0f1419\">Follow</div>\n  </a>\n")
		if err != nil {
			return err
		}
		return nil
	})
}
