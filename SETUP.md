```
# Generate templates

templ generate && go run *.go

# Output templates

```

Or via Nix:

```
nix develop  .#devShells.build
```
```
nix develop  .#devShells.dev
```

```
tmux new-session -d -s KiwiKidDev \; \
  split-window -h \; \
  send-keys -t 0 'templ generate --watch' C-m \; \
  attach-session -t KiwiKidDev


tmux attach-session -t KiwiKidDev

tmux kill-session -t KiwiKidDev
```
