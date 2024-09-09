{
  description = "A basic gomod2nix flake";
  inputs.ewx.url = "github:kyp0717/ewxback";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";
  inputs.gomod2nix.url = "github:nix-community/gomod2nix";
  inputs.gomod2nix.inputs.nixpkgs.follows = "nixpkgs";
  inputs.gomod2nix.inputs.flake-utils.follows = "flake-utils";

  outputs = { self, nixpkgs, flake-utils, gomod2nix, ewx }:
    (flake-utils.lib.eachDefaultSystem
      (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};

          # The current default sdk for macOS fails to compile go projects, so we use a newer one for now.
          # This has no effect on other platforms.
          callPackage = pkgs.darwin.apple_sdk_11_0.callPackage or pkgs.callPackage;
	  goEnv = callPackage ./shell.nix {
            inherit (gomod2nix.legacyPackages.${system}) mkGoEnv gomod2nix;
	  };
        in
        {
          packages.default = callPackage ./. {
            inherit (gomod2nix.legacyPackages.${system}) buildGoApplication;
          };
          #devShells.default = callPackage ./shell.nix {
            #inherit (gomod2nix.legacyPackages.${system}) mkGoEnv gomod2nix;
          #};

	  devShells.default = pkgs.mkShell {
	    buildInputs  = [goEnv];
	    inputsFrom = [ewx.devShells.${system}.default];
	  };
	  
        })
    );
}
