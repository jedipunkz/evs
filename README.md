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
```

## Credits

Tomokazu HIRAI (@jedipunkz)


## License
The source code is licensed MIT. The website content is licensed CC BY 4.0,see LICENSE.
