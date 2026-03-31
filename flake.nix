{
  description = "Common Go libraries for cffnpwr";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
    go-overlay = {
      url = "github:purpleclay/go-overlay";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    nixpkgs-extras = {
      url = "github:cffnpwr/nixpkgs-extras";
      inputs = {
        nixpkgs.follows = "nixpkgs";
        go-overlay.follows = "go-overlay";
      };
    };
  };

  outputs =
    inputs@{
      nixpkgs,
      flake-parts,
      go-overlay,
      nixpkgs-extras,
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
        {
          _module.args.pkgs = import nixpkgs {
            inherit system;
            overlays = [
              go-overlay.overlays.default
              nixpkgs-extras.overlays.default
            ];
          };

          formatter = pkgs.treefmt;

          devShells.default =
            let
              miseConfig = fromTOML (builtins.readFile ./mise.toml);

              go = pkgs.go-bin.versions.${miseConfig.tools.go};
              golangci-lint = go.tools.golangci-lint.${miseConfig.tools.golangci-lint};
              gopls = go.tools.gopls.${miseConfig.tools."aqua:golang.org/x/tools/gopls"};
              yamlfmt = pkgs.yamlfmt.versions.${miseConfig.tools.yamlfmt};
            in
            pkgs.mkShell {
              packages = [
                # development tools
                pkgs.git
                gopls

                # linter/formatter
                golangci-lint
                pkgs.nixd
                pkgs.nixfmt
                pkgs.treefmt
                yamlfmt

                # build tools/dependencies
                go
              ];
            };
        };
    };
}
