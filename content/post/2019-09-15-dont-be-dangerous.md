---
date: "2019-09-15T00:00:00Z"
image: /img/maverick.jpg
published: true
tags:
- programming
- software development
- software engineering
title: Don't be Dangerous
---

Shipping production software requires a lot of housekeeping; so much so that many developers mentally block out just how
much time a day they spend endlessly shepherding their work through the release lifecycle. Does the pull request have merge
conflicts? Are tests passing? Is the linter happy? Is some other automated check failing?

Most of these failures require intervention; whether it's a finicky spec failing in a test suite or something more
serious like a merge conflict. Intervening on behalf of multiple in-flight features, at a certain point, is a full time
job. I set out a long time ago (maybe more on that in another post) to fix one of these problems because of the
unappreciated amount of risk it poses to releasing software: *merge conflicts*.

If you ask yourself whether you actually have a process for how to resolve merge conflicts at various phases of the
release lifecycle, you probably don't have a good answer. In my opinion there are three critical phases — from least to
most risk — when merge conflicts must be addressed:

1. You're still developing `feature y`, but `feature x` is merged to `master`, conflicting with your branch
2. You've started `feature y`, but `feature x` is in testing with QA. It's partially through testing
3. QA stamped `feature x` "ready to ship" and the world rejoices in unison for your new feature 🚢

Phase 1 is fairly straight-forward. You fix the conflict. Your head is already in the code; you have all the context in
the world. Maybe you have a few questions for the author of `feature x`, but the merge conflict resulting from that
feature being merged is relatively low risk to `feature y`. Maybe `feature x` changes some assumptions you made
going into development, but that sort of risk is standard-issue with a fast-moving or large code base.

Phase 2 is where things become interesting. What does your process say about conflict resolution when features are
partially tested? I can tell you what it should say: everything that has been tested up to the point of conflict _needs
to be retested_. No matter how far-ranging the merge conflict was, if you have a manual QA process, you should believe
in its ability to catch bugs and that means letting it runs its course any time that code changes. That includes bugs
introduced by merge conflict resolution.

If Phase 2 is where things become interesting, Phase 3 is where you'd better have a very precise idea of how to address
merge conflicts. If you think you can do a conflict resolution at this stage and still ship your feature without it
being completely re-tested, I'll assume you're Maverick in Top Gun.

![maverick being dangerous](https://66.media.tumblr.com/747f9516cbd32fea2b32171f91be10d5/tumblr_pv61ptymtc1qhskw9o3_500.gif "That's right, Ice...Man --
I am Dangerous")

Don't be Maverick. Don't be dangerous. I've never seen good developers do worse work than when they're resolving merge
conflicts. I rarely do worse work than when I'm resolving merge conflicts. Merge conflicts are treated like nuisances
that are relatively easy to fix, but the reality is they're one of the greatest opportunities in the software
development lifecycle to introduce bugs. Because so few teams have a well-defined process for how to address merge
conflicts, they're a great opportunity for bugs to jump aboard your ship and cause untold havoc. If you're thinking only
subtle and insidious bugs get introduced during conflict resolutions — I've seen at least one developer ship a _syntax
error_ to production because of a bad conflict resolution during Phase 3.

If you're performing a Phase 3 merge conflict resolution and still think your feature should ship when you hoped it
would, please let me know what product your work on so I can make sure my life never depends on it. You _will_ break
something; it's only a matter of time.

---

I set out to solve this problem with [prwatch-action](https://github.com/marketplace/actions/prwatch-action), but in
doing so, discovered that there's a whole world of automations that can improve developers' lives. Automatic rebases,
automated test pass updates, notifications in the places they're most active...the list goes on.

Currently, `prwatch-action` does one thing: when issues have merge conflicts, they're sent back to a status reflecting
that development work needs to be done. If an issue is "In QA" and has a conflict, suddenly it's "In Progress", and the
author is notified why. If an issue is "Verified by QA", it's suddenly "In Progress" and not going to muck up the
release pipeline when your CI/CD service is unable to create a release branch. Everything is about de-risking merge
conflict resolution by moving issues back to the "in development" phase, and calling attention to conflicts as early as
possible. There's no reason humans should be doing this. Let the robots take the wheel 🤖.

Looking for a project to contribute to? Pull requests welcome for
[prwatch-action](https://github.com/acaloiaro/prwatch-action) over on Github. I'd love to add more notification methods
and issue management systems in addition to Jira.

