{pkgs ? import <nixpkgs> {}}:
pkgs.mkShell {
  buildInputs = [
    pkgs.cosign
    pkgs.go
    pkgs.go-task
    pkgs.go-tools
    pkgs.gopls
    pkgs.goreleaser
    pkgs.pre-commit
    pkgs.python312
    pkgs.syft
  ];
}
