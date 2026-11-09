# Reginald Alfret V — Portfolio (Go Buffalo)

> **Go & Buffalo Web Framework Implementation** of Reginald Alfret's Liquid-Glass Obsidian Resume Portfolio.

Explore the live unified multi-stack portfolio at [reginaldalfret.tech](https://reginaldalfret.tech) or the direct Go Buffalo build at [reginaldalfret.tech/stacks/buffalo](https://reginaldalfret.tech/stacks/buffalo).

---

## ⚡ Architecture & Tech Stack

- **Framework**: [Go Buffalo](https://gobuffalo.io/) (v1.1.4) + [Go 1.27](https://golang.org/)
- **Templating**: [Plush v5](https://github.com/gobuffalo/plush) Server-Side Rendering
- **Styling**: Tailwind CSS + Obsidian Glassmorphic Palette (`#06B6D4` Buffalo Turquoise branding)
- **Features**:
  - 30 Merged Open-Source PRs across Apple, Google, Microsoft, Meta, AWS, NVIDIA, NASA, Netflix, Linux Foundation & Apache
  - Professional Experience, Hackathon Victories, Research, and Certified Credentials
  - Interactive PDF Modal / Original Resume Viewer
  - Seamless Multi-Stack Cross-Navigation Dock & Switcher (9 Stacks)
  - Synchronous `is-embedded` iframe detection preventing duplicate navigation UI
  - Standalone embedded binary compilation via Go `embed.FS`

---

## 🚀 Running Locally

```bash
# Ensure Go 1.27+ is installed
go version

# Clone and run
git clone https://github.com/reginaldalfret/portfolio-buffalo.git
cd portfolio-buffalo

# Run server on port 8087
go run .
```

Open [http://127.0.0.1:8087](http://127.0.0.1:8087) in your browser.

---

## 📦 Static Export

```bash
go run . -export -out ./dist
```
Exports standalone, self-contained static HTML assets ready for edge hosting on Vercel, Netlify, or AWS CloudFront.
