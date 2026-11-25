package main

// Disguise mode: when enabled, every unauthenticated visitor sees a plain
// commercial website. The panel UI, login page, and panel-flavored API
// responses stay behind a secret gate path / a gate cookie.

import (
	"crypto/subtle"
	"net/http"
	"os"
	"strings"
)

const gateCookieName = "sid"

func disguiseEnabled() bool {
	State.mu.RLock()
	defer State.mu.RUnlock()
	v, _ := State.Settings["disguise_enabled"].(bool)
	return v
}

func disguiseSecret() string {
	State.mu.RLock()
	defer State.mu.RUnlock()
	s, _ := State.Settings["disguise_secret"].(string)
	return strings.TrimSpace(s)
}

// disguiseGateToken is the value the gate cookie must carry. It is derived
// from the secret so that rotating the secret invalidates old cookies and
// so the cookie itself does not equal the secret.
func disguiseGateToken() string {
	return sha256Hex("gate:" + disguiseSecret())[:32]
}

func disguiseGateCookie(r *http.Request) bool {
	c, err := r.Cookie(gateCookieName)
	if err != nil || c.Value == "" {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(c.Value), []byte(disguiseGateToken())) == 1
}

func disguiseSetGateCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     gateCookieName,
		Value:    disguiseGateToken(),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   86400,
	})
}

// writeDisguise404 renders the cover site for unknown paths. Visitors must
// not be able to probe which paths exist: everything non-essential looks
// like a 404 of the cover site.
func writeDisguise404(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte(coverPageHTML("404", "")))
}

// disguiseGuard wraps handlers so that, in disguise mode, panel surfaces are
// hidden from anyone without the gate cookie. Transport endpoints (/ws/,
// /xhttp-siz10/, /sub/, /link/, /api/sub/) are exempt: VPN clients carry no
// cookie and must keep working.
func disguiseGuard(exempt bool, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if exempt || !disguiseEnabled() || disguiseGateCookie(r) {
			next(w, r)
			return
		}
		writeDisguise404(w)
	}
}

// coverPageHTML renders the cover site. It looks for a user-supplied
// cover.html in the data dir first so operators can fully customize it;
// otherwise it falls back to the built-in template parameterized by
// disguise_site_name / disguise_site_tagline.
func coverPageHTML(title, body string) string {
	State.mu.RLock()
	name, _ := State.Settings["disguise_site_name"].(string)
	tagline, _ := State.Settings["disguise_site_tagline"].(string)
	dataDir := State.DataDir
	State.mu.RUnlock()

	if dataDir != "" {
		if b, err := os.ReadFile(dataDir + "/cover.html"); err == nil && len(b) > 0 {
			return string(b)
		}
	}
	if strings.TrimSpace(name) == "" {
		name = "Nova Digital"
	}
	if strings.TrimSpace(tagline) == "" {
		tagline = "Modern tools for growing businesses"
	}
	return builtinCoverHTML(name, tagline, title, body)
}
