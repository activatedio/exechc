# exechc

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/activatedio/exechc/ci.yaml?branch=main&style=flat-square)](https://github.com/activatedio/exechc/actions?query=workflow%3ACI)
[![Go Report Card](https://goreportcard.com/badge/github.com/activatedio/exechc?style=flat-square)](https://goreportcard.com/report/github.com/activatedio/exechc)
![Go Version](https://img.shields.io/github/go-mod/go-version/activatedio/exechc?style=flat-square)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/activatedio/exechc)](https://pkg.go.dev/mod/github.com/activatedio/exechc)


Command EXECution Health Check

A simple HTTP server which returns a 200 or 500 based on the output of a
specified command. Useful for applications such as health checks for supporting
UDP-based load balancer targets.

This was inspried by usage for Wireguard as inspiried by this post: 
https://www.procustodibus.com/blog/2021/10/ha-wireguard-site-to-site/






