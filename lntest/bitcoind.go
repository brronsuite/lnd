//go:build bitcoind && !notxindex
// +build bitcoind,!notxindex

package lntest

import (
	"github.com/brronsuite/brond/chaincfg"
)

// NewBackend starts a bitcoind node with the txindex enabled and returns a
// BitcoindBackendConfig for that node.
func NewBackend(miner string, netParams *chaincfg.Params) (
	*BitcoindBackendConfig, func() error, error) {

	extraArgs := []string{
		"-debug",
		"-regtest",
		"-txindex",
		"-disablewallet",
	}

	return newBackend(miner, netParams, extraArgs)
}
