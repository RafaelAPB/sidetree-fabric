/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package protocol

import (
	"sort"

	"github.com/trustbloc/sidetree-core-go/pkg/api/protocol"
)

// Client is a struct which holds a list of protocols.
type Client struct {
	protocols []protocol.Protocol
}

//New initializes the protocol parameters from file
func New(protocolVersions map[string]protocol.Protocol) *Client {
	// Creating the list of the protocol versions
	protocols := make([]protocol.Protocol, 0, len(protocolVersions))
	for _, v := range protocolVersions {
		protocols = append(protocols, v)
	}

	// Sorting the protocolParameter list based on blockChain start time
	sort.SliceStable(protocols, func(i, j int) bool {
		return protocols[j].StartingBlockChainTime > protocols[i].StartingBlockChainTime
	})

	return &Client{protocols: protocols}
}

//Current returns the latest version of protocol
func (c *Client) Current() protocol.Protocol {
	return c.protocols[len(c.protocols)-1]
}
