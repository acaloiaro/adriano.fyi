---
date: "2020-10-21"
published: true
title: My Wireless Internet Setup
image: https://nextcloud.adriano.fyi/s/PdSy96KmXjPiDt6/preview
share_img: https://nextcloud.adriano.fyi/s/PdSy96KmXjPiDt6/preview
tags:
- roaming
- travel
- personal
---

If you're reading this, you're probably familiar with the fact that I'm [permanently roaming](https://adriano.fyi/post/2020-08-26-roaming/), and I'd be foolish not to acknowledge the importance of my wireless setup in making that possible. Getting -- and staying -- connected while permanently roaming is nothing like paying the local cable cabal to get your house or apartment connected.

Remember the last show you binged on Netflix over the weekend? That consumed 20x most mobile users' monthly data allotment, but you never had to think about data because when you're wired you generally don't pay for consumption. How many hours of video conferencing do you do every day? That'll be ~1GB/hr, please.

Let me make one thing clear from the outset: I'm by no means an expert in wireless technology, and there may be *many* better ways to get connected than the one(s) I've chosen. So if you spot a flaw or have suggestions for me, be sure to get in touch and let me know. I'll keep this post up to date as reasonably possible as I get more feedback and make changes in the future.

I also feel compelled to call out **"Unlimited Data"** plans before we get started. For us mere mortals who are not grandfathered into legacy plans: unlimited data plans no longer exist. You may find a fly-by-night company offering unlimited data plans for a short time, but you'll probably soon realize that they're doing it in violation of their agreement with the upstream provider when they get shut down. In the best case scenario, you'll need to get a new SIM card if they're able to remain in operation and switch to a new upstream provider. These fly-by-nights are more commonly known as MVNOs (Mobile Virtual Network Operators). They're the likes of Cricket Wireless, Boost Mobile, etc. These companies are piggybacked on AT&T, Verizon, T-Mobile, and Sprint. As a result, they're subject to the true operator's rules and the terms of their agreement.

Figuring all this out was a bit of a journey for me and I've had a ton of conversations with friends and acquaintances about the details. So I figured it was time to finally capture the details.

---

# Choices

Between equipment and service providers, there's no shortage of choices to be made early in the wireless setup journey. I'll try to capture a non-exhaustive list of pros and cons for a few options before getting into the choices I've personally made, and the motivation for having made them. However, be sure to consider your data needs and access patterns before taking my opinions and choices as gospel, because if you won't be working or streaming media -- your choice might be quite simple: **cell phone hotspot**.

# Equipment

Equipment is the most technical aspect of getting yourself connected while roaming. For the time being, I'll forego talking at all about dedicated antennas, because that's a topic that is both very deep, and one I cannot speak very intelligently about. However, if you want a serious setup, you'll want to buy a dedicated antenna; trust me.

## Cell Phone Hotspot

This is the most obvious choice if you're just dipping your toes in the roaming waters. It doesn't require any equipment or service that you don't already have, and is actually quite fast and versatile. If cell phone antennas were as good as external, dedicated antennas, and the wireless cabals would stop pretending that anyone who wants to hotspot from their phone is some kind deviant, phone hotspots would be a fantastic option. And for many people, they still can be.

### Pros

* No new equipment or service required
* Cheap

### Cons

* Cell phone antennas don't have enough gain to pick up the weak radio signals that whisper across BLM (Bureau of Land management), National Forest, and other public lands. "No bars" will quickly render those lands uninhabitable.
* The wireless cabals often place very low "Hotspot" data caps on their "Unlimited" plans. While on-device data may be relatively unlimited, you'll find that hotspot data is capped at 5-15GB.

## Signal Boosters

Signal boosters can be great, especially if you're using your cell phone as a hotspot. Boosters can range from tens of dollars to several hundred. These can be thought of as cell signal re-transmitters; they make a weak signal stronger -- which is why they're a good complement to a cell phone hotspot. I've never owned one, so be aware that my opinion is of limited value here.

### Pros

* Usually cheaper than dedicated wireless routers.
* Can be used in conjunction with your existing equipment (cell phone)

### Cons

