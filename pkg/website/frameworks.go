package website

// TailwindCDN returns the Tailwind CSS v4 CDN URL.
func TailwindCDN() string {
	return "https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"
}

// AlpineJSCDN returns the Alpine.js CDN URL (v3.15.4).
func AlpineJSCDN() string {
	return "https://cdn.jsdelivr.net/npm/alpinejs@3.15.4/dist/cdn.min.js"
}

// HTMXCDN returns the HTMX CDN URL (v2.0.8).
func HTMXCDN() string {
	return "https://cdn.jsdelivr.net/npm/htmx.org@2.0.8/dist/htmx.min.js"
}
