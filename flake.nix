{
  description = "A flake for building and running a Go program with devShells";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-23.05";
    nixpkgs-unstable.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs, nixpkgs-unstable }: 
    let 
      system = "x86_64-linux";
      pkgs = import nixpkgs { inherit system; };
      unstablePkgs = import nixpkgs-unstable { inherit system; };
      go = unstablePkgs.go_1_21;
      templ = unstablePkgs.templ;
    in
    {
      devShells = {
        build = pkgs.mkShell {
          buildInputs = [
            go
           templ
          ];
          shellHook = ''
            echo "Building the Go project..."
            git config user.name $GIT_AUTHOR_USER
            git config user.email $GIT_AUTHOR_EMAIL
           templ generate && go run *.go
          '';
        };

        dev = pkgs.mkShell {
          buildInputs = [ 
            pkgs.air
            templ
            go 
          ];
          shellHook = ''
            git config user.name $GIT_AUTHOR_USER
            git config user.email $GIT_AUTHOR_EMAIL
            tmux new-session -d -s ms \; \
              split-window -h \; \
              send-keys -t 0 'templ generate --watch' C-m \; \
              send-keys -t 1 'go run *.go' \; \
            attach-session -t ms
          '';
          shellExit = ''
            tmux kill-session -t ms
          '';
        };
      };
    };
}


