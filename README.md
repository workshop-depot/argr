# argr
tokenize string as command line args

For example, tokenizes this:

```
s := "-v 1 -s=1 --t 1 --q 11 -w \"tt\\\"tt\" -z=\"tt\\\"tt\" -zz /tmp/dir/"
```

To:

```
[-v 1 -s=1 --t 1 --q 11 -w tt"tt -z=tt"tt -zz /tmp/dir/]
```
