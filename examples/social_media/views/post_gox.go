package views

import (
	"fmt"
	"github.com/buemura/gox"
	"html"
	"io"
	"strconv"
)

func PostCard(post PostData, currentUser User) gox.Component {
	return gox.ComponentFunc(func(w io.Writer) error {
		var err error
		_ = err
		_, err = io.WriteString(w, "\n  ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<article")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"flex gap-3 px-4 py-3 transition-colors cursor-pointer\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " style=\"border-bottom:1px solid #2f3336\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " onmouseover=\"this.style.backgroundColor='rgba(231,233,234,0.03)'\"")
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
		_, err = io.WriteString(w, "<a")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " href=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "/user/"+post.Username)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " onclick=\"event.stopPropagation()\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " style=\"text-decoration:none;flex-shrink:0\"")
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
		err = Avatar(post.Username, "md").Render(w)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n    ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</a>")
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
		if IsReply(post) {
			_, err = io.WriteString(w, "\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"text-[13px] mb-0.5\"")
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
			_, err = io.WriteString(w, "\n          Replying to ")
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
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "/user/"+post.ReplyToUser)))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\"")
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
			_, err = io.WriteString(w, " onclick=\"event.stopPropagation()\"")
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
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", post.ReplyToUser)))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</a>")
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
		_, err = io.WriteString(w, "\n      ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<div")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"flex items-center gap-1 mb-0.5\"")
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
		_, err = io.WriteString(w, "<a")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " href=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "/user/"+post.Username)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"font-bold text-[15px] hover:underline truncate\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " style=\"color:#e7e9ea;text-decoration:none\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " onclick=\"event.stopPropagation()\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", post.DisplayName)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</a>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n        ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<span")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"text-[15px] truncate\"")
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
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", post.Username)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</span>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n        ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<span")
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
		_, err = io.WriteString(w, "&#xB7;")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</span>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n        ")
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
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "/post/"+strconv.Itoa(post.ID))))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"text-[15px] hover:underline flex-shrink-0\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " style=\"color:#71767b;text-decoration:none\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " onclick=\"event.stopPropagation()\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", post.TimeAgo)))
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
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n      ")
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
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "/post/"+strconv.Itoa(post.ID))))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " style=\"text-decoration:none;color:inherit\"")
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
		_, err = io.WriteString(w, " class=\"text-[15px] leading-5 post-content\"")
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
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", post.Content)))
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
		_, err = io.WriteString(w, "</a>")
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
		_, err = io.WriteString(w, " class=\"flex items-center gap-0 mt-3 -ml-2 max-w-[425px] justify-between\"")
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
		_, err = io.WriteString(w, "<a")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " href=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "/post/"+strconv.Itoa(post.ID))))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"flex items-center gap-1 group\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " style=\"text-decoration:none\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " onclick=\"event.stopPropagation()\"")
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
		_, err = io.WriteString(w, "<div")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"p-2 rounded-full transition-colors group-hover:bg-blue-500/10\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n            ")
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
		_, err = io.WriteString(w, " class=\"w-[18px] h-[18px] transition-colors group-hover:fill-blue-400\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " fill=\"#71767b\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n              ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<path")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " d=\"M1.751 10c0-4.42 3.584-8 8.005-8h4.366c4.49 0 8.129 3.64 8.129 8.13 0 2.25-.893 4.34-2.457 5.86l-7.17 6.97c-.77.75-2.03.74-2.79-.02L2.929 16.1C2.12 15.26 1.751 14.09 1.751 12.85V10zm8.005-6c-3.317 0-6.005 2.69-6.005 6v2.85c0 .828.26 1.6.73 2.2l6.87 6.61 7.13-6.93c1.17-1.14 1.77-2.71 1.77-4.39C20.251 6.64 17.621 4 14.122 4H9.756z\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " />")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n            ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</svg>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n          ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n          ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<span")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"text-[13px] transition-colors group-hover:text-blue-400\"")
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
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", FormatCount(post.CommentCount))))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</span>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n        ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</a>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n\n        ")
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
		_, err = io.WriteString(w, " action=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", fmt.Sprintf("/post/%d/reply", post.ID))))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"inline\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " onclick=\"event.stopPropagation()\"")
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
		_, err = io.WriteString(w, " type=\"hidden\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " name=\"content\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " value=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", fmt.Sprintf("RT @%s", post.Username))))
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
		_, err = io.WriteString(w, "\n          ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<button")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " type=\"button\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"flex items-center gap-1 group\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " onclick=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", fmt.Sprintf("window.location.href='/post/%d'", post.ID))))
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
		_, err = io.WriteString(w, "\n            ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<div")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"p-2 rounded-full transition-colors group-hover:bg-green-500/10\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n              ")
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
		_, err = io.WriteString(w, " class=\"w-[18px] h-[18px] transition-colors group-hover:fill-green-400\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " fill=\"#71767b\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n                ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<path")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " d=\"M4.5 3.88l4.432 4.14-1.364 1.46L5.5 7.55V16c0 1.1.896 2 2 2H13v2H7.5c-2.209 0-4-1.79-4-4V7.55L1.432 9.48.068 8.02 4.5 3.88zM16.5 6H11V4h5.5c2.209 0 4 1.79 4 4v8.45l2.068-1.93 1.364 1.46-4.432 4.14-4.432-4.14 1.364-1.46 2.068 1.93V8c0-1.1-.896-2-2-2z\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " />")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n              ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</svg>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n            ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n            ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<span")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"text-[13px] transition-colors group-hover:text-green-400\"")
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
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", FormatCount(post.ReplyCount))))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</span>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n          ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</button>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n        ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</form>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n\n        ")
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
		_, err = io.WriteString(w, " action=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", fmt.Sprintf("/post/%d/like", post.ID))))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"inline\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " onclick=\"event.stopPropagation()\"")
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
		_, err = io.WriteString(w, "<button")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " type=\"submit\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"flex items-center gap-1 group\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n            ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<div")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"p-2 rounded-full transition-colors group-hover:bg-pink-500/10\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n              ")
		if err != nil {
			return err
		}
		if post.Liked {
			_, err = io.WriteString(w, "\n                ")
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
			_, err = io.WriteString(w, " class=\"w-[18px] h-[18px]\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " fill=\"#f91880\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n                  ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<path")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " d=\"M20.884 13.19c-1.351 2.48-4.001 5.12-8.379 7.67l-.503.3-.504-.3c-4.379-2.55-7.029-5.19-8.382-7.67-1.36-2.5-1.45-4.55-.782-6.14.602-1.43 1.743-2.51 2.992-3.11 1.252-.6 2.641-.76 3.88-.34.793.27 1.462.72 1.983 1.24l.248.24c.07-.08.146-.15.228-.22.521-.52 1.189-.97 1.983-1.26 1.239-.42 2.628-.26 3.88.34 1.249.6 2.39 1.68 2.992 3.11.668 1.59.578 3.64-.782 6.14z\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " />")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n                ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</svg>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n              ")
			if err != nil {
				return err
			}
		} else {
			_, err = io.WriteString(w, "\n                ")
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
			_, err = io.WriteString(w, " class=\"w-[18px] h-[18px] transition-colors group-hover:fill-pink-500\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " fill=\"#71767b\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n                  ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<path")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " d=\"M16.697 5.5c-1.222-.06-2.679.51-3.89 2.16l-.805 1.09-.806-1.09C9.984 6.01 8.526 5.44 7.304 5.5c-1.243.07-2.349.78-2.91 1.91-.552 1.12-.633 2.78.479 4.82 1.074 1.97 3.257 4.27 7.129 6.61 3.87-2.34 6.052-4.64 7.126-6.61 1.111-2.04 1.03-3.7.477-4.82-.56-1.13-1.666-1.84-2.908-1.91zm4.187 7.69c-1.351 2.48-4.001 5.12-8.379 7.67l-.503.3-.504-.3c-4.379-2.55-7.029-5.19-8.382-7.67-1.36-2.5-1.45-4.55-.782-6.14.602-1.43 1.743-2.51 2.992-3.11 1.252-.6 2.641-.76 3.88-.34.793.27 1.462.72 1.983 1.24.521-.52 1.189-.97 1.982-1.24 1.24-.42 2.629-.26 3.881.34 1.25.6 2.39 1.68 2.993 3.11.667 1.59.577 3.64-.783 6.14z\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " />")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n                ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</svg>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n              ")
			if err != nil {
				return err
			}
		}
		_, err = io.WriteString(w, "\n            ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</div>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n            ")
		if err != nil {
			return err
		}
		if post.Liked {
			_, err = io.WriteString(w, "\n              ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<span")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"text-[13px]\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"color:#f91880\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", FormatCount(post.LikeCount))))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</span>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n            ")
			if err != nil {
				return err
			}
		} else {
			_, err = io.WriteString(w, "\n              ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<span")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"text-[13px] transition-colors group-hover:text-pink-500\"")
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
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", FormatCount(post.LikeCount))))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</span>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n            ")
			if err != nil {
				return err
			}
		}
		_, err = io.WriteString(w, "\n          ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</button>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n        ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</form>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n\n        ")
		if err != nil {
			return err
		}
		if post.UserID == currentUser.ID {
			_, err = io.WriteString(w, "\n          ")
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
			_, err = io.WriteString(w, " action=\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", fmt.Sprintf("/post/%d/delete", post.ID))))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"inline\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " onclick=\"event.stopPropagation()\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n            ")
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
			_, err = io.WriteString(w, " class=\"flex items-center group\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " onclick=\"return confirm('Delete this post?')\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n              ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"p-2 rounded-full transition-colors group-hover:bg-red-500/10\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n                ")
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
			_, err = io.WriteString(w, " class=\"w-[18px] h-[18px] transition-colors group-hover:fill-red-500\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " fill=\"#71767b\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n                  ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<path")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " d=\"M16 6V4.5C16 3.12 14.88 2 13.5 2h-3C9.11 2 8 3.12 8 4.5V6H3v2h1.06l.81 11.21C4.98 20.78 6.28 22 7.86 22h8.27c1.58 0 2.88-1.22 3-2.79L19.93 8H21V6h-5zm-6-1.5c0-.28.22-.5.5-.5h3c.27 0 .5.22.5.5V6h-4V4.5zm7.13 14.57c-.04.52-.47.93-1 .93H7.86c-.53 0-.96-.41-1-.93L6.07 8h11.85l-.79 11.07zM9 17v-6h2v6H9zm4 0v-6h2v6h-2z\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " />")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n                ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</svg>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n              ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</div>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n            ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</button>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</form>")
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
		_, err = io.WriteString(w, "\n  ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</article>")
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

