{
  description = "didweb";

  inputs = {
    nixpkgs-unstable.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, flake-utils, ... }@inputs : flake-utils.lib.eachDefaultSystem(system: let
    pkgs = import inputs.nixpkgs-unstable { inherit system; };
  in {
    packages.default = pkgs.buildGoModule {
      pname = "didweb";
      version = "0.1";
      # In 'nix develop', we don't need a copy of the source tree
      # in the Nix store.
      src = ./.;

      preBuild = ''
        export GOWORK=off
      '';

      vendorHash = "sha256-nen6MZ1MQ18ZvNA6gyl0jCPLkYfu3OGIl5n+qRcbuPo=";
    };
    devShells.default = pkgs.mkShell {
      packages = with pkgs; [
        hello
      ];
    };
  });
}
