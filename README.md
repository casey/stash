WORM
====

A write once read many key value storage service.

[Test instance here.](http://rodarmor-worm.appspot.com)


API
---

PUT /KEY/VALUE associates KEY with VALUE. Puts after the first will be ignored.

GET /KEY will return the value associated with KEY.

KEYs match `/[a-zA-Z0-9.-_]+/`.
VALUEs match `/[a-zA-Z0-9.-_]{0,64}/`.

Errata
------

KEYs can be any length, but VALUEs are limited to 64 characters just to avoid too much spam in the test instance datastore. The allowed characters are exactly enough for URL-safe base64 encoding.

Unfortunately, due to app engine limitations, PUTs and GETs with an empty key, `//VALUE`, don't work.

The sha256 hash of KEYs are used as datastore string IDs, instead of the KEY itself. This allows keys to be arbitrarily long, since they aren't actually stored in the datastore. Also, this lessens the severity of a [potential attack vector](http://ikaisays.com/2011/01/25/app-engine-datastore-tip-monotonically-increasing-values-are-bad/).
