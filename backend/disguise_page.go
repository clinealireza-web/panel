package main

// Built-in cover site template: a plain, generic digital-services landing
// page. No panel branding, no panel wording, no references to VPN or proxy
// anywhere in markup, copy, class names, or asset names.

import (
	"fmt"
	"html"
	"strings"
)

func builtinCoverHTML(name, tagline, title, body string) string {
	name = strings.TrimSpace(name)
	tagline = strings.TrimSpace(tagline)
	if name == "" {
		name = "Nova Digital"
	}
	if tagline == "" {
		tagline = "Modern tools for growing businesses"
	}
	pageTitle := name
	if strings.TrimSpace(title) != "" {
		pageTitle = fmt.Sprintf("%s | %s", title, name)
	}
	extra := ""
	if strings.TrimSpace(body) != "" {
		extra = "<section class=\"card\"><h2>" + html.EscapeString(title) + "</h2><p>" + html.EscapeString(body) + "</p></section>"
	}
	return `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>` + html.EscapeString(pageTitle) + `</title>
<meta name="description" content="` + html.EscapeString(tagline) + `">
<style>
:root{--bg:#f6f7fb;--fg:#1c2333;--mut:#5b6478;--acc:#3457d5;--card:#ffffff;--line:#e4e7f0}
*{margin:0;padding:0;box-sizing:border-box}
body{font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Helvetica,Arial,sans-serif;background:var(--bg);color:var(--fg);line-height:1.6}
header{background:var(--card);border-bottom:1px solid var(--line);position:sticky;top:0;z-index:10}
.wrap{max-width:1080px;margin:0 auto;padding:0 20px}
.nav{display:flex;align-items:center;justify-content:space-between;height:64px}
.brand{font-weight:700;font-size:20px;color:var(--fg);text-decoration:none;letter-spacing:-.02em}
.brand span{color:var(--acc)}
nav ul{display:flex;gap:26px;list-style:none}
nav a{color:var(--mut);text-decoration:none;font-size:15px}
nav a:hover{color:var(--fg)}
.btn{display:inline-block;background:var(--acc);color:#fff;padding:10px 22px;border-radius:8px;text-decoration:none;font-weight:600;font-size:15px}
.hero{padding:90px 0 70px;text-align:center}
.hero h1{font-size:44px;line-height:1.15;letter-spacing:-.03em;max-width:720px;margin:0 auto 18px}
.hero p{font-size:18px;color:var(--mut);max-width:560px;margin:0 auto 32px}
.grid{display:grid;grid-template-columns:repeat(auto-fit,minmax(240px,1fr));gap:20px;padding:30px 0 70px}
.card{background:var(--card);border:1px solid var(--line);border-radius:12px;padding:26px}
.card h2{font-size:17px;margin-bottom:8px}
.card p{font-size:14px;color:var(--mut)}
.ico{width:42px;height:42px;border-radius:10px;background:#eef1fe;color:var(--acc);display:flex;align-items:center;justify-content:center;font-size:20px;margin-bottom:14px}
footer{border-top:1px solid var(--line);padding:28px 0;color:var(--mut);font-size:14px;text-align:center}
@media(max-width:640px){nav ul{display:none}.hero h1{font-size:32px}.hero{padding:60px 0 40px}}
</style>
</head>
<body>
<header><div class="wrap"><div class="nav">
<a class="brand" href="/"><span>` + html.EscapeString(name[:1]) + `</span>` + html.EscapeString(name[1:]) + `</a>
<nav><ul>
<li><a href="/">Home</a></li>
<li><a href="/features">Features</a></li>
<li><a href="/pricing">Pricing</a></li>
<li><a href="/contact">Contact</a></li>
</ul></nav>
<a class="btn" href="/contact">Get Started</a>
</div></div></header>
<main class="wrap">
` + extra + `
<section class="hero">
<h1>` + html.EscapeString(tagline) + `</h1>
<p>Everything your team needs to ship faster, in one simple workspace. No credit card required.</p>
<a class="btn" href="/contact">Start Free Trial</a>
</section>
<section class="grid">
<div class="card"><div class="ico">&#9889;</div><h2>Fast Setup</h2><p>Get running in minutes with a guided onboarding flow designed for teams of any size.</p></div>
<div class="card"><div class="ico">&#128274;</div><h2>Secure by Default</h2><p>Industry-standard encryption and access controls are enabled out of the box.</p></div>
<div class="card"><div class="ico">&#128200;</div><h2>Detailed Analytics</h2><p>Track usage and performance with clear, real-time dashboards and reports.</p></div>
<div class="card"><div class="ico">&#129309;</div><h2>Team Collaboration</h2><p>Share projects, assign roles, and keep everyone aligned in one place.</p></div>
</section>
</main>
<footer><div class="wrap">&copy; 2026 ` + html.EscapeString(name) + `. All rights reserved.</div></footer>
</body>
</html>`
}
