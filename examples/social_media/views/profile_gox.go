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
			_, err = io.WriteString(w, "\n\n    <div class=\"h-[200px]\" style=\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "background:linear-gradient(135deg, "+AvatarColor(data.ProfileUser.Username)+" 0%, #000 100%)")))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\"></div>\n\n    <div class=\"px-4 pb-3\" style=\"border-bottom:1px solid #2f3336\">\n      <div class=\"flex justify-between items-start -mt-[42px] mb-3\">\n        <div class=\"rounded-full p-1\" style=\"background:#000\">\n          ")
			if err != nil {
				return err
			}
			err = Avatar(data.ProfileUser.Username, "lg").Render(w)
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n        </div>\n\n        ")
			if err != nil {
				return err
			}
			if data.IsOwnProfile {
				_, err = io.WriteString(w, "\n          <a href=\"#\" class=\"mt-[52px] px-4 py-1.5 rounded-full font-bold text-[15px] transition-colors\" style=\"border:1px solid #536471;color:#e7e9ea;text-decoration:none\" onmouseover=\"this.style.backgroundColor='rgba(239,243,244,0.1)'\" onmouseout=\"this.style.backgroundColor='transparent'\">\n            Edit profile\n          </a>\n        ")
				if err != nil {
					return err
				}
			} else {
				_, err = io.WriteString(w, "\n          <form method=\"POST\" action=\"")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", fmt.Sprintf("/user/%s/follow", data.ProfileUser.Username))))
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\" class=\"mt-[52px]\">\n            ")
				if err != nil {
					return err
				}
				if data.IsFollowing {
					_, err = io.WriteString(w, "\n              <button type=\"submit\" class=\"px-4 py-1.5 rounded-full font-bold text-[15px] transition-colors\" style=\"border:1px solid #536471;color:#e7e9ea\" onmouseover=\"this.style.borderColor='#67070f';this.style.color='#f4212e';this.textContent='Unfollow'\" onmouseout=\"this.style.borderColor='#536471';this.style.color='#e7e9ea';this.textContent='Following'\">\n                Following\n              </button>\n            ")
					if err != nil {
						return err
					}
				} else {
					_, err = io.WriteString(w, "\n              <button type=\"submit\" class=\"px-4 py-1.5 rounded-full font-bold text-[15px] transition-colors\" style=\"background:#e7e9ea;color:#0f1419\" onmouseover=\"this.style.background='#d7dbdc'\" onmouseout=\"this.style.background='#e7e9ea'\">\n                Follow\n              </button>\n            ")
					if err != nil {
						return err
					}
				}
				_, err = io.WriteString(w, "\n          </form>\n        ")
				if err != nil {
					return err
				}
			}
			_, err = io.WriteString(w, "\n      </div>\n\n      <div class=\"mb-3\">\n        <h2 class=\"text-xl font-extrabold leading-6\" style=\"color:#e7e9ea\">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", data.ProfileUser.DisplayName)))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</h2>\n        <span class=\"text-[15px]\" style=\"color:#71767b\">@")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", data.ProfileUser.Username)))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</span>\n      </div>\n\n      ")
			if err != nil {
				return err
			}
			if data.ProfileUser.Bio != "" {
				_, err = io.WriteString(w, "\n        <p class=\"text-[15px] leading-5 mb-3\" style=\"color:#e7e9ea\">")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", data.ProfileUser.Bio)))
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "</p>\n      ")
				if err != nil {
					return err
				}
			}
			_, err = io.WriteString(w, "\n\n      <div class=\"flex gap-5\">\n        <span class=\"text-sm\" style=\"color:#71767b\">\n          <span class=\"font-bold\" style=\"color:#e7e9ea\">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", strconv.Itoa(data.ProfileUser.FollowingCount))))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</span> Following\n        </span>\n        <span class=\"text-sm\" style=\"color:#71767b\">\n          <span class=\"font-bold\" style=\"color:#e7e9ea\">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", strconv.Itoa(data.ProfileUser.FollowerCount))))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</span> Followers\n        </span>\n      </div>\n    </div>\n\n    <div class=\"flex\" style=\"border-bottom:1px solid #2f3336\">\n      <div class=\"flex-1 flex justify-center py-4 transition-colors hover:bg-xhover relative\">\n        <span class=\"font-bold text-[15px]\" style=\"color:#e7e9ea\">Posts</span>\n        <div class=\"absolute bottom-0 h-1 w-14 rounded-full\" style=\"background:#1d9bf0\"></div>\n      </div>\n    </div>\n\n    ")
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
