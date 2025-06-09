{pkgs ? import <nixpkgs> {}}:
pkgs.mkShell {
  buildInputs = with pkgs; [
    go
    go-task
    go-tools
    gofumpt
    golangci-lint
    gopls
    universal-ctags
  ];
}
