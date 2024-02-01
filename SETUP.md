```
# Generate templates
templ generate && go run *.go

# Output templates

```


```
tmux new-session -d -s KiwiKidDev \; \
  split-window -h \; \
  send-keys -t 0 'templ generate --watch' C-m \; \
  send-keys -t 1 'go run *.go' C-m


tmux attach-session -t KiwiKidDev
```
