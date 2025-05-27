{
  pkgs ? import <nixpkgs> { },
}:
pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.go-task
    pkgs.go-tools
    pkgs.gopls
  ];
}
