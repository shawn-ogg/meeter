# meeter

This is a simple redirect server for the awesome meet.jit.si video conference service.

## What problem does it solve?

Every meet.jit.si/ URI is a conference room. The problem is, that the good ones are already taken or the likelyhood of being disturbed by a
stranger is very high. Using a password is also another concern you have to manage.
Meeter alows you to host a redirect server which alows you to use nice and short names e.g. `meet.myhost.com/pubquiz`.
It will redirect to a random `meet.jit.si/name+uuid` url. The url will change by default once a day. 
