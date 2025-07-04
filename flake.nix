{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05";
    utils.url = "github:numtide/flake-utils";
  };
  outputs =
    {
      nixpkgs,
      utils,
      ...
    }:
    utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShell =
          with pkgs;
          mkShell {
            buildInputs = with pkgs; [
              go
              go-task
              go-tools
              gofumpt
              golangci-lint
              gopls
              gosec
              pre-commit
              universal-ctags
            ];
          };
      }
    );
}
