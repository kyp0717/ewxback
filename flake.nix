{
  description = "A basic gomod2nix flake";
  inputs.ewxfront.url = "github:kyp0717/ewxfront?ref=test-react";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";
  inputs.gomod2nix.url = "github:nix-community/gomod2nix";
  inputs.gomod2nix.inputs.nixpkgs.follows = "nixpkgs";
  inputs.gomod2nix.inputs.flake-utils.follows = "flake-utils";
  inputs.templ.url = "github:a-h/templ";

  outputs = { self, nixpkgs, flake-utils, gomod2nix, ewxfront, templ }:
    (flake-utils.lib.eachDefaultSystem
      (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
          templOverlay = templ.packages.${system}.templ;

          # The current default sdk for macOS fails to compile go projects, so we use a newer one for now.
          # This has no effect on other platforms.
          callPackage = pkgs.darwin.apple_sdk_11_0.callPackage or pkgs.callPackage;
        in
        {
          packages.default = callPackage ./nix {
            inherit (gomod2nix.legacyPackages.${system}) buildGoApplication;
          };
          devShells.default = callPackage ./nix/shell.nix {
            inherit (gomod2nix.legacyPackages.${system}) mkGoEnv gomod2nix ;
          };

          devShells.nvim = ewxfront.devShells.${system}.default ;
          devShells.templ = with pkgs; mkShell {
         	   buildInputs = [ templOverlay 
               ewxfront.devShells.${system}.default 
              ];
	  };
	  
        })
    );
}
