{
  description = "Ebitengine (Ebiten) Go Development Environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, utils }:
    utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShells.default = pkgs.mkShell {
          # Tools needed at build time
          nativeBuildInputs = with pkgs; [
            go
            pkg-config
            gcc
          ];

          # Libraries Ebiten needs to link against (CGO)
          buildInputs = with pkgs; [
            # X11 dependencies
            xorg.libX11
            xorg.libXcursor
            xorg.libXi
            xorg.libXinerama
            xorg.libXrandr
            xorg.libXxf86vm
            # Graphics and Audio
            libGL
            alsa-lib
          ];

          # Ebitengine specific environment setup
          shellHook = ''
            export CGO_ENABLED=1
            
            # This allows the compiled binary to find the OpenGL driver at runtime
            export LD_LIBRARY_PATH="${pkgs.lib.makeLibraryPath [ pkgs.libGL pkgs.xorg.libX11 ]}:$LD_LIBRARY_PATH"
            
            echo "🔥 Ebitengine dev environment loaded!"
            echo "Go version: $(go version)"
          '';
        };
      });
}