* Don't provide connectivity to non-cell devices (laptops, tablets, smart TVs, etc.)
* The range over which they amplify cell signal can be quite small. Sometimes less than a 10 feet.
* Not all allow for external antennas, so their gain (ability to pick up signal) and placement can be limited.

## Wireless Routers

Wireless routers are the most versatile option. And generally the most expensive.

### Pros

* Usually provide a WiFi network, allowing you to bring non-cell devices online.
* Have external antennas, allowing for optimal placement/gain.
* Can be purchased directly from your wireless provider or 3rd party.

### Cons

* Many wireless providers only _officially_ support their own "hotspot" devices. E.g. "Nighthawks" and "MIFIs"
* Can be quite expensive

# Wireless Providers

This section could become _very long_. So instead of providing all the details here, I'll point you to the [Mobile Internet Resource Center](https://www.rvmobileinternet.com), which is an invaluable resource for understanding the current mobile service provider landscape. Things change _fast_ when it comes to providers, so by the time you read this there may be a new kid on the block. I used Mobile Internet Resource Center extensively when I was researching both gear and providers.

## Quick Summary

Instead of getting into the details here, I'll leave you with the most important thing that you should take away:

**Redundant service providers is essential**

I cannot stress that enough. If you can afford it, and you know you'll move locations often -- have at lease two service providers at all times. Whether it's Verzion, Sprint, T-Mobile, or AT&T -- have at least two. If you can afford more, get more than two; especially if you don't have a direct relationship with the provider and one is at risk of going out of business (i.e. you're connecting through an MVNO).

 ## Throttling and Deprioritization

You'll **need** to know the difference between throttling and deprioritization before purchasing a data plan. Every wireless provider engages in one, if not both of these two practices, and the difference between them are stark.

## Throttling

Providers may refer to throttling by some other name, but the details are generally the same.

Scenario:

You have an "unlimited" data plan with some fine print that says after `X` gigabytes of usage, you'll be throttled to `2G` speed.

You're half way through the month and you've already consumed `X` gigabytes of data. Now it's important to know how fast "2G speed" is. Almost all providers throttle to 2G speed when they throttle, so make no mistake: **2G SPEED IS WAY TOO SLOW**

You're not getting anything done at `2G` speed. Specifically, it's `128Kbps`, or an ISDN line, or `~2x` the speed of your childhood AOL dialup connection. Scientifically, it's the speed at which you test a web page in your browser and you give up saying: "Nope, no internet here". So if you're getting throttled, you're neither playing nor working on the interwebs. And let me be clear, if you're getting throttled to `128Kbps/2G` speed, you're _lucky_ if you see that kind of speed. You're probably running at 1995 `56K` dialup at best.

Throttling to `2G` speed on a `4G` data plan should be _illegal_. That's not hyperbole.

## Deprioritization

This is the lesser of two evils.

Scenario:

You have an "unlimited" data plan with some fine print that says after `X` gigabytes of usage, you'll be deprioritized.

Someone who hasn't exceeded their soft cap needs top download a file and the tower is already overloaded. Your Netflix stream get a little stuttery.

This one's a bit more ambiguous. But in my experience, and in the places I tend to stay, it's hard to tell the difference between mediocre signal and deprioritization. Basically, your traffic gets put behind everyone else's traffic. If you're somewhere rural, you're probably not going to notice because service already sucks. That's the reality.

---

# My Setup

If you look at the pros/cons above, you'll probably have guessed my choice. I decided to purchase a LTE Router with an omnidirectional antenna, and service provided by *Verizon* and *AT&T*.

# LTE Router

I'm specifically calling out that I have a _LTE_ router and not 5G because in 2020 it's relevant to note that we're on the verge of a nationwide 5G rollout. There are two reasons I chose 4G LTE:

1. 5G is far from rolled out nation-wide and coverage in rural areas hardly exists.
2. 5G in the high frequencies does not penetrate obstacles or refract well. Making it non-ideal for hilly and wooded areas.

