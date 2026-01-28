---
title: "How I work with AI: my coding partner instructions for Go"
date: "2026-01-26"
author: "Maciej Adamski"
published: true
description: "I wrote 1,400 lines of instructions that turn any AI into my ideal coding mentor. Simplicity-first philosophy, Go patterns, NASA-inspired defensive programming, and structured observability."
---

**TL;DR:** I created a 1,400-line document that transforms AI coding assistants into principled mentors. It encodes my software philosophy (simplicity, YAGNI, explicit control), Go-specific patterns, NASA-inspired defensive programming, and a structured engineering journal system. The AI guides me by default but switches to implementation mode on command.

---

**Link to the full instructions: [AI Coding Partner Instructions](https://gist.github.com/maciejadamski89/06d6e4a3a8c1d6d481f974d763c16c6e)**

I spend most of my coding hours paired with AI. Not as a replacement for thinking, but as a mentor, reviewer, and sparring partner. Over months of daily collaboration, I noticed a pattern: the AI would sometimes write clever code instead of simple code, dump entire solutions when I wanted to learn, or forget the principles I care about most.

So I wrote a document. A comprehensive set of instructions that tells the AI exactly how I want to work.

This is not a prompt template. It is a philosophy of software development, encoded into rules that any AI can follow. I am sharing it publicly because I believe it can help others build better software and become better engineers.


## Why did I create custom AI instructions?

I have two goals when I code:

1. Build simple, maintainable software that works.
2. Become a better engineer with every line I write.

These goals are often in tension. If I want to learn, I should struggle. If I want to ship, I should let the AI write the code. The instructions solve this by giving me explicit control. By default, the AI guides me through problems without giving answers. But when I say "just do it", it switches to implementation mode.

The instructions also encode my philosophy. I follow the masters: Rob Pike's "less is more", Rich Hickey's "simple vs easy", Linus Torvalds' hatred of complexity, Sandi Metz's refactoring rules. Every time I start a new chat, the AI already knows what I value.


## What is the core philosophy?

The document opens with a clear statement:

Simplicity is the only metric that matters. All code must be, above all else, simple, clear, and obvious.

This is not a platitude. It is a filter for every decision. Should I add this abstraction? Is it simpler? Should I use this clever one-liner? Is it clearer? Should I optimize this function? Do I have proof it is slow?

The instructions include YAGNI (You Ain't Gonna Need It) as a core principle. We ruthlessly cut any code, argument, or conditional that is not required right now. The AI is trained to challenge me when I start building for hypothetical futures.


## How does the collaboration workflow work?

This is perhaps the most important section. It defines how we work together:

1. The AI provides a statement and proposition about the topic.
2. It provides examples and paths to follow.
3. It stops and waits for my input.
4. I implement it myself.
5. The AI reviews my code and provides feedback.
6. We repeat until the task is complete.

The AI does not dump full solutions. It does not proceed without my confirmation. It treats me as the engineer who must understand and own every line.

But there is an escape hatch. When I say "please implement it" or "vibe coding session", the AI switches modes and provides complete code immediately. Sometimes I need to ship fast. The instructions respect that.


## What Go patterns do I enforce?

The document contains an extensive section on Go patterns. These are not suggestions. They are rules that shape every function, every interface, every package.

Interface design: accept interfaces, return concrete types. Keep interfaces small (one to three methods). Define interfaces where they are used, not where they are implemented.

Package design: package by feature, not by layer. Use internal/ to hide implementation. Short, lowercase names. No circular dependencies.

Error handling: wrap errors with context using fmt.Errorf and %w. Use errors.Is() and errors.As() for checking. Define sentinel errors for expected failures. Never panic in library code.

Concurrency: never start a goroutine you cannot stop. Always pass context.Context as the first parameter. Use errgroup for coordination. Close channels from the sender.

The document includes concrete examples for each rule. Bad code and good code, side by side. The AI sees exactly what I expect.


## What is defensive programming?

I borrowed principles from NASA's Power of 10 rules. These are practices from safety-critical software development that prevent entire categories of bugs.

Loop safety: every loop must have a clear termination condition. Input collections must have size limits. Retry loops must have maximum attempts. Pagination must have safety limits. Worker loops must be stoppable via context.

Bounds checking: never trust slice indices from external input. Check explicitly before accessing.

Static analysis: run golangci-lint with strict settings. Zero warnings tolerance. Never merge code with linter warnings.

These rules sound paranoid until you have debugged an infinite loop in production at 3 AM.


## How do I handle observability and logging?

Logging is not an afterthought. It is a core feature. The instructions define exactly what good logging looks like.

When I switch the log level to DEBUG, I want to see the complete data flow. Every request should have a trace ID. Every log entry should include structured context. The message should be static; dynamic data goes in fields.

The document includes complete examples of middleware, context propagation, and handler logging. When DEBUG is enabled, the logs tell a story:

```
DEBUG trace_id=abc123 msg="incoming request" method=POST path=/api/orders
DEBUG trace_id=abc123 msg="parsing request body" content_type=application/json
DEBUG trace_id=abc123 msg="validation passed"
DEBUG trace_id=abc123 msg="checking inventory" sku=WIDGET-001
INFO  trace_id=abc123 msg="order created" order_id=order_789 total=107.99
DEBUG trace_id=abc123 msg="request completed" status=201 duration_ms=45
```

This level of detail saves hours when something goes wrong.


## What is the engineering journal system?

The instructions include a system for documentation that I call the engineering journal. It has three components:

ADR (architecture decision records): when we agree on a significant decision, the AI generates a formatted entry I can save. Each decision has a global number (D-001, D-002) that I can reference in code comments.

Learning entries: when the AI explains a concept I asked about, it generates a summary I can save. This prevents me from asking the same question twice.

Daily log: a simple bullet journal for capturing thoughts during work. One file per month.

The AI knows to ask me for the next number before generating entries. It knows to generate ADRs automatically when we reach agreement. Knowledge does not disappear into chat history.


## What changes after using these instructions?

After using these instructions for months, I notice several changes:

Code quality improved. The AI challenges my names, my structure, my error handling. It asks questions like "What would you name this function if it could only do one thing?"

Learning accelerated. Because I write the code myself by default, I understand it deeply. The AI is a mentor, not a crutch.

Debugging became systematic. The instructions define a factual debugging process. No guessing. Check each component one by one.

Consistency increased. Every file follows the same patterns. Every error is handled the same way. Every log entry has the same structure.


## How can you use these instructions?

The full document is available as a GitHub Gist. You can use it directly or adapt it to your own philosophy.

To use it with Claude, Cursor, Windsurf, or similar tools, paste it into your system prompt or custom instructions. The AI will follow it from the first message.

I encourage you to read it completely before using it. Understand the philosophy. Disagree with parts if you want. Make it yours. The goal is not to follow my rules. The goal is to have rules at all.

Link to the full instructions: [AI Coding Partner Instructions](https://gist.github.com/maciejadamski89/06d6e4a3a8c1d6d481f974d763c16c6e)


## Frequently Asked Questions

**Can I use these instructions with any AI coding assistant?**

Yes. The instructions work with Claude, ChatGPT, Cursor, Windsurf, and any AI that accepts system prompts or custom instructions. The format is plain markdown.

**Do I need to use Go to benefit from this document?**

No. While the Go patterns section is language-specific, the philosophy, collaboration workflow, and engineering journal system apply to any language. You can replace the Go section with patterns for Python, TypeScript, or Rust.

**How do I switch between learning mode and "vibe coding" mode?**

Say "please implement it" or "vibe coding session" to trigger implementation mode. The AI will provide complete code immediately instead of guiding you step by step.

**How long did it take to write these instructions?**

The initial version took a weekend. But I refine it continuously as I discover new patterns or edge cases. It is a living document.


## Final thoughts

AI is a tool. Like any tool, it amplifies what you bring to it. If you bring no principles, you get inconsistent code. If you bring clear principles, you get a collaborator who helps you uphold them.

These instructions are my principles, encoded. They represent years of learning from masters like Pike, Hickey, Torvalds, and Metz. They represent months of refining how I work with AI.

I share them hoping they help you build simpler software and become a better engineer.

Clear is better than clever. Always.


### About the Author

Maciej Adamski is a software engineer and founder of Dataglitch, specializing in Go backend development and high-performance web applications. He writes about software craftsmanship, AI-assisted development, and the pursuit of simplicity in code.