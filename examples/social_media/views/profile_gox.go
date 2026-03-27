package views

import (
	"fmt"
	"github.com/buemura/gox"
	"html"
	"io"
	"strconv"
)

func ProfilePage(data ProfileData) gox.Component {
	return gox.ComponentFunc(func(w io.Writer) error {
		var err error
		_ = err
		_, err = io.WriteString(w, "\n  ")
		if err != nil {
			return err
		}
		err = Layout(data.ProfileUser.DisplayName+" (@"+data.ProfileUser.Username+") / Gox Social", data.CurrentUser, gox.ComponentFunc(func(w io.Writer) error {
			var err error
			_ = err
			_, err = io.WriteString(w, "\n    ")
			if err != nil {
				return err
			}
			err = PageHeaderWithBack(data.ProfileUser.DisplayName, "/").Render(w)
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
			_, err = io.WriteString(w, " class=\"h-[200px]\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "background:linear-gradient(135deg, "+AvatarColor(data.ProfileUser.Username)+" 0%, #000 100%)")))
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
			_, err = io.WriteString(w, "</div>")
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
			_, err = io.WriteString(w, " class=\"px-4 pb-3\"")
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
			_, err = io.WriteString(w, " class=\"flex justify-between items-start -mt-[42px] mb-3\"")
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
			_, err = io.WriteString(w, " class=\"rounded-full p-1\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"background:#000\"")
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
			err = Avatar(data.ProfileUser.Username, "lg").Render(w)
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
			if data.IsOwnProfile {
				_, err = io.WriteString(w, "\n          ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "<a")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " href=\"#\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " class=\"mt-[52px] px-4 py-1.5 rounded-full font-bold text-[15px] transition-colors\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " style=\"border:1px solid #536471;color:#e7e9ea;text-decoration:none\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " onmouseover=\"this.style.backgroundColor='rgba(239,243,244,0.1)'\"")
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
				_, err = io.WriteString(w, "\n            Edit profile\n          ")
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
			} else {
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
				_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", fmt.Sprintf("/user/%s/follow", data.ProfileUser.Username))))
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " class=\"mt-[52px]\"")
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
				if data.IsFollowing {
					_, err = io.WriteString(w, "\n              ")
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
					_, err = io.WriteString(w, " class=\"px-4 py-1.5 rounded-full font-bold text-[15px] transition-colors\"")
					if err != nil {
						return err
					}
					_, err = io.WriteString(w, " style=\"border:1px solid #536471;color:#e7e9ea\"")
					if err != nil {
						return err
					}
					_, err = io.WriteString(w, " onmouseover=\"this.style.borderColor='#67070f';this.style.color='#f4212e';this.textContent='Unfollow'\"")
					if err != nil {
						return err
					}
					_, err = io.WriteString(w, " onmouseout=\"this.style.borderColor='#536471';this.style.color='#e7e9ea';this.textContent='Following'\"")
					if err != nil {
						return err
					}
					_, err = io.WriteString(w, ">")
					if err != nil {
						return err
					}
					_, err = io.WriteString(w, "\n                Following\n              ")
					if err != nil {
						return err
					}
					_, err = io.WriteString(w, "</button>")
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
					_, err = io.WriteString(w, "<button")
					if err != nil {
						return err
					}
					_, err = io.WriteString(w, " type=\"submit\"")
					if err != nil {
						return err
					}
					_, err = io.WriteString(w, " class=\"px-4 py-1.5 rounded-full font-bold text-[15px] transition-colors\"")
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
					_, err = io.WriteString(w, "\n                Follow\n              ")
					if err != nil {
						return err
					}
					_, err = io.WriteString(w, "</button>")
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
			_, err = io.WriteString(w, "\n\n      ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"mb-3\"")
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
			_, err = io.WriteString(w, " class=\"text-xl font-extrabold leading-6\"")
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
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", data.ProfileUser.DisplayName)))
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
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", data.ProfileUser.Username)))
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
			_, err = io.WriteString(w, "\n\n      ")
			if err != nil {
				return err
			}
			if data.ProfileUser.Bio != "" {
				_, err = io.WriteString(w, "\n        ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "<p")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, " class=\"text-[15px] leading-5 mb-3\"")
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
				_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", data.ProfileUser.Bio)))
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "</p>")
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
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"flex gap-5\"")
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
			_, err = io.WriteString(w, "<span")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"text-sm\"")
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
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<span")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"font-bold\"")
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
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", strconv.Itoa(data.ProfileUser.FollowingCount))))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</span>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " Following\n        ")
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
			_, err = io.WriteString(w, " class=\"text-sm\"")
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
			_, err = io.WriteString(w, "\n          ")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "<span")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"font-bold\"")
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
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", strconv.Itoa(data.ProfileUser.FollowerCount))))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</span>")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " Followers\n        ")
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
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"flex\"")
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
			_, err = io.WriteString(w, " class=\"flex-1 flex justify-center py-4 transition-colors hover:bg-xhover relative\"")
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
			_, err = io.WriteString(w, "Posts")
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
			_, err = io.WriteString(w, "<div")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " class=\"absolute bottom-0 h-1 w-14 rounded-full\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, " style=\"background:#1d9bf0\"")
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
			_, err = io.WriteString(w, "\n\n    ")
			if err != nil {
				return err
			}
			if len(data.Posts) == 0 {
				_, err = io.WriteString(w, "\n      ")
				if err != nil {
					return err
				}
				err = EmptyState("No posts yet.").Render(w)
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
