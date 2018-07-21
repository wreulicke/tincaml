

# Tincaml


Below example is tincaml file.

```
let f n {
  if n < 2 then n
  else f(n - 2) + f(n -1)
}
f(10)
```

## Try

```bash
$ go run sample
&{f [{n}] [0xc42005c0c0]}
55
```