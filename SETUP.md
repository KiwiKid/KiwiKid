```
# Generate templates
templ generate && go run *.go

# Output templates

```


```
tmux new-session -d -s KiwiKidDev \; \
  split-window -h \; \
  send-keys -t 0 'templ generate --watch' C-m \; \
  send-keys -t 1 'go run *.go' C-m \; \
  attach-session -t KiwiKidDev

firefox "file://$HOME/src/local/KiwiKid/docs/index.html"

tmux attach-session -t KiwiKidDev

tmux kill-session -t KiwiKidDev
```
