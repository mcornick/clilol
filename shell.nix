{pkgs ? import <nixpkgs> {}}:
pkgs.mkShell {
  buildInputs = [
    pkgs.cosign
    pkgs.go
    pkgs.go-task
    pkgs.go-tools
    pkgs.gofumpt
    pkgs.golangci-lint
    pkgs.gopls
    pkgs.goreleaser
    pkgs.python312Packages.mkdocs-material
    pkgs.pre-commit
    pkgs.syft
  ];
}
