package views

import (
	"fmt"
	"github.com/buemura/gox"
	"html"
	"io"
)

func ComposeBox(action string, placeholder string, buttonText string, currentUser User) gox.Component {
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
		_, err = io.WriteString(w, " class=\"px-4 py-3 flex gap-3\"")
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
		_, err = io.WriteString(w, "\n    ")
		if err != nil {
			return err
		}
		err = Avatar(currentUser.Username, "md").Render(w)
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
		_, err = io.WriteString(w, " action=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", action)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"flex-1 flex flex-col\"")
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
		_, err = io.WriteString(w, "<textarea")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " name=\"content\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " rows=\"3\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " placeholder=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", placeholder)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " required")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"w-full bg-transparent text-xl py-3 placeholder-opacity-60\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " style=\"color:#e7e9ea;border:none\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " onfocus=\"this.parentElement.querySelector('button').style.opacity='1'\"")
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
		_, err = io.WriteString(w, "\n      ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<div")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"flex justify-end pt-3\"")
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
		_, err = io.WriteString(w, "<button")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " type=\"submit\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " class=\"px-5 py-2 rounded-full font-bold text-[15px] transition-all\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " style=\"background:#1d9bf0;color:#fff;opacity:0.85\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " onmouseover=\"this.style.background='#1a8cd8'\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " onmouseout=\"this.style.background='#1d9bf0'\"")
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
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", buttonText)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n        ")
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
		_, err = io.WriteString(w, "</div>")
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

func FeedPage(data FeedPageData) gox.Component {
	return gox.ComponentFunc(func(w io.Writer) error {
		var err error
		_ = err
		_, err = io.WriteString(w, "\n  ")
		if err != nil {
			return err
		}
		err = Layout("Home / Gox Social", data.CurrentUser, gox.ComponentFunc(func(w io.Writer) error {
			var err error
			_ = err
			_, err = io.WriteString(w, "\n    ")
			if err != nil {
				return err
			}
			err = PageHeader("Home").Render(w)
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n    ")
			if err != nil {
				return err
			}
			err = ComposeBox("/posts", "What is happening?!", "Post", data.CurrentUser).Render(w)
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n    ")
			if err != nil {
				return err
			}
			if len(data.Posts) == 0 {
				_, err = io.WriteString(w, "\n      ")
				if err != nil {
					return err
				}
				err = EmptyState("No posts yet. Follow some users or create your first post!").Render(w)
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
				for _, post := range data.Posts {
					_, err = io.WriteString(w, "\n        ")
					if err != nil {
						return err
					}
					err = PostCard(post, data.CurrentUser).Render(w)
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

func ExplorePage(data ExploreData) gox.Component {
	return gox.ComponentFunc(func(w io.Writer) error {
		var err error
		_ = err
		_, err = io.WriteString(w, "\n  ")
		if err != nil {
			return err
		}
		err = Layout("Explore / Gox Social", data.CurrentUser, gox.ComponentFunc(func(w io.Writer) error {
			var err error
			_ = err
			_, err = io.WriteString(w, "\n    ")
			if err != nil {
				return err
			}
			err = PageHeader("Explore").Render(w)
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n\n    ")
			if err != nil {
				return err
			}
			if len(data.SuggestedUsers) > 0 {
				_, err = io.WriteString(w, "\n      ")
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
				_, err = io.WriteString(w, "\n        ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "<h2")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " class=\"px-4 pt-3 pb-1 font-extrabold text-xl\"")
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
				_, err = io.WriteString(w, "Who to follow")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "</h2>")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n        ")
				if err != nil {
					return err
				}
				for _, user := range data.SuggestedUsers {
					_, err = io.WriteString(w, "\n          ")
					if err != nil {
						return err
					}
					err = UserCard(user).Render(w)
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
			if len(data.Posts) == 0 {
				_, err = io.WriteString(w, "\n      ")
				if err != nil {
					return err
				}
				err = EmptyState("No posts yet. Be the first to post!").Render(w)
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
				for _, post := range data.Posts {
					_, err = io.WriteString(w, "\n        ")
					if err != nil {
						return err
					}
					err = PostCard(post, data.CurrentUser).Render(w)
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
