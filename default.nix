{ pkgs ? import <nixpkgs> { } }:
pkgs.buildGo117Module {
  pname = "ircd-in-go";
  version = "0";
  src = builtins.path { path = ./.; name = "ircd-in-go"; };
  vendorSha256 = null;
}
