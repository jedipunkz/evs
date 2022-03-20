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
$ evs scan --image scantest:latest --region ap-northeast-1
+----------------+--------+
| SEVERITY LEVEL | COUNTS |
+----------------+--------+
| MEDIUM         |      2 |
| LOW            |     12 |
| INFORMATIONAL  |      4 |
+----------------+--------+
```

## Commiter

jedipunkz

## License

MIT LICENSE

