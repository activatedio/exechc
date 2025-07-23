with import <nixpkgs> {};

stdenv.mkDerivation {

  name = "ops";

  buildInputs = with pkgs; [
    go
    gnumake
  ];

}

