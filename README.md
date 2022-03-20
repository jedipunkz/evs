# evs

evs is CLI Tool for scainng AWS ECR Vulnerabilities. 

## scan image

```shell
$ evs scan --image scantest:latest --region ap-northeast-1
ImageName: scantest, Tag: latest
+----------------+--------+
| SEVERITY LEVEL | COUNTS |
+----------------+--------+
| MEDIUM         |      2 |
| LOW            |     12 |
| INFORMATIONAL  |      4 |
+----------------+--------+
```

## Commiter

- jedipunkz

## LICENSE

MIT LICENSE

