---
title: "Cleaning the digital attic; automating Cloudflare Pages with Go"
date: "2026-01-21"
draft: false
description: "A concurrency-first Go utility to clean up old Cloudflare Pages deployments safely and fast."
---

In modern web development, "set it and forget it" is a lie. I ran into a different kind of mess while cleaning up my Cloudflare organization. I could not delete unused Cloudflare Workers for my websites because there were hundreds of old Cloudflare Pages deployments tied to those previews. Every branch, every preview, and every minor tweak was living forever in the Cloudflare cloud.

Cloudflare is generous, but thousands of dead deployments make the UI hard to navigate and clutter the project history. I needed a way to wipe the old deployments while keeping production safe and untouched.

## The problem: manual cleanup doesn't scale

Cloudflare's dashboard is fine for a few deployments, but manually clicking "Delete" on hundreds is not a task for an engineer. In my case I had more than 400 deployments, and deleting them by hand would take a lot of time. I needed a tool that was:

- **Fast** > It should handle deletions in parallel.
- **Safe** > It must never, under any circumstances, delete the active production site.
- **Simple** > No heavy dependencies. Just a single binary I can run from my terminal.

## The solution: a concurrency-first Go utility

I chose Go for this task because of its first-class support for concurrency and its clear, practical approach to systems tooling.

### 1. The "bouncer" pattern (semaphores)

When hitting an API, you cannot just fire off 1,000 requests at once. You will hit rate limits or exhaust your local network sockets. I implemented a semaphore pattern using Go channels.

By using a buffered channel, I created a "bouncer" that only allows 15 workers into the "club" at any given time. This keeps the execution fast but respectful of Cloudflare's API limits.

### 2. Failing fast and being "production-aware"

The most dangerous part of an automated deletion script is the risk of deleting the live site. I designed the filtering logic to be strictly production-aware.

Before the script even attempts a deletion, it inspects the deployment metadata. If the environment is marked as production, the script skips it entirely. Even if that check fails, I added a secondary layer: the script parses the raw JSON error bodies from Cloudflare. When we hit the specific error `8000034` ("Cannot delete active production"), the script handles it gracefully instead of crashing.

### 3. The "maintainless" workflow

Following the KISS (Keep It Simple, Stupid) principle, I wrapped the entire logic in a Makefile. I do not want to remember complex CLI flags or environment variable exports.

- `make run`: Cleans up the default project.
- `make nuke project_name=xxx`: Targets a specific site.

This local-first approach ensures I can verify the cleanup on my machine before even thinking about putting it into a CI/CD pipeline.

## Engineering lessons learned

- **Simple is not easy** > Writing a script that does not get stuck in an infinite loop when it hits a protected resource takes more thought than just a basic for loop.
- **Stream the data** > Using `json.NewDecoder` allowed the script to handle large batches of deployment data without spiking memory usage.
- **Secrets belong in the environment** > Using `.env` files and `os.Getenv` kept my API tokens safe and out of the source code.

## Check out the code

You can find the full source code, the Makefile, and the documentation on my GitHub: <https://github.com/dg-dataglitch/cloudflare-pages-cleanup>
