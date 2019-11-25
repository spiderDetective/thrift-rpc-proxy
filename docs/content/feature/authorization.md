---
title: "Authorization"
since: "1.5.11"
---

fabio supports basic http authorization on a per-route basis.

<!--more-->

Authorization schemes are configured with the `proxy.auth` option.
You can configure one or multiple schemes.

Each authorization scheme is configured with a list of
key/value options.

    name=<name>;type=<type>;opt=arg;opt[=arg];...

Each scheme must have a **unique name** which is then
referenced in a route configuration.

    proxy.auth = name=myauth;type=...

When you configure the route, you must reference the unique name for the authorization scheme:

    route add svc / https://127.0.0.1:8080 auth=<name>

    urlprefix-/ proto=https auth=<name>

The following types of authorization schemes are available:

* [`basic`](#basic): legacy store for a single TLS and a set of client auth certificates

At the end you also find a list of [examples](#examples).

### Basic

The basic authorization scheme leverages [Http Basic Auth](https://en.wikipedia.org/wiki/Basic_access_authentication) and reads a [htpasswd](https://httpd.apache.org/docs/2.4/misc/password_encryptions.html) file at startup and credentials are cached until the service exits.

The `file` option contains the path to the htpasswd file. The `realm` parameter is optional (default is to use the `name`). The `refresh` option can set the htpasswd file refresh interval. Minimal refresh interval is `1s` to void busy loop. By default refresh is disabled i.e. set to zero.
Note: removing the htpasswd file will cause all requests to fail with HTTP status code 401 (Unauthorized) until the file is restored.

    name=<name>;type=basic;file=<file>;realm=<realm>;refresh=<interval>

Supported htpasswd formats are detailed [here](https://github.com/tg123/go-htpasswd)

#### Examples

    # single basic auth scheme
    name=mybasicauth;type=basic;file=p/creds.htpasswd;

    # single basic auth scheme with refresh interval set to 30 seconds
    name=mybasicauth;type=basic;file=p/creds.htpasswd;refresh=30s

    # basic auth with multiple schemes
    proxy.auth = name=mybasicauth;type=basic;file=p/creds.htpasswd;refresh=30s
                 name=myotherauth;type=basic;file=p/other-creds.htpasswd;realm=myrealm