# go-authcrunch-secrets-static-secrets-manager

<a href="https://github.com/greenpau/go-authcrunch-secrets-static-secrets-manager/actions/" target="_blank"><img src="https://github.com/greenpau/go-authcrunch-secrets-static-secrets-manager/workflows/build/badge.svg?branch=main"></a>
<a href="https://pkg.go.dev/github.com/greenpau/go-authcrunch-secrets-static-secrets-manager" target="_blank"><img src="https://img.shields.io/badge/godoc-reference-blue.svg"></a>

[AuthCrunch](https://github.com/greenpau/go-authcrunch) Secrets Plugin
for statically configured secrets.

<!-- begin-markdown-toc -->
## Table of Contents

* [Getting Started](#getting-started)
  * [Password Hashing](#password-hashing)
  * [Secrets Management](#secrets-management)
    * [User Credentials](#user-credentials)
* [AuthCrunch Usage](#authcrunch-usage)

<!-- end-markdown-toc -->

## Getting Started

### Password Hashing

Install `bcrypt-cli` for password hashing:

```bash
go install github.com/bitnami/bcrypt-cli@latest
```

Install `pwgen` for password generation:

```bash
sudo yum -y install pwgen
```

Generate a password:

```
$ pwgen -cnvB1 32
rbrH97m9bpbk3qRphHFNM9ksJfRcWdvr
```

Next, hash the `password` and the `api_key`:

```
$ echo -n "rbrH97m9bpbk3qRphHFNM9ksJfRcWdvr" | bcrypt-cli -c 10
$2a$10$iqq53VjdCwknBSBrnyLd9OH1Mfh6kqPezMMy6h6F41iLdVDkj13I6
```

Repeat the same thing for `api_key`.

```
$ pwgen -cnvB1 32
kqvc7cgk44dtpX9nXx4NL9krH4g7fqdJ
$ echo -n "kqvc7cgk44dtpX9nXx4NL9krH4g7fqdJ" | bcrypt-cli -c 10
$2a$10$TEQ7ZG9cAdWwhQK36orCGOlokqQA55ddE0WEsl00oLZh567okdcZ6
```

### Secrets Management

#### User Credentials

Create a set of credentials for a management user, `jsmith` and use it in
configuring a secret:

* `username`: `jsmith`
* `password`: `bcrypt:10:$2a$10$iqq53VjdCwknBSBrnyLd9OH1Mfh6kqPezMMy6h6F41iLdVDkj13I6`
* `api_key`: `bcrypt:10:$2a$10$TEQ7ZG9cAdWwhQK36orCGOlokqQA55ddE0WEsl00oLZh567okdcZ6`
* `email`: `jsmith@localhost.localdomain`
* `name`: 'John Smith`

## AuthCrunch Usage

TODO.

