package views

import (
	"context"
	"fmt"
	"github.com/buemura/gox"
	"html"
	"io"
)

func LoginPage(errorMsg string) gox.Component {
	return gox.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		var err error
		_ = err
		_, err = io.WriteString(w, "\n  ")
		if err != nil {
			return err
		}
		err = AuthLayout("Log in / Gox Social", gox.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			var err error
			_ = err
			_, err = io.WriteString(w, "\n    <div class=\"w-full max-w-[364px] px-8\">\n      <div class=\"flex justify-center mb-8\">\n        <svg viewBox=\"0 0 24 24\" class=\"w-9 h-9\" fill=\"currentColor\" style=\"color:#e7e9ea\">\n          <path d=\"M18.244 2.25h3.308l-7.227 8.26 8.502 11.24H16.17l-5.214-6.817L4.99 21.75H1.68l7.73-8.835L1.254 2.25H8.08l4.713 6.231zm-1.161 17.52h1.833L7.084 4.126H5.117z\" />\n        </svg>\n      </div>\n\n      <h1 class=\"text-3xl font-bold mb-8\" style=\"color:#e7e9ea\">Sign in</h1>\n\n      ")
			if err != nil {
				return err
			}
			if errorMsg != "" {
				_, err = io.WriteString(w, "\n        <div class=\"mb-4 px-4 py-3 rounded-xl text-sm\" style=\"background:rgba(244,33,46,0.1);border:1px solid rgba(244,33,46,0.2);color:#f4212e\">\n          ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", errorMsg)))
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n        </div>\n      ")
				if err != nil {
					return err
				}
			}
			_, err = io.WriteString(w, "\n\n      <form method=\"POST\" action=\"/login\" class=\"flex flex-col gap-5\">\n        <div class=\"relative\">\n          <input type=\"text\" name=\"username\" required placeholder=\" \" id=\"login-user\" class=\"peer w-full px-3 pt-6 pb-2 rounded-md text-[17px] bg-transparent\" style=\"border:1px solid #2f3336;color:#e7e9ea;caret-color:#1d9bf0\" />\n          <label for=\"login-user\" class=\"absolute left-3 top-1/2 -translate-y-1/2 text-[17px] transition-all pointer-events-none peer-focus:top-4 peer-focus:text-xs peer-[:not(:placeholder-shown)]:top-4 peer-[:not(:placeholder-shown)]:text-xs\" style=\"color:#71767b\">Username</label>\n        </div>\n\n        <div class=\"relative\">\n          <input type=\"password\" name=\"password\" required placeholder=\" \" id=\"login-pass\" class=\"peer w-full px-3 pt-6 pb-2 rounded-md text-[17px] bg-transparent\" style=\"border:1px solid #2f3336;color:#e7e9ea;caret-color:#1d9bf0\" />\n          <label for=\"login-pass\" class=\"absolute left-3 top-1/2 -translate-y-1/2 text-[17px] transition-all pointer-events-none peer-focus:top-4 peer-focus:text-xs peer-[:not(:placeholder-shown)]:top-4 peer-[:not(:placeholder-shown)]:text-xs\" style=\"color:#71767b\">Password</label>\n        </div>\n\n        <button type=\"submit\" class=\"w-full py-3 rounded-full font-bold text-[17px] transition-colors\" style=\"background:#e7e9ea;color:#0f1419\" onmouseover=\"this.style.background='#d7dbdc'\" onmouseout=\"this.style.background='#e7e9ea'\">\n          Sign in\n        </button>\n      </form>\n\n      <p class=\"mt-10 text-base\" style=\"color:#71767b\">\n        Don't have an account?\n        <a href=\"/register\" class=\"hover:underline\" style=\"color:#1d9bf0\"> Sign up</a>\n      </p>\n    </div>\n  ")
			if err != nil {
				return err
			}
			return nil
		})).Render(ctx, w)
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
	return gox.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		var err error
		_ = err
		_, err = io.WriteString(w, "\n  ")
		if err != nil {
			return err
		}
		err = AuthLayout("Create account / Gox Social", gox.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			var err error
			_ = err
			_, err = io.WriteString(w, "\n    <div class=\"w-full max-w-[364px] px-8\">\n      <div class=\"flex justify-center mb-8\">\n        <svg viewBox=\"0 0 24 24\" class=\"w-9 h-9\" fill=\"currentColor\" style=\"color:#e7e9ea\">\n          <path d=\"M18.244 2.25h3.308l-7.227 8.26 8.502 11.24H16.17l-5.214-6.817L4.99 21.75H1.68l7.73-8.835L1.254 2.25H8.08l4.713 6.231zm-1.161 17.52h1.833L7.084 4.126H5.117z\" />\n        </svg>\n      </div>\n\n      <h1 class=\"text-3xl font-bold mb-8\" style=\"color:#e7e9ea\">Create your account</h1>\n\n      ")
			if err != nil {
				return err
			}
			if errorMsg != "" {
				_, err = io.WriteString(w, "\n        <div class=\"mb-4 px-4 py-3 rounded-xl text-sm\" style=\"background:rgba(244,33,46,0.1);border:1px solid rgba(244,33,46,0.2);color:#f4212e\">\n          ")
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", errorMsg)))
				if err != nil {
					return err
				}
				_, err = io.WriteString(w, "\n        </div>\n      ")
				if err != nil {
					return err
				}
			}
			_, err = io.WriteString(w, "\n\n      <form method=\"POST\" action=\"/register\" class=\"flex flex-col gap-5\">\n        <div class=\"relative\">\n          <input type=\"text\" name=\"display_name\" required placeholder=\" \" id=\"reg-name\" class=\"peer w-full px-3 pt-6 pb-2 rounded-md text-[17px] bg-transparent\" style=\"border:1px solid #2f3336;color:#e7e9ea;caret-color:#1d9bf0\" />\n          <label for=\"reg-name\" class=\"absolute left-3 top-1/2 -translate-y-1/2 text-[17px] transition-all pointer-events-none peer-focus:top-4 peer-focus:text-xs peer-[:not(:placeholder-shown)]:top-4 peer-[:not(:placeholder-shown)]:text-xs\" style=\"color:#71767b\">Display name</label>\n        </div>\n\n        <div class=\"relative\">\n          <input type=\"text\" name=\"username\" required placeholder=\" \" id=\"reg-user\" class=\"peer w-full px-3 pt-6 pb-2 rounded-md text-[17px] bg-transparent\" style=\"border:1px solid #2f3336;color:#e7e9ea;caret-color:#1d9bf0\" />\n          <label for=\"reg-user\" class=\"absolute left-3 top-1/2 -translate-y-1/2 text-[17px] transition-all pointer-events-none peer-focus:top-4 peer-focus:text-xs peer-[:not(:placeholder-shown)]:top-4 peer-[:not(:placeholder-shown)]:text-xs\" style=\"color:#71767b\">Username</label>\n        </div>\n\n        <div class=\"relative\">\n          <input type=\"email\" name=\"email\" required placeholder=\" \" id=\"reg-email\" class=\"peer w-full px-3 pt-6 pb-2 rounded-md text-[17px] bg-transparent\" style=\"border:1px solid #2f3336;color:#e7e9ea;caret-color:#1d9bf0\" />\n          <label for=\"reg-email\" class=\"absolute left-3 top-1/2 -translate-y-1/2 text-[17px] transition-all pointer-events-none peer-focus:top-4 peer-focus:text-xs peer-[:not(:placeholder-shown)]:top-4 peer-[:not(:placeholder-shown)]:text-xs\" style=\"color:#71767b\">Email</label>\n        </div>\n\n        <div class=\"relative\">\n          <input type=\"password\" name=\"password\" required placeholder=\" \" id=\"reg-pass\" class=\"peer w-full px-3 pt-6 pb-2 rounded-md text-[17px] bg-transparent\" style=\"border:1px solid #2f3336;color:#e7e9ea;caret-color:#1d9bf0\" />\n          <label for=\"reg-pass\" class=\"absolute left-3 top-1/2 -translate-y-1/2 text-[17px] transition-all pointer-events-none peer-focus:top-4 peer-focus:text-xs peer-[:not(:placeholder-shown)]:top-4 peer-[:not(:placeholder-shown)]:text-xs\" style=\"color:#71767b\">Password</label>\n        </div>\n\n        <button type=\"submit\" class=\"w-full py-3 rounded-full font-bold text-[17px] transition-colors\" style=\"background:#e7e9ea;color:#0f1419\" onmouseover=\"this.style.background='#d7dbdc'\" onmouseout=\"this.style.background='#e7e9ea'\">\n          Create account\n        </button>\n      </form>\n\n      <p class=\"mt-10 text-base\" style=\"color:#71767b\">\n        Already have an account?\n        <a href=\"/login\" class=\"hover:underline\" style=\"color:#1d9bf0\"> Sign in</a>\n      </p>\n    </div>\n  ")
			if err != nil {
				return err
			}
			return nil
		})).Render(ctx, w)
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
