# Mnemonic-as-a-Service using Golang and Buffalo!

## The basic code behind mnemonic-as-a-service.com

>There are only two hard things in Computer Science: cache invalidation and naming things.
>
>-- Phil Karlton

The nice folks at [mnx.io](http://mnx.io/) wrote a fun [blog post](http://mnx.io/blog/a-proper-server-naming-scheme/) about server naming schemes.

One of the pearls of wisdom is the recommendation to use Oren Tirosh’s [mnemonic encoding project](http://web.archive.org/web/20090918202746/http://tothink.com/mnemonic/wordlist.html) as a wordlist.

GENIUS!

Since I originally started playing with this wordlist, it has become my ```hello world``` example for dealing with new languages and frameworks.  I have examples in [flask](https://github.com/alexlovelltroy/mnemonic-as-a-service), [django](https://github.com/alexlovelltroy/django_mnemonic), nodejs, and lambda, that have all been behind menmonic-as-a-service.com at one time or another.

# Enter Golang, Docker and Buffalo

Enough time has passed that I needed to play with a few more ideas.  The first is golang, so that's the language I've used here.  The second is Docker which is a natural fit for running Go binaries with repeatable environment variables.  It's certianly not necessary, but fun to learn.  I also used this opportunity to try out [Buffalo](http://gobuffalo.io) which is only subtly different than my [last Go attempt](https://github.com/alexlovelltroy/mnemonic-gin) using [Gin](https://gin-gonic.github.io/gin/).  

[Powered by Buffalo](http://gobuffalo.io)