func PostDetailPage(data PostDetailData) gox.Component {
	return gox.ComponentFunc(func(w io.Writer) error {
		var err error
		_ = err
		_, err = io.WriteString(w, "\n  ")
		if err != nil {
			return err
		}
		err = Layout("Post / Gox Social", data.CurrentUser, gox.ComponentFunc(func(w io.Writer) error {
			var err error
			_ = err
			_, err = io.WriteString(w, "\n    ")
			if err != nil {
				return err
			}
			err = PageHeaderWithBack("Post", "/").Render(w)
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n\n    ")
			if err != nil {
				return err
			}
			for _, parent := range data.Thread {
				_, err = io.WriteString(w, "\n      ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "<article")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " class=\"flex gap-3 px-4 pt-3 relative\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " style=\"border-bottom:none\"")
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
				_, err = io.WriteString(w, " class=\"flex flex-col items-center\"")
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
				err = Avatar(parent.Username, "md").Render(w)
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n          ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "<div")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " class=\"w-0.5 flex-1 mt-1\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " style=\"background:#2f3336\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, ">")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "</div>")
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
				_, err = io.WriteString(w, "\n        ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "<div")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " class=\"flex-1 min-w-0 pb-3\"")
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
				_, err = io.WriteString(w, "<div")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " class=\"flex items-center gap-1 mb-0.5\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, ">")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n            ")
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
				_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "/user/"+parent.Username)))
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " class=\"font-bold text-[15px] hover:underline\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " style=\"color:#e7e9ea;text-decoration:none\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, ">")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", parent.DisplayName)))
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "</a>")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n            ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "<span")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " class=\"text-[15px]\"")
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
				_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", parent.Username)))
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "</span>")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n            ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "<span")
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
				_, err = io.WriteString(w, "&#xB7;")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "</span>")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n            ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "<span")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " class=\"text-[15px]\"")
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
				_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", parent.TimeAgo)))
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "</span>")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n          ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "</div>")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n          ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "<div")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " class=\"text-[15px] leading-5 post-content\"")
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
				_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", parent.Content)))
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "</div>")
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
				_, err = io.WriteString(w, "</article>")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n    ")
				if err != nil {
					return err
				}
			}
			_, err = io.WriteString(w, "\n\n    ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<article")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"px-4 py-3\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"border-bottom:1px solid #2f3336\"")
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
			_, err = io.WriteString(w, " class=\"flex items-center gap-3 mb-3\"")
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
			_, err = io.WriteString(w, "<a")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " href=\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "/user/"+data.Post.Username)))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"text-decoration:none\"")
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
			err = Avatar(data.Post.Username, "md").Render(w)
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</a>")
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
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          ")
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
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "/user/"+data.Post.Username)))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"font-bold text-[17px] hover:underline block leading-5\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"color:#e7e9ea;text-decoration:none\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", data.Post.DisplayName)))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</a>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<span")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"text-[15px]\"")
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
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", data.Post.Username)))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</span>")
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
			_, err = io.WriteString(w, " class=\"text-[17px] leading-6 post-content mb-3\"")
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
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", data.Post.Content)))
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
			_, err = io.WriteString(w, " class=\"py-3 text-[15px]\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"color:#71767b;border-top:1px solid #2f3336\"")
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
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", data.Post.TimeAgo)))
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
			_, err = io.WriteString(w, "\n      ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"flex gap-6 py-3\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"border-top:1px solid #2f3336\"")
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
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<span")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"font-bold text-[15px]\"")
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
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", FormatCount(data.Post.ReplyCount))))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</span>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<span")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"text-[15px] ml-1\"")
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
			_, err = io.WriteString(w, "Replies")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</span>")
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
			_, err = io.WriteString(w, "\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<div")
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
			_, err = io.WriteString(w, "<span")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"font-bold text-[15px]\"")
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
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", FormatCount(data.Post.LikeCount))))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</span>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<span")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"text-[15px] ml-1\"")
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
			_, err = io.WriteString(w, "Likes")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</span>")
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
			_, err = io.WriteString(w, "\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<div")
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
			_, err = io.WriteString(w, "<span")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"font-bold text-[15px]\"")
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
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", FormatCount(data.Post.CommentCount))))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</span>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<span")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"text-[15px] ml-1\"")
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
			_, err = io.WriteString(w, "Comments")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</span>")
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
			_, err = io.WriteString(w, " class=\"flex items-center justify-around py-1\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"border-top:1px solid #2f3336\"")
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
			_, err = io.WriteString(w, "<form")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " method=\"POST\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " action=\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", fmt.Sprintf("/post/%d/like", data.Post.ID))))
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
			_, err = io.WriteString(w, "\n          ")
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
			_, err = io.WriteString(w, " class=\"p-2 rounded-full transition-colors hover:bg-pink-500/10 group\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n            ")
			if err != nil {
				return err
			}
			if data.Post.Liked {
				_, err = io.WriteString(w, "\n              ")
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
				_, err = io.WriteString(w, " class=\"w-[22px] h-[22px]\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " fill=\"#f91880\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, ">")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n                ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "<path")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " d=\"M20.884 13.19c-1.351 2.48-4.001 5.12-8.379 7.67l-.503.3-.504-.3c-4.379-2.55-7.029-5.19-8.382-7.67-1.36-2.5-1.45-4.55-.782-6.14.602-1.43 1.743-2.51 2.992-3.11 1.252-.6 2.641-.76 3.88-.34.793.27 1.462.72 1.983 1.24l.248.24c.07-.08.146-.15.228-.22.521-.52 1.189-.97 1.983-1.26 1.239-.42 2.628-.26 3.88.34 1.249.6 2.39 1.68 2.992 3.11.668 1.59.578 3.64-.782 6.14z\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " />")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n              ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "</svg>")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n            ")
				if err != nil {
					return err
				}
			} else {
				_, err = io.WriteString(w, "\n              ")
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
				_, err = io.WriteString(w, " class=\"w-[22px] h-[22px] transition-colors group-hover:fill-pink-500\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " fill=\"#71767b\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, ">")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n                ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "<path")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " d=\"M16.697 5.5c-1.222-.06-2.679.51-3.89 2.16l-.805 1.09-.806-1.09C9.984 6.01 8.526 5.44 7.304 5.5c-1.243.07-2.349.78-2.91 1.91-.552 1.12-.633 2.78.479 4.82 1.074 1.97 3.257 4.27 7.129 6.61 3.87-2.34 6.052-4.64 7.126-6.61 1.111-2.04 1.03-3.7.477-4.82-.56-1.13-1.666-1.84-2.908-1.91zm4.187 7.69c-1.351 2.48-4.001 5.12-8.379 7.67l-.503.3-.504-.3c-4.379-2.55-7.029-5.19-8.382-7.67-1.36-2.5-1.45-4.55-.782-6.14.602-1.43 1.743-2.51 2.992-3.11 1.252-.6 2.641-.76 3.88-.34.793.27 1.462.72 1.983 1.24.521-.52 1.189-.97 1.982-1.24 1.24-.42 2.629-.26 3.881.34 1.25.6 2.39 1.68 2.993 3.11.667 1.59.577 3.64-.783 6.14z\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " />")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n              ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "</svg>")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n            ")
				if err != nil {
					return err
				}
			}
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</button>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n        ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</form>")
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
			_, err = io.WriteString(w, "</article>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n\n    ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"border-bottom:1px solid #2f3336\"")
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
			err = ComposeBox(fmt.Sprintf("/post/%d/reply", data.Post.ID), "Post your reply", "Reply", data.CurrentUser).Render(w)
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
			_, err = io.WriteString(w, "\n\n    ")
			if err != nil {
				return err
			}
			if len(data.Replies) > 0 {
				_, err = io.WriteString(w, "\n      ")
				if err != nil {
					return err
				}
				for _, reply := range data.Replies {
					_, err = io.WriteString(w, "\n        ")
					if err != nil {
						return err
					}
					err = PostCard(reply, data.CurrentUser).Render(w)
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
			}
			_, err = io.WriteString(w, "\n\n    ")
			if err != nil {
				return err
			}
			if len(data.Comments) > 0 {
				_, err = io.WriteString(w, "\n      ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "<div")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " class=\"px-4 py-3\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " style=\"border-bottom:1px solid #2f3336\"")
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
				_, err = io.WriteString(w, "<h3")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " class=\"font-bold text-[17px] mb-3\"")
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
				_, err = io.WriteString(w, "Comments")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "</h3>")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n        ")
				if err != nil {
					return err
				}
				for _, comment := range data.Comments {
					_, err = io.WriteString(w, "\n          ")
					if err != nil {
						return err
					}
					err = CommentItem(comment).Render(w)
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
				_, err = io.WriteString(w, "</div>")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n    ")
				if err != nil {
					return err
				}
			}
			_, err = io.WriteString(w, "\n\n    ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"px-4 py-3\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"border-bottom:1px solid #2f3336\"")
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
			_, err = io.WriteString(w, "<form")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " method=\"POST\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " action=\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", fmt.Sprintf("/post/%d/comment", data.Post.ID))))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"flex gap-3\"")
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
			err = Avatar(data.CurrentUser.Username, "sm").Render(w)
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
			_, err = io.WriteString(w, " class=\"flex-1\"")
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
			_, err = io.WriteString(w, "<textarea")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " name=\"content\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " rows=\"2\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " placeholder=\"Add a comment...\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " required")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"w-full bg-transparent text-[15px] py-2\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"color:#e7e9ea;border:none\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</textarea>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"flex justify-end\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n            ")
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
			_, err = io.WriteString(w, " class=\"px-4 py-1.5 rounded-full font-bold text-sm\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"background:#1d9bf0;color:#fff\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, ">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "Comment")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</button>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</div>")
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
			_, err = io.WriteString(w, "</form>")
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

func CommentItem(comment CommentData) gox.Component {
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
		_, err = io.WriteString(w, " class=\"flex gap-3 py-3\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " style=\"border-bottom:1px solid rgba(47,51,54,0.5)\"")
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
		_, err = io.WriteString(w, "<a")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " href=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "/user/"+comment.Username)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " style=\"text-decoration:none;flex-shrink:0\"")
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
		err = Avatar(comment.Username, "sm").Render(w)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n    ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</a>")
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
		_, err = io.WriteString(w, " class=\"flex items-center gap-1 mb-0.5\"")
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
		_, err = io.WriteString(w, "<a")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " href=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "/user/"+comment.Username)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"font-bold text-[15px] hover:underline\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " style=\"color:#e7e9ea;text-decoration:none\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, ">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", comment.DisplayName)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</a>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n        ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<span")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"text-[15px]\"")
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
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", comment.Username)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</span>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n        ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<span")
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
		_, err = io.WriteString(w, "&#xB7;")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</span>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n        ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<span")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"text-[13px]\"")
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
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", comment.TimeAgo)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</span>")
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
		_, err = io.WriteString(w, "\n      ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<div")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"text-[15px] leading-5\"")
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
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", comment.Content)))
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
