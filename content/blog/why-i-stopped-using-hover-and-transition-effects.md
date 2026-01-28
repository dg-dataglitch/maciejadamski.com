---
title: "Why I stopped using hover and transition effects"
date: "2026-01-23"
author: "Maciej Adamski"
published: true
description: "In a mobile-first world, hover and transition effects are mostly dead code. Why I prefer simpler Tailwind, faster pages, and better readability."
---

**TL;DR:** Hover and transition effects are desktop ghosts in a mobile-first world. They double your Tailwind class lists, cause "jank" on budget smartphones, and add zero value for touch users. I removed them from my projects, and the code became more readable, the pages faster, and the UX more consistent.

---

In a mobile-first world, fancy UI is just technical debt.


## Why do hover effects hurt code readability?

I use TailwindCSS for every project; it offers total control without leaving the HTML. But that control comes with a price: **visual noise**. Between layout, spacing, and typography, a single button can easily carry ten classes; when you add "hover," "transition-all," and "duration-300," you are not just styling; you are cluttering. You are making the code harder to read for a feature most users will never see.


## How bloated can a single button get?

Consider a standard "get started" button. To make it functional, accessible, and "pretty" by modern standards, your HTML looks like this:

```html
<div class="mt-10 flex items-center justify-center gap-x-6">
  <a href="#" class="rounded-md bg-indigo-500 px-3.5 py-2.5 text-sm font-semibold text-white shadow-xs transition-all duration-300 ease-in-out hover:bg-indigo-400 hover:scale-105 active:scale-95 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500">
    Get started
  </a>
  <a href="#" class="text-sm/6 font-semibold text-white">
    Learn more <span aria-hidden="true">→</span>
  </a>
</div>
```

Look at that class string. Half of those classes; specifically hover:bg-indigo-400, hover:scale-105, transition-all, and duration-300; are desktop ghosts. They exist for a mouse cursor that is not there on a mobile device; by chasing "polish," we have doubled the length of our code. It is now harder for any developer to scan the file and understand the actual structure of the page.


## Why are hover effects useless on mobile?

We need to stop designing for ourselves; most users visit your site from a phone. Hover effects do not work on mobile; when a user taps a button on a smartphone, the "hover" state might flicker for a millisecond; or worse, it gets stuck in a "persistent hover" state that requires a second tap to clear. We are shipping dead code to our users' browsers.

Transitions and backdrop-blur are even worse; they require the mobile processor to "repaint" the screen constantly. On a budget smartphone, this does not look smooth; it looks like "jank." It makes your premium UI feel like a cheap toy.


## What can Amazon teach us about UI design?

If "pretty" was the secret to success, Amazon would have gone bankrupt a decade ago. The Amazon UI looks like a relic from the 90s; it is "ugly" by design school standards. But Amazon understands something most designers forget: users do not visit your site to admire the buttons. They are there to solve a problem.

Amazon spent millions to learn that:

- **Predictability beats novelty:** Users know exactly where the "Add to Cart" button is.
- **Speed is a feature:** Every millisecond of delay costs them money.
- **Consistency creates trust:** When a system works the same way every time, users feel safe.

They do not change their UI to match the latest trends; they know it adds zero value to the customer's mission.


## What should you prioritize instead?

A senior engineer's value is not measured by how many fancy CSS tricks they can cram into a tag; it is measured by how much value they deliver with the least amount of complexity.

- **Prioritize focus over hover:** Spend your time on focus-visible and accessibility; that helps everyone.
- **Prioritize speed over polish:** A site that loads in 500ms with a "90s design" will always have a better UX than a "modern" site that takes 3 seconds to animate in.
- **Prioritize readability over control:** Keep your Tailwind classes lean; if you do not need the transition, delete it.


## Frequently Asked Questions

**Should I remove ALL hover effects?**

Not necessarily. Subtle hover effects on navigation links or important CTAs can provide useful feedback for desktop users. The key is to be intentional: if it doesn't add value, remove it.

**What about accessibility? Don't hover effects help?**

No. Focus states help accessibility; hover states do not. Spend your effort on `focus-visible` and proper tab navigation, which benefit keyboard users and screen readers.

**Will my site look "boring" without transitions?**

It will look faster. Users notice speed far more than they notice a 300ms fade-in. If you want visual interest, focus on typography, spacing, and color—not animation.

**How do I convince designers to drop hover effects?**

Show them mobile analytics. If 60%+ of your traffic is mobile, you're building features for the minority. Then run a Lighthouse audit with and without the effects—the performance difference speaks for itself.


## Summary

Stop building for your desktop monitor; keep your UI clean, minimal, and lightning-fast. In a world of "fancy" technical debt, simplicity is the ultimate competitive advantage.


### About the Author

Maciej Adamski is a software engineer and founder of Dataglitch, specializing in high-performance web applications and frontend optimization. He writes about pragmatic UI development, web performance, and the pursuit of simplicity in code.
