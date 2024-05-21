```
# Generate templates

templ generate && go run *.go

# Output templates

```

Build production html file:
```
nix develop  .#devShells.build
```

tmux session live-templ generation:
```
nix develop  .#devShells.dev
```
[TODO: would be cool to do the build on templ re-gen, currently needs a manual 'go run *.go' to generate the final html file]


```
 tmux new-session -d -s ms \; \
              split-window -h \; \
              send-keys -t 0 'templ generate --watch' C-m \; \
              send-keys -t 1 'go run *.go' \; \
            attach-session -t ms

tmux kill-session -t ms
```


