package views

import (
	"fmt"
	"github.com/buemura/gox"
	"html"
	"io"
	"strconv"
)

func Layout(title string, currentUser User, children gox.Component) gox.Component {
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
		_, err = io.WriteString(w, "</title>\n      <script src=\"https://cdn.tailwindcss.com\"></script>\n      <link rel=\"preconnect\" href=\"https://fonts.googleapis.com\" />\n      <link rel=\"preconnect\" href=\"https://fonts.gstatic.com\" crossorigin />\n      <link href=\"https://fonts.googleapis.com/css2?family=Chirp:wght@400;500;700&family=Inter:wght@400;500;600;700&display=swap\" rel=\"stylesheet\" />\n      ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, fmt.Sprintf("%v", `<script>tailwind.config={theme:{extend:{colors:{xbg:'#000000',xsurface:'#16181c',xborder:'#2f3336',xtext:'#e7e9ea',xmuted:'#71767b',xaccent:'#1d9bf0',xlike:'#f91880',xrepost:'#00ba7c',xhover:'rgba(231,233,234,0.1)'},fontFamily:{chirp:['"Inter"','system-ui','sans-serif']}}}};</script>`))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n      ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, fmt.Sprintf("%v", `<style>*{font-family:'Inter',system-ui,sans-serif}body{background:#000;color:#e7e9ea;margin:0}a{color:inherit}::selection{background:#1d9bf0;color:#fff}::-webkit-scrollbar{width:4px}::-webkit-scrollbar-track{background:transparent}::-webkit-scrollbar-thumb{background:#2f3336;border-radius:4px}textarea{resize:none}textarea:focus,input:focus{outline:none}.post-content{white-space:pre-wrap;word-break:break-word}</style>`))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n    </head>\n    <body>\n      <div class=\"flex mx-auto min-h-screen\" style=\"max-width:1265px\">\n        <nav class=\"hidden lg:flex flex-col items-end w-[275px] pr-3 sticky top-0 h-screen flex-shrink-0\" style=\"border-right:none\">\n          <div class=\"flex flex-col gap-1 py-2 w-[250px]\">\n            <a href=\"/\" class=\"flex items-center gap-4 px-4 py-3 rounded-full transition-colors hover:bg-xhover\">\n              <svg viewBox=\"0 0 24 24\" class=\"w-7 h-7\" fill=\"currentColor\">\n                <path d=\"M18.244 2.25h3.308l-7.227 8.26 8.502 11.24H16.17l-5.214-6.817L4.99 21.75H1.68l7.73-8.835L1.254 2.25H8.08l4.713 6.231zm-1.161 17.52h1.833L7.084 4.126H5.117z\" />\n              </svg>\n            </a>\n\n            <a href=\"/\" class=\"flex items-center gap-4 px-4 py-3 rounded-full transition-colors hover:bg-xhover\">\n              <svg viewBox=\"0 0 24 24\" class=\"w-[26px] h-[26px]\" fill=\"currentColor\">\n                <path d=\"M21.591 7.146L12.52 1.157c-.316-.21-.724-.21-1.04 0l-9.071 5.99c-.26.173-.409.456-.409.757v13.183c0 .502.418.913.929.913h6.638c.511 0 .929-.41.929-.913v-7.075h3.008v7.075c0 .502.418.913.929.913h6.638c.511 0 .929-.41.929-.913V7.904c0-.301-.158-.584-.408-.758z\" />\n              </svg>\n              <span class=\"text-xl font-bold\">Home</span>\n            </a>\n\n            <a href=\"/explore\" class=\"flex items-center gap-4 px-4 py-3 rounded-full transition-colors hover:bg-xhover\">\n              <svg viewBox=\"0 0 24 24\" class=\"w-[26px] h-[26px]\" fill=\"currentColor\">\n                <path d=\"M10.25 3.75c-3.59 0-6.5 2.91-6.5 6.5s2.91 6.5 6.5 6.5c1.795 0 3.419-.726 4.596-1.904 1.178-1.177 1.904-2.801 1.904-4.596 0-3.59-2.91-6.5-6.5-6.5zm-8.5 6.5c0-4.694 3.806-8.5 8.5-8.5s8.5 3.806 8.5 8.5c0 1.986-.682 3.815-1.824 5.262l4.781 4.781-1.414 1.414-4.781-4.781c-1.447 1.142-3.276 1.824-5.262 1.824-4.694 0-8.5-3.806-8.5-8.5z\" />\n              </svg>\n              <span class=\"text-xl\">Explore</span>\n            </a>\n\n            ")
		if err != nil {
			return err
		}
		if currentUser.ID > 0 {
			_, err = io.WriteString(w, "\n              <a href=\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "/user/"+currentUser.Username)))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\" class=\"flex items-center gap-4 px-4 py-3 rounded-full transition-colors hover:bg-xhover\">\n                <svg viewBox=\"0 0 24 24\" class=\"w-[26px] h-[26px]\" fill=\"currentColor\">\n                  <path d=\"M5.651 19h12.698c-.337-1.8-1.023-3.21-1.945-4.19C15.318 13.65 13.838 13 12 13s-3.317.65-4.404 1.81c-.922.98-1.608 2.39-1.945 4.19zm.486-5.56C7.627 11.85 9.648 11 12 11s4.373.85 5.863 2.44c1.477 1.58 2.366 3.8 2.632 6.46l.11 1.1H3.395l.11-1.1c.266-2.66 1.155-4.88 2.632-6.46zM12 4c-1.105 0-2 .9-2 2s.895 2 2 2 2-.9 2-2-.895-2-2-2zM8 6c0-2.21 1.791-4 4-4s4 1.79 4 4-1.791 4-4 4-4-1.79-4-4z\" />\n                </svg>\n                <span class=\"text-xl\">Profile</span>\n              </a>\n\n              <form method=\"POST\" action=\"/logout\" class=\"mt-2\">\n                <button type=\"submit\" class=\"flex items-center gap-4 px-4 py-3 rounded-full transition-colors hover:bg-xhover w-full text-left\">\n                  <svg viewBox=\"0 0 24 24\" class=\"w-[26px] h-[26px]\" fill=\"currentColor\">\n                    <path d=\"M5 2h14c1.1 0 2 .9 2 2v16c0 1.1-.9 2-2 2H5c-1.1 0-2-.9-2-2v-4h2v4h14V4H5v4H3V4c0-1.1.9-2 2-2zm6.5 14l-1.41-1.41L12.67 12H3v-2h9.67l-2.58-2.59L11.5 6l5 5-5 5z\" />\n                  </svg>\n                  <span class=\"text-xl\">Logout</span>\n                </button>\n              </form>\n\n              <a href=\"/\" class=\"mt-4 flex items-center justify-center py-3 rounded-full font-bold text-[17px] transition-colors\" style=\"background:#1d9bf0;color:#fff\" onmouseover=\"this.style.background='#1a8cd8'\" onmouseout=\"this.style.background='#1d9bf0'\">\n                Post\n              </a>\n            ")
			if err != nil {
				return err
			}
		}
		_, err = io.WriteString(w, "\n          </div>\n\n          ")
		if err != nil {
			return err
		}
		if currentUser.ID > 0 {
			_, err = io.WriteString(w, "\n            <div class=\"mt-auto mb-3 w-[250px]\">\n              <a href=\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "/user/"+currentUser.Username)))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\" class=\"flex items-center gap-3 px-4 py-3 rounded-full transition-colors hover:bg-xhover\" style=\"text-decoration:none\">\n                ")
			if err != nil {
				return err
			}
			err = Avatar(currentUser.Username, "sm").Render(w)
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\n                <div class=\"flex-1 min-w-0\">\n                  <div class=\"font-bold text-[15px] leading-5 truncate\" style=\"color:#e7e9ea\">")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", currentUser.DisplayName)))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</div>\n                  <div class=\"text-[13px] leading-4 truncate\" style=\"color:#71767b\">@")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", currentUser.Username)))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "</div>\n                </div>\n              </a>\n            </div>\n          ")
			if err != nil {
				return err
			}
		}
		_, err = io.WriteString(w, "\n        </nav>\n\n        <main class=\"flex-1 max-w-[600px] min-h-screen\" style=\"border-left:1px solid #2f3336;border-right:1px solid #2f3336\">\n          ")
		if err != nil {
			return err
		}
		if children != nil {
			err = children.Render(w)
			if err != nil {
				return err
			}
		}
		_, err = io.WriteString(w, "\n        </main>\n\n        <aside class=\"hidden xl:block w-[350px] pl-7 sticky top-0 h-screen overflow-y-auto py-3\">\n          <div class=\"rounded-2xl p-4 mb-4\" style=\"background:#16181c\">\n            <h2 class=\"text-xl font-extrabold mb-1\" style=\"color:#e7e9ea\">Subscribe to Premium</h2>\n            <p class=\"text-[15px] mb-3\" style=\"color:#e7e9ea\">Subscribe to unlock new features.</p>\n            <a href=\"#\" class=\"inline-block px-4 py-2 rounded-full font-bold text-sm\" style=\"background:#1d9bf0;color:#fff\">Subscribe</a>\n          </div>\n\n          <div class=\"rounded-2xl overflow-hidden\" style=\"background:#16181c\">\n            <h2 class=\"text-xl font-extrabold px-4 pt-3 pb-2\" style=\"color:#e7e9ea\">What's happening</h2>\n            <div class=\"px-4 py-3 transition-colors\" style=\"border-bottom:1px solid #2f3336\" onmouseover=\"this.style.backgroundColor='rgba(231,233,234,0.03)'\" onmouseout=\"this.style.backgroundColor='transparent'\">\n              <div class=\"text-[13px]\" style=\"color:#71767b\">Trending</div>\n              <div class=\"font-bold text-[15px]\" style=\"color:#e7e9ea\">#GoxTemplates</div>\n              <div class=\"text-[13px]\" style=\"color:#71767b\">1,234 posts</div>\n            </div>\n            <div class=\"px-4 py-3 transition-colors\" onmouseover=\"this.style.backgroundColor='rgba(231,233,234,0.03)'\" onmouseout=\"this.style.backgroundColor='transparent'\">\n              <div class=\"text-[13px]\" style=\"color:#71767b\">Technology</div>\n              <div class=\"font-bold text-[15px]\" style=\"color:#e7e9ea\">#GoLang</div>\n              <div class=\"text-[13px]\" style=\"color:#71767b\">5,678 posts</div>\n            </div>\n          </div>\n        </aside>\n      </div>\n\n      <nav class=\"lg:hidden fixed bottom-0 left-0 right-0 flex justify-around py-3 px-4\" style=\"background:#000;border-top:1px solid #2f3336;z-index:50\">\n        <a href=\"/\" class=\"p-2\">\n          <svg viewBox=\"0 0 24 24\" class=\"w-6 h-6\" fill=\"currentColor\" style=\"color:#e7e9ea\">\n            <path d=\"M21.591 7.146L12.52 1.157c-.316-.21-.724-.21-1.04 0l-9.071 5.99c-.26.173-.409.456-.409.757v13.183c0 .502.418.913.929.913h6.638c.511 0 .929-.41.929-.913v-7.075h3.008v7.075c0 .502.418.913.929.913h6.638c.511 0 .929-.41.929-.913V7.904c0-.301-.158-.584-.408-.758z\" />\n          </svg>\n        </a>\n        <a href=\"/explore\" class=\"p-2\">\n          <svg viewBox=\"0 0 24 24\" class=\"w-6 h-6\" fill=\"currentColor\" style=\"color:#e7e9ea\">\n            <path d=\"M10.25 3.75c-3.59 0-6.5 2.91-6.5 6.5s2.91 6.5 6.5 6.5c1.795 0 3.419-.726 4.596-1.904 1.178-1.177 1.904-2.801 1.904-4.596 0-3.59-2.91-6.5-6.5-6.5zm-8.5 6.5c0-4.694 3.806-8.5 8.5-8.5s8.5 3.806 8.5 8.5c0 1.986-.682 3.815-1.824 5.262l4.781 4.781-1.414 1.414-4.781-4.781c-1.447 1.142-3.276 1.824-5.262 1.824-4.694 0-8.5-3.806-8.5-8.5z\" />\n          </svg>\n        </a>\n        ")
		if err != nil {
			return err
		}
		if currentUser.ID > 0 {
			_, err = io.WriteString(w, "\n          <a href=\"")
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", "/user/"+currentUser.Username)))
			if err != nil {
				return err
			}
			_, err = io.WriteString(w, "\" class=\"p-2\">\n            <svg viewBox=\"0 0 24 24\" class=\"w-6 h-6\" fill=\"currentColor\" style=\"color:#e7e9ea\">\n              <path d=\"M5.651 19h12.698c-.337-1.8-1.023-3.21-1.945-4.19C15.318 13.65 13.838 13 12 13s-3.317.65-4.404 1.81c-.922.98-1.608 2.39-1.945 4.19zm.486-5.56C7.627 11.85 9.648 11 12 11s4.373.85 5.863 2.44c1.477 1.58 2.366 3.8 2.632 6.46l.11 1.1H3.395l.11-1.1c.266-2.66 1.155-4.88 2.632-6.46zM12 4c-1.105 0-2 .9-2 2s.895 2 2 2 2-.9 2-2-.895-2-2-2zM8 6c0-2.21 1.791-4 4-4s4 1.79 4 4-1.791 4-4 4-4-1.79-4-4z\" />\n            </svg>\n          </a>\n        ")
			if err != nil {
				return err
			}
		}
		_, err = io.WriteString(w, "\n      </nav>\n    </body>\n  </html>\n")
		if err != nil {
			return err
		}
		return nil
	})
}

func AuthLayout(title string, children gox.Component) gox.Component {
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
		_, err = io.WriteString(w, "</title>\n      <script src=\"https://cdn.tailwindcss.com\"></script>\n      <link href=\"https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap\" rel=\"stylesheet\" />\n      ")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, fmt.Sprintf("%v", `<style>*{font-family:'Inter',system-ui,sans-serif}body{background:#000;color:#e7e9ea;margin:0}::selection{background:#1d9bf0;color:#fff}input:focus{outline:none}</style>`))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\n    </head>\n    <body class=\"min-h-screen flex items-center justify-center\">\n      ")
		if err != nil {
			return err
		}
		if children != nil {
			err = children.Render(w)
			if err != nil {
				return err
			}
		}
		_, err = io.WriteString(w, "\n    </body>\n  </html>\n")
		if err != nil {
			return err
		}
		return nil
	})
}

func PageHeader(title string) gox.Component {
	return gox.ComponentFunc(func(w io.Writer) error {
		var err error
		_ = err
		_, err = io.WriteString(w, "\n  <div class=\"sticky top-0 px-4 py-3 flex items-center gap-6 backdrop-blur-md\" style=\"background:rgba(0,0,0,0.65);z-index:40;border-bottom:1px solid #2f3336\">\n    <h1 class=\"text-xl font-bold\" style=\"color:#e7e9ea\">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", title)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</h1>\n  </div>\n")
		if err != nil {
			return err
		}
		return nil
	})
}