My router is a [Sierra Wireless RV55](https://www.sierrawireless.com/products-and-solutions/routers-gateways/rv55/). It's not cheap, but it's reliable and rugged as hell.

Purchase Motivation:

* *It's LTE-A Pro (CAT 12)*, supporting up to 600Mbps/150Mbps (DL/UL) [But let's be real, you'll never see that kind of speed on _any_ provider. I think I wasted my money on this feature]
* *Supports Ethernet for when it's available*
* *Supports dual-mode WiFi.* I have two 9dBi WiFi antennas attached to it -- one can pick up WiFi, and the other can transmit it. This way, my Kindle, Apple TV, laptops, etc. don't need to be reconfigured to use a new WiFi provider if I'm on campground WiFi. It sees the new WiFi connection as one they're already familiar with: `roam`.
* *Has two SIM slots.* This means I can switch between Verizon and AT&T automatically without having to replace any SIM cards. Granted, network switching is slow because it has to reload network-specific firmware any time it switches network.
* *Supports both AC and DC power.* Living in an RV this is huge. The router is wired directly into the RV's 12V DC power, drawing less than an amp of power. Not running a DC->AC inverter is key to saving battery.

# Antenna(s)

I actually have both directional (9 dBi) and omnidirectional (6 dBi) antennas. Now you've seen me talk about both _gain_ and this _dBi_ nonsense. Gain is measured in dBi. The higher the number, the more signal you can grab out of the air. 6 dBi is decent, 9 dBi is pretty good. A good "Yagi" antenna gain can be > 10 dBi.

I use my omnidirectional antenna almost 100% of the time because I generally have no idea where the tower is and would rather not spend 20 minutes pointing my directional antenna all over creation trying to find a tower when I arrive somewhere new. Not when there's only 3 dBi difference between my omni/directional antennas. If 3 dBi is the difference between me being connected and disconnected -- chances are, I'm not going to get much done on that connection anyway. It's _going_ to be volatile. Temperature changes will affect it, humidity will affect it; maybe even star alignment. I don't know.

My antenna is a [Proxicast Pro-Gain 4G/LTE MIMO Antenna](https://www.amazon.com/gp/product/B07ZBVX8VL/). That this antenna is a MIMO antenna is what helps me take advantage of the LTE-A tech in my router. Am I getting 600Mbps download speed? Not a chance! But it does allow the router to bond together the signal from two different channels into a more reliable connection. Neat.

# Signal Tester

I'll come back to this one. I have a signal tester, but I'm not looking at it and can't remember what brand/model it is. These are great if you want to do some recon in a new area. Bring your signal tester and the antenna that you use with your device to get an idea of whether or not the signal is viable. I won't bet my life on the results I get from my tester, but it's a good gut check if I don't want to pull the rig into a new location blind.

# Wireless Providers

As I already let on earlier, I have two redundant providers: AT&T is my primary and Verizon is my backup. At one point, I had T-Mobile provided by a MVNO called OTR Mobile. That gravy train was great while it lasted: $80/mo for a truly unlimited plan. It lasted all of four months before (I'm guessing), T-Mobile terminated the relationship for violation of contract.

## AT&T

My AT&T service is a bit unique, and I acknowledge that this option is not available to everyone because it requires one to have either a business of their own, or a business willing to open a line on their behalf.

My AT&T service is a "unlimited" 12Mbps Mobile Broadband plan for ~$150/mo (after taxes). This is as near unlimited as I believe possible today, as my data is simply "deprioritized" after 75GB. I've not been able to notice deprioritization. See the [Throttling & Deprioritization](#throttling-and-deprioritization) for more details on what this means.

## Verizon

My backup Verizon service is through [US Mobile](https://www.usmobile.com). It's "Unlimited", but they can throttle after 50GB. Granted, I've never noticed any throttling because it's always fairly slow. It's reliable though, and acceptably slow.

---

# Conclusions

I've glossed over _a lot_ of details and yet this post is still too long. My apologies. If there's anything I can help clarify, or details you think I may be able to reveal, please feel free to reach out to me. I can probably feign some degree of competence in providing answers. Also, please let me know if I've gotten anything wrong. I researched a lot of service providers and equipment throughout this process, so if there's a router, booster, or providers you have questions about -- I may know a thing or two.

Cheers

