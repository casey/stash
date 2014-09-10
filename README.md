WORM
----

A write once read many key value storage service.

[Test instance here.](http://rodarmor-worm.appspot.com)

KEYs and VALUEs match `/[a-zA-Z0-9.-_]*/`.

KEYs can be any length, but VALUEs are limited to 64 characters just to avoid too much spam in the test instance datastore. The allowed characters are exactly enough for URL-safe base64 encoding.

PUT /KEY/VALUE associates KEY with VALUE. Puts after the first will be ignored.

GET /KEY will return the value associated with KEY.

Unfortunately, due to app engien limitations, PUTs and GETs with an empty key, `//VALUE`, don't work.