func PageHeaderWithBack(title string, backURL string) gox.Component {
	return gox.ComponentFunc(func(w io.Writer) error {
		var err error
		_ = err
		_, err = io.WriteString(w, "\n  <div class=\"sticky top-0 px-4 py-1 flex items-center gap-6 backdrop-blur-md\" style=\"background:rgba(0,0,0,0.65);z-index:40;border-bottom:1px solid #2f3336\">\n    <a href=\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", backURL)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "\" class=\"p-2 -ml-2 rounded-full transition-colors hover:bg-xhover\" style=\"color:#e7e9ea\">\n      <svg viewBox=\"0 0 24 24\" class=\"w-5 h-5\" fill=\"currentColor\">\n        <path d=\"M7.414 13l5.043 5.04-1.414 1.42L3.586 12l7.457-7.46 1.414 1.42L7.414 11H21v2H7.414z\" />\n      </svg>\n    </a>\n    <div>\n      <h1 class=\"text-xl font-bold leading-6\" style=\"color:#e7e9ea\">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", title)))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "</h1>\n      <span class=\"text-[13px]\" style=\"color:#71767b\">")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", strconv.Itoa(0))))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " posts</span>\n    </div>\n  </div>\n")
		if err != nil {
			return err
		}
		return nil
	})
}
