---
title: "proxy.auth"
---

`proxy.auth` configures one or more authorization schemes.

Each authorization scheme is configured with a list of
key/value options. Each scheme must have a unique
name which can then be referred to in a routing rule.

    name=<name>;type=<type>;opt=arg;opt[=arg];...

The following types of authorization schemes are available:

#### Basic

The basic authorization scheme leverages [Http Basic Auth](https://en.wikipedia.org/wiki/Basic_access_authentication) and reads a [htpasswd](https://httpd.apache.org/docs/2.4/misc/password_encryptions.html) file at startup and credentials are cached until the service exits.

The `file` option contains the path to the htpasswd file. The `realm` parameter is optional (default is to use the `name`). The `refresh` option can set the htpasswd file refresh interval. Minimal refresh interval is `1s` to void busy loop. By default refresh is disabled i.e. set to zero.
Note: removing the htpasswd file will cause all requests to fail with HTTP status code 401 (Unauthorized) until the file is restored.

    name=<name>;type=basic;file=<file>;realm=<realm>;refresh=<interval>

Supported htpasswd formats are detailed [here](https://github.com/tg123/go-htpasswd)

#### Examples

    # single basic auth scheme
    name=mybasicauth;type=basic;file=p/creds.file;

    # single basic auth scheme with refresh interval set to 30 seconds
    name=mybasicauth;type=basic;file=p/creds.htpasswd;refresh=30s

    # basic auth with multiple schemes
    proxy.auth = name=mybasicauth;type=basic;file=p/creds.htpasswd;refresh=30s
                 name=myotherauth;type=basic;file=p/other-creds.htpasswd;realm=myrealm

The default is

    proxy.auth =
