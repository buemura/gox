package views

import (
	"fmt"
	"github.com/buemura/gox"
	"html"
	"io"
)

func LoginPage(errorMsg string) gox.Component {
	return gox.ComponentFunc(func(w io.Writer) error {
		var err error
		_ = err
		_, err = io.WriteString(w, "\n  ")
		if err != nil {
			return err
		}
		err = AuthLayout("Log in / Gox Social", gox.ComponentFunc(func(w io.Writer) error {
			var err error
			_ = err
			_, err = io.WriteString(w, "\n    ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"w-full max-w-[364px] px-8\"")
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
			_, err = io.WriteString(w, " class=\"flex justify-center mb-8\"")
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
			_, err = io.WriteString(w, "<svg")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " viewBox=\"0 0 24 24\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"w-9 h-9\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " fill=\"currentColor\"")
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
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<path")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " d=\"M18.244 2.25h3.308l-7.227 8.26 8.502 11.24H16.17l-5.214-6.817L4.99 21.75H1.68l7.73-8.835L1.254 2.25H8.08l4.713 6.231zm-1.161 17.52h1.833L7.084 4.126H5.117z\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " />")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</svg>")
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
			_, err = io.WriteString(w, "\n\n      ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<h1")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"text-3xl font-bold mb-8\"")
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
			_, err = io.WriteString(w, "Sign in")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</h1>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n\n      ")
			if err != nil {
				return err
			}
			if errorMsg != "" {
				_, err = io.WriteString(w, "\n        ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "<div")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " class=\"mb-4 px-4 py-3 rounded-xl text-sm\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " style=\"background:rgba(244,33,46,0.1);border:1px solid rgba(244,33,46,0.2);color:#f4212e\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, ">")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n          ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", errorMsg)))
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n        ")
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
			}
			_, err = io.WriteString(w, "\n\n      ")
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
			_, err = io.WriteString(w, " action=\"/login\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"flex flex-col gap-5\"")
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
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"relative\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          ")
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
			_, err = io.WriteString(w, " name=\"username\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " required")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " placeholder=\" \"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " id=\"login-user\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"peer w-full px-3 pt-6 pb-2 rounded-md text-[17px] bg-transparent\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"border:1px solid #2f3336;color:#e7e9ea;caret-color:#1d9bf0\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " />")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<label")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " for=\"login-user\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"absolute left-3 top-1/2 -translate-y-1/2 text-[17px] transition-all pointer-events-none peer-focus:top-4 peer-focus:text-xs peer-[:not(:placeholder-shown)]:top-4 peer-[:not(:placeholder-shown)]:text-xs\"")
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
			_, err = io.WriteString(w, "Username")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</label>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</div>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"relative\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<input")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " type=\"password\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " name=\"password\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " required")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " placeholder=\" \"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " id=\"login-pass\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"peer w-full px-3 pt-6 pb-2 rounded-md text-[17px] bg-transparent\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"border:1px solid #2f3336;color:#e7e9ea;caret-color:#1d9bf0\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " />")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<label")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " for=\"login-pass\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"absolute left-3 top-1/2 -translate-y-1/2 text-[17px] transition-all pointer-events-none peer-focus:top-4 peer-focus:text-xs peer-[:not(:placeholder-shown)]:top-4 peer-[:not(:placeholder-shown)]:text-xs\"")
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
			_, err = io.WriteString(w, "Password")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</label>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</div>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n\n        ")
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
			_, err = io.WriteString(w, " class=\"w-full py-3 rounded-full font-bold text-[17px] transition-colors\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"background:#e7e9ea;color:#0f1419\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " onmouseover=\"this.style.background='#d7dbdc'\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " onmouseout=\"this.style.background='#e7e9ea'\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          Sign in\n        ")
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
			_, err = io.WriteString(w, "</form>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n\n      ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<p")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"mt-10 text-base\"")
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
			_, err = io.WriteString(w, "\n        Don't have an account?\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<a")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " href=\"/register\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"hover:underline\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"color:#1d9bf0\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " Sign up")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</a>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n      ")
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
			_, err = io.WriteString(w, "</div>")
			if err != nil {
				return err
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

func RegisterPage(errorMsg string) gox.Component {
	return gox.ComponentFunc(func(w io.Writer) error {
		var err error
		_ = err
		_, err = io.WriteString(w, "\n  ")
		if err != nil {
			return err
		}
		err = AuthLayout("Create account / Gox Social", gox.ComponentFunc(func(w io.Writer) error {
			var err error
			_ = err
			_, err = io.WriteString(w, "\n    ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"w-full max-w-[364px] px-8\"")
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
			_, err = io.WriteString(w, " class=\"flex justify-center mb-8\"")
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
			_, err = io.WriteString(w, "<svg")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " viewBox=\"0 0 24 24\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"w-9 h-9\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " fill=\"currentColor\"")
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
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<path")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " d=\"M18.244 2.25h3.308l-7.227 8.26 8.502 11.24H16.17l-5.214-6.817L4.99 21.75H1.68l7.73-8.835L1.254 2.25H8.08l4.713 6.231zm-1.161 17.52h1.833L7.084 4.126H5.117z\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " />")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</svg>")
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
			_, err = io.WriteString(w, "\n\n      ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<h1")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"text-3xl font-bold mb-8\"")
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
			_, err = io.WriteString(w, "Create your account")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</h1>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n\n      ")
			if err != nil {
				return err
			}
			if errorMsg != "" {
				_, err = io.WriteString(w, "\n        ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "<div")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " class=\"mb-4 px-4 py-3 rounded-xl text-sm\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " style=\"background:rgba(244,33,46,0.1);border:1px solid rgba(244,33,46,0.2);color:#f4212e\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, ">")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n          ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", errorMsg)))
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n        ")
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
			}
			_, err = io.WriteString(w, "\n\n      ")
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
			_, err = io.WriteString(w, " action=\"/register\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"flex flex-col gap-5\"")
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
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"relative\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          ")
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
			_, err = io.WriteString(w, " name=\"display_name\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " required")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " placeholder=\" \"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " id=\"reg-name\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"peer w-full px-3 pt-6 pb-2 rounded-md text-[17px] bg-transparent\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"border:1px solid #2f3336;color:#e7e9ea;caret-color:#1d9bf0\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " />")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<label")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " for=\"reg-name\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"absolute left-3 top-1/2 -translate-y-1/2 text-[17px] transition-all pointer-events-none peer-focus:top-4 peer-focus:text-xs peer-[:not(:placeholder-shown)]:top-4 peer-[:not(:placeholder-shown)]:text-xs\"")
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
			_, err = io.WriteString(w, "Display name")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</label>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</div>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"relative\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          ")
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
			_, err = io.WriteString(w, " name=\"username\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " required")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " placeholder=\" \"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " id=\"reg-user\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"peer w-full px-3 pt-6 pb-2 rounded-md text-[17px] bg-transparent\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"border:1px solid #2f3336;color:#e7e9ea;caret-color:#1d9bf0\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " />")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<label")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " for=\"reg-user\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"absolute left-3 top-1/2 -translate-y-1/2 text-[17px] transition-all pointer-events-none peer-focus:top-4 peer-focus:text-xs peer-[:not(:placeholder-shown)]:top-4 peer-[:not(:placeholder-shown)]:text-xs\"")
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
			_, err = io.WriteString(w, "Username")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</label>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</div>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"relative\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<input")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " type=\"email\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " name=\"email\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " required")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " placeholder=\" \"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " id=\"reg-email\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"peer w-full px-3 pt-6 pb-2 rounded-md text-[17px] bg-transparent\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"border:1px solid #2f3336;color:#e7e9ea;caret-color:#1d9bf0\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " />")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<label")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " for=\"reg-email\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"absolute left-3 top-1/2 -translate-y-1/2 text-[17px] transition-all pointer-events-none peer-focus:top-4 peer-focus:text-xs peer-[:not(:placeholder-shown)]:top-4 peer-[:not(:placeholder-shown)]:text-xs\"")
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
			_, err = io.WriteString(w, "Email")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</label>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</div>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"relative\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<input")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " type=\"password\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " name=\"password\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " required")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " placeholder=\" \"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " id=\"reg-pass\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"peer w-full px-3 pt-6 pb-2 rounded-md text-[17px] bg-transparent\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"border:1px solid #2f3336;color:#e7e9ea;caret-color:#1d9bf0\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " />")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<label")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " for=\"reg-pass\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"absolute left-3 top-1/2 -translate-y-1/2 text-[17px] transition-all pointer-events-none peer-focus:top-4 peer-focus:text-xs peer-[:not(:placeholder-shown)]:top-4 peer-[:not(:placeholder-shown)]:text-xs\"")
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
			_, err = io.WriteString(w, "Password")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</label>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</div>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n\n        ")
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
			_, err = io.WriteString(w, " class=\"w-full py-3 rounded-full font-bold text-[17px] transition-colors\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"background:#e7e9ea;color:#0f1419\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " onmouseover=\"this.style.background='#d7dbdc'\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " onmouseout=\"this.style.background='#e7e9ea'\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          Create account\n        ")
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
			_, err = io.WriteString(w, "</form>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n\n      ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<p")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"mt-10 text-base\"")
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
			_, err = io.WriteString(w, "\n        Already have an account?\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<a")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " href=\"/login\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"hover:underline\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"color:#1d9bf0\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " Sign in")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</a>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n      ")
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
			_, err = io.WriteString(w, "</div>")
			if err != nil {
				return err
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
