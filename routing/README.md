routing
=======

[![Build Status](http://img.shields.io/travis/brronsuite/lnd.svg)](https://travis-ci.org/brronsuite/lnd) 
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/brronsuite/lnd/blob/master/LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/brronsuite/lnd/routing)

The routing package implements authentication+validation of channel
announcements, pruning of the channel graph, path finding within the network,
sending outgoing payments into the network and synchronizing new peers to our
channel graph state.

## Installation and Updating

```shell
⛰  go get -u github.com/brronsuite/lnd/routing
```
