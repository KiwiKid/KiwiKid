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
          buildInputs = [ go ];
          shellHook = ''
            echo "Building the Go project..."
            go build -o ./tmp/main .
          '';
        };

        dev = pkgs.mkShell {
          buildInputs = [ 
            pkgs.air
            templ
            go 
          ];
          shellHook = ''
            echo "firefox /home/greg/mine/KiwiKid/docs/index.html to preview"
            templ generate --watch

          '';
        };
      };
    };
}


