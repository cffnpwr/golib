{
  description = "Common Go libraries for cffnpwr";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
    go-overlay = {
      url = "github:purpleclay/go-overlay";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs =
    inputs@{
      nixpkgs,
      flake-parts,
      go-overlay,
      ...
    }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [
        "x86_64-linux"
        "aarch64-linux"
        "x86_64-darwin"
        "aarch64-darwin"
      ];

      perSystem =
        { pkgs, system, ... }:
        let
          go = pkgs.go-bin.fromGoMod ./go.mod;
        in
        {
          _module.args.pkgs = import nixpkgs {
            inherit system;
            overlays = [ go-overlay.overlays.default ];
          };

          devShells.default = pkgs.mkShell {
            packages = [
              (go.withTools [
                "golangci-lint"
                "gopls"
              ])
            ];
          };
        };
    };
}
