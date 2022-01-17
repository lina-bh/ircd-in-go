{ pkgs ? import <nixpkgs> { } }:
let
  default = pkgs.callPackage ./default.nix { };
in
default.overrideAttrs (oldAttrs: {
  nativeBuildInputs = with pkgs; [
    irssi
    go-tools
    gopls
    gopkgs
    go-outline
    delve
  ] ++ oldAttrs.nativeBuildInputs;
  shellHook = "echo shell";
})
