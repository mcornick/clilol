{
  description = "virtual environments";
  inputs = {
    devshell.url = "github:numtide/devshell";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs =
    {
      flake-utils,
      devshell,
      nixpkgs,
      ...
    }:
    flake-utils.lib.eachDefaultSystem (system: {
      devShells.default =
        let
          pkgs = import nixpkgs {
            inherit system;
            overlays = [ devshell.overlays.default ];
          };
        in
        pkgs.devshell.mkShell (
          { extraModulesPath, ... }:
          {
            imports = [ "${extraModulesPath}/language/go.nix" ];
            packages = [
              "go-task"
              "go-tools"
              "gofumpt"
              "golangci-lint"
              "gopls"
            ];
          }
        );
    });
}
