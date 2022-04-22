# evs

evs is CLI Tool for scainng AWS ECR Vulnerabilities. 

## Installation

```shell
go install github.com/jedipunkz/evs
```

## Requirement

- go 1.18 or later

## Scan Image

```shell
$ evs scan --image testimage:latest --region ap-northeast-1
+----------------+--------+
| SEVERITY LEVEL | COUNTS |
+----------------+--------+
| MEDIUM         |      2 |
| LOW            |     12 |
| INFORMATIONAL  |      4 |
+----------------+--------+
```

## List Image's Vulnerabilities

```
$ evs list --image testimage:latest --region ap-northeast-1
+------------------+---------------+----------------------------------------------------------------+---------------------------------+
|     CVE NAME     |   SEVERITY    |                              URI                               |           DESCRIPTION           |
+------------------+---------------+----------------------------------------------------------------+---------------------------------+
| CVE-2021-20305   | MEDIUM        | http://people.ubuntu.com/~ubuntu-security/cve/CVE-2021-20305   | A flaw was found in Nettle      |
|                  |               |                                                                | in versions before 3.7.2,       |
|                  |               |                                                                | where several Nettle signature  |
|                  |               |                                                                | verification functions          |
|                  |               |                                                                | (GOST DSA, EDDSA & ECDSA)       |

<snip>

| CVE-2020-14155   | INFORMATIONAL | http://people.ubuntu.com/~ubuntu-security/cve/CVE-2020-14155   | libpcre in PCRE before 8.44     |
|                  |               |                                                                | allows an integer overflow      |
|                  |               |                                                                | via a large number after a (?C  |
|                  |               |                                                                | substring.                      |
| CVE-2017-11164   | INFORMATIONAL | http://people.ubuntu.com/~ubuntu-security/cve/CVE-2017-11164   | In PCRE 8.41, the OP_KETRMAX    |
|                  |               |                                                                | feature in the match function   |
|                  |               |                                                                | in pcre_exec.c allows stack     |
|                  |               |                                                                | exhaustion (uncontrolled        |
|                  |               |                                                                | recursion) when processing a    |
|                  |               |                                                                | crafted regular expression.     |
+------------------+---------------+----------------------------------------------------------------+---------------------------------+
```

## Credits

Tomokazu HIRAI (@jedipunkz)


## License
The source code is licensed MIT. The website content is licensed CC BY 4.0,see LICENSE.
