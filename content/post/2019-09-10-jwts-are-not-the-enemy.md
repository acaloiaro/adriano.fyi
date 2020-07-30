---
date: "2019-09-10T00:00:00Z"
image: /img/no_fud_zone.png
published: true
subtitle: get your FUD off my lawn
tags:
- programming
- software development
- software engineering
title: JWTs Are Not the Enemy
---

I finally felt compelled to write this after reading a [little tidbit](https://www.moesif.com/blog/technical/restful-apis/Authorization-on-RESTful-APIs/#) from a blog posted on Hacker News containing this perennial piece of techno[FUD](https://en.wikipedia.org/wiki/Fear,_uncertainty,_and_doubt)

> One of the downsides with JWTs is that banning users or adding/removing roles is a little harder if you need the action to be immediate.
>
> ...
>
> Since the token is stored client side, there is no way to directly invalidate the token even if you mark the user as disabled in your database. Rather, you must wait until it expires.

Are you sure about that, Derric? Or are you parroting someone else's bad blog post?

---

Let me preface everything I'm about to write by saying that I don't care about [JSON Web Tokens (JWTs)](https://jwt.io). On a couple occasions I've considered using JWTs because they seemed like a simple, yet powerful option for conveying identity between services. But let me be abundantly clear -- I'd be hard-pressed to care any less about JWTs and have never used them for anything approaching "production-worthy". I've started this post at least twice before, but I deleted it because I knew better, but now I don't.

On the first occasion I considered using JWTs, it was for my company's mobile application. I was in the process of rebuilding mobile API authentication and I was researching how other apps in the industry were doing it. I don't consider myself a mobile developer by trade, so I wanted to gain a basic understanding of what "everyone else" was up to. A finger in the wind sort of approach.

On the second occasion, I was building a document conversion service and wanted to base authentication on some sort of "industry standard"; for whatever that's worth. Ultimately, I concluded that JWTs would do little more than introduce complexity to a relatively simple thing. So I scrapped the idea and rolled my own ™️ — and that all turned out to be fine.

However, what struck me throughout my research was just how deeply opinionated people are about JWTs: their security; their ability to "live up to their promises"; and essentially how well they aligned with developers' religious sects of authentication technologies.

There's a bit more juice to the "security" FUD, but it's still FUD because it was based on a [few bad JWT implementations](https://auth0.com/blog/critical-vulnerabilities-in-json-web-token-libraries/) that have all since been fixed. In other words, it's not the JWT spec that's insecure, but how a few people implemented it.

It's the "JWTs can't live up to their promises" variety of FUD that really struck me as odd.

To explain the "promise" that JWTs are supposed to live up to — it's one of statelessness. Statelessness is perhaps JWTs' finest property.

Because JWTs are able to contain within them every bit of information you need to know about users, there's no need to perform a database lookup when a system is presented with one. JWTs are contracts that, at some point, the person or system presenting you with one was considered valid. I.e. that the person (for simplicity) presenting the token *authenticated* successfully and now that person is presenting themselves with their token of authenticity. JWTs should make claims about *who* a person is, not *what* they can do. The difference between knowing who a person is and what that person can do is that of authentication (authn) and authorization (authz).

The argument about whether JWTs can live up to their stateless promise goes, "Well what if I want to disable a user from using my service? I still have to check my database to see whether the user is enabled, so they're not really stateless.". And my argument goes, "No you don't". Does your desire to disable a user in your system change their identity?. No. When disabled users present themselves to your system, they're still the same person; you've simply decided they're less worthy ones. When you decide to disable a user, that's an authorization concern, not identity. You're conflating authn and authz. JWTs never set out to solve the authz problem for you. Were you not going to check your database to see if the user is authorized to do the thing they wanted to do anyway?

I suspect a lot of the people writing FUDdy drivel about JWTs would consider themselves computer scientists, but they
demonstrate all the scientific rigor of a televangelist audience hungry to be bequeathed the words of the almighty from
an ultra-hip lord-whisperer in a pair of [$4,000 sneakers](https://www.instagram.com/p/BvmwrFHBg7V/). Reject using JWTs
for good reasons like they're too big, or they're too expensive to validate given your performance demands, or they're
too complex for your needs; but please cut the FUD. This is a...

![no fud zone](/img/no_fud_zone.png)

