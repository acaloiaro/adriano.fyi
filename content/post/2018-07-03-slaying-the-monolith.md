---
date: "2018-07-03T00:00:00Z"
image: /img/negative_man.png
published: true
subtitle: Documents @ Greenhouse
tags:
- Golang
- Libreoffice
- Software Development
- Web Development
- Technology
title: Slaying The Monolith
---
This is the first in what I hope becomes a series of posts on how we’re slaying the monolith at Greenhouse. Over the course of the next year, Greenhouse Engineering is aiming to break down our monolithic Rails application into a more scalable and robust collection of services. This won’t happen quickly, but it will happen deliberately and with the clear goal of safely scaling Greenhouse Recruiting into the future.

This is a cross-post from Greenhouse's Engineering blog, _[In the Weeds](https://medium.com/in-the-weeds/slaying-the-monolith-3b7a017faf02)_.

---

# Documents On Documents
It probably goes without saying that Greenhouse Recruiting handles a lot of documents. Every day 75,000 people apply to jobs through Greenhouse Recruiting. Nearly every one of them has a resume and cover letter and every document must be converted to a format compatible with Greenhouse’s desktop and mobile applications. On top of that, every document gets thumbnailed and scanned for viruses and malware.

Let’s get this out of the way — thumbnails are the easy part. And you might be wondering why documents need to be converted at all. Why not leave them in their original format? Well, because for better or worse, a large percentage of every resume ever authored since the 1980s is in a format that we all know and love — Microsoft Word — a format that every browser neither knows nor loves.

Browsers can’t display Word documents natively. It’s possible to render Word documents in browser plugins, but it’s simply not acceptable to build a great product like Greenhouse Recruiting and ask users to install a browser plugin to use one of its core features: application review. And boy do customers take application review seriously; whether on desktop or mobile, it’s a critical part of the hiring process, and it must work flawlessly on every platform. Naturally, cover letters and resumes are a big part of application review and we needed a way for recruiters and hiring managers to review them as painlessly as possible. For years we’ve had an internal microservice that scans documents for viruses and malware, but we’ve relied on a separate third party service to perform document conversion, thumbnail generation, and previewing … until now.

# How We Arrived Here
Technical decisions at early stage startups are usually born out of absolute necessity and with very little room for failure. Decisions made at the six month mark make a lot less sense two or more years later. So when we decided to outsource document conversion and previewing at Greenhouse, it made a lot of sense — it was the right thing to do. The cost was nominal and no Greenhouse engineers had to maintain or develop it; they could focus on our core product. Big wins all around.

But when third-party services start costing an engineers’ salary every year, at the very least you should begin questioning whether your rationale for having purchased the service still holds. As a more mature company with the good fortune of having built a product that our customers love, Greenhouse is obligated to make the sort of strategic decisions that optimize long-term sustainability over the short-term fight for survival. It was time to bring documents in house.

# Dochouse is Born
We began this project with a few guiding principles:

1. Previewing documents could work better, but not worse than with our 3rd party provider or else recruiters would be rioting in the streets.
2. Word documents had to be converted to some other format.
3. Some other format would probably be PDF because PDFs are portable, right? (spoiler alert: they are)

Early on we mulled over the idea of bolting Dochouse onto our Rails monolith. It’s never easy dropping the habit of throwing everything onto that familiar pile of code you’ve worked on the last X years. It’s comfortable there. The problem is, that pile of code is a productivity-sucking monster. It’s where “move fast and break things” goes to die a very slow, and very miserable death. The monolith moves slow and breaks your spirit.

So we decided on a two-pronged solution. Greenhouse Recruiting would get a shiny new PDF.js viewer that lives on our CDN, and some other new and shiny thing would start converting documents from Microsoft Word, TXT, and RTF to PDF. In other words, PDF was becoming la lingua franca for documents.

The commercial document conversion space is full of capable products, but we don’t like black boxes at Greenhouse and we certainly don’t like swapping out one black box as a service (BBaS) for another black box on premises (BBoP — let’s make these acronyms happen 🤞). Anyone who has done work in this space can probably tell where this is headed by now: LibreOffice and open source software to the rescue (OSS >BBaS && BBoP )! You may be thinking that LibreOffice is what your company’s resident neckbeards use to annoy business and IT folks with their wonky document formats, and you’re right, but it’s also what powers Collabora’s fantastic productivity suite and document conversion service. Indeed, Collabora’s Michael Meeks provided critical insights as we navigated the waters of running LibreOffice at scale (more on that shortly!).

