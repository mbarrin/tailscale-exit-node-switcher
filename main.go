package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type exitNode struct {
	name     string
	country  string
	city     string
	selected bool
}

func (e *exitNode) String() string {
	return fmt.Sprintf("%s %s %s %t\n", e.name, e.country, e.city, e.selected)
}

type exitNodes []exitNode

func (e *exitNodes) String() string {
	nodes := ""
	for _, exitNode := range *e {
		nodes += exitNode.String()
	}
	return nodes
}

func main() {
	tailscaleStatusCmd := exec.Command("tailscale", "status", "--json")
	output, err := tailscaleStatusCmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	status := Status{}
	json.Unmarshal(output, &status)

	allExitNodes := exitNodes{}

	for _, peer := range status.Peer {
		if peer.ExitNodeOption {
			node := exitNode{
				name:     peer.DnsName,
				country:  peer.Location.Country,
				city:     peer.Location.City,
				selected: peer.ExitNode,
			}
			if peer.ExitNode {
				allExitNodes = append(exitNodes{node}, allExitNodes...)
			} else {
				allExitNodes = append(allExitNodes, node)
			}
		}
	}

	nodeListString := allExitNodes.String()
	if allExitNodes[0].selected {
		nodeListString = " disable\n" + nodeListString
	}

	wofiCmd := exec.Command("wofi", "--show", "dmenu", "-W", "50%")
	wofiCmd.Stdin = strings.NewReader(nodeListString)
	wofiOuput, err := wofiCmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	name := strings.Split(string(wofiOuput), " ")[0]

	tailscaleSetExitNodeCmd := exec.Command("tailscale", "set", "--exit-node", name)
	err = tailscaleSetExitNodeCmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
