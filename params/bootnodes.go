// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
    "enode://218aa67fa58287049e6856380e910bdf7d5f3aec6e1384ae53d473745a49aecef7e73eff371fe487aa6dab8185f938a1d294fe94c8027537da8d59b8875359b2@45.90.57.29:30303",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the test network.
var TestnetBootnodes = []string{
}

var V5Bootnodes = []string{
}

const dnsPrefix = "enrtree://APPP5TMYEINO4O5JHEHQWDJR3WVNB4QKO7YJERUKFA4EKP7YJ76A4@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case TestnetGenesisHash:
		net = "testnet"
	default:
		return ""
	}
	return dnsPrefix + protocol + ".disco." + net + ".bpxchain.cc"
}