The remainder of the stack consists of a simple Golang web server wrapping LibreOffice for conversion and Libvips for thumbnail generation. Libvips is fast, stable, and more importantly doesn’t throw temper tantrums when applicants upload corrupt or password-protected PDF files (yes, that happens. often). Finally, when documents are converted and thumbnails generated, everything gets dumped into a secure Amazon S3 bucket. This all sounds great and ready to ship; what could go wrong?

# The Rollout
Like any self-respecting SaaS company, we obsessively use our own products in house and we maintain fine-grained control over feature rollouts. And a feature like this is not something that gets rolled out to every customer at once, and certainly not in a way that users are able to see. So our zero-impact rollout strategy looked something like this (try to guess the phase where things began to catastrophically fail):

1. Begin rolling Dochouse out to our own company’s recruiting team; in parallel with the existing system
2. Include small customers
3. Include small and medium customers
4. Include small, medium, and enterprise customers

If you guessed that it failed in phase 1, you get an A+. After less than a day of processing documents for Greenhouse Software’s instance of Greenhouse Recruiting, the system ground to a screeching halt. No documents were processing, CPU was pegged at 100%, memory consumption was through the roof—turnkey black box solutions had a sudden allure.

Let me preface this by saying LibreOffice is an absolute triumph of the open source movement. Most users never have the opportunity of seeing it fail because 1. it’s good, and 2. most users aren’t opening tens of thousands of documents in it every day. Any time you use software in unexpected ways, failure should be the expected outcome. Similarly, with any software that handles user input, failure should be the expected outcome.

![Occasionally LibreOffice can't event ](/img/negative_man.png)


Because people upload weird things — text documents with Word document extensions (.doc/.docx). Word documents with text extensions (.txt). Clips of The Bachelor with rich text format extensions (.rtf). Your garden variety weird stuff; failures are bound to happen. The most insidious, though, is a Word document with the correct file extension; a file to which LibreOffice promptly responds, “Ok, I’ve got this” — followed by several minutes of silence, and finally — “I can’t even. I’ve lost my capacity to even with this document”. These documents may be valid in Microsoft Word, but they’re LibreOffice kryptonite. This is where Dochouse was on day one — with significantly less than one percent of real-world load, Dochouse couldn’t even. We needed to get a handle on LibreOffice processes; they were running amok with neither time nor resource bounds. LibreOffice is persistent; when it encounters documents that it can’t convert, it doesn’t quit. Which, for Dochouse, meant that goroutines and memory consumption were on a monotonic path to abject failure. A full production rollout was absolutely out of the question.

Collabora had already solved this problem by incorporating JODConverter into their products. JODConverter manages a pool of Libre/Open Office instances; it detects when instances get stuck and evacuates them from the worker pool. Additionally it adds work deadlines around every job. This is exactly what we needed. The problem is that JODConverter is written in Java and Java isn’t nearly as hip as Golang. I mean, Golang has a cute gopher logo and Java has what … Duke? Please.

![Golang Gopher](/img/gopher.png)

Case closed.

Poor humor aside, we love Golang at Greenhouse for legitimate reasons. Go code is beautifully simple, and channels/CSP are a delight to work with compared to other concurrency models. We knew Dochouse would need a high degree of concurrency, we just weren’t sure to what extent (turns out in production Dochouse averages ~70 concurrent goroutines); so Golang was a natural choice. Personally, I don’t consider myself an accomplished Go developer. I fall somewhere between novice and “pretty alright”, but that doesn’t change how I feel about the language.

Armed with the universal human desire not to maintain Go code that contains embedded Java/JODConverter (`go → cgo →  JNI →  Java = NOPE!`)and the knowledge that a simple process barrier between LibreOffice and the Golang web server would solve all of our woes — we built a minimal process manager and worker pool in Go and exposed it locally over a tcp/ip RPC interface. The process manager assigns two things to every worker process in the pool.

A status change channel to convey changes in worker status to listeners: e.g. missed work deadline notifications, process terminations.
A monitor goroutine that waits for events on the status change channel and takes action. The most common action is restarting workers that are in a bad state.
Most importantly, what this process barrier gives us is the critical ability to kill unhealthy LibreOffice workers and assign deadlines to RPC calls. In short, it’s what prevents a single bad document conversion from dragging the entire system down with it, allowing us to scale from tens of documents per day to tens of thousands.

If we built this on our Rails monolith, we would have been in for a world of hurt. What would have happened had our production application been directly exposed to runaway LibreOffice processes? What would the development and debugging lifecycle look like had this been done inside of a Sidekiq worker? Would our infrastructure and security teams be comfortable shipping LibreOffice in the monolith Docker image? I feel crippled by indecision even asking these questions hypothetically. I’d rather be building things and making real progress.

Interested in working on projects like Dochouse that impact thousands of users every day? [Greenhouse is hiring](https://grnh.se/716d31af1) Software Engineers. Join us!
