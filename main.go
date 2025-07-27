package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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
	return fmt.Sprintf("%s - %s %s\n", e.city, e.country, e.name)
}

func newNode(peer Node) *exitNode {
	var country, city string

	if peer.Location.Country == "" {
		country = "N/A"
	} else {
		country = peer.Location.Country
	}

	if peer.ExitNode {
		city = "Disable: "
	}

	if peer.Location.City == "" {
		city += "N/A"
	} else {
		city += peer.Location.City
	}

	node := exitNode{
		name:     peer.DnsName,
		country:  country,
		city:     city,
		selected: peer.ExitNode,
	}

	return &node
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
	tailscaleStatusCmd.Stderr = os.Stderr
	output, err := tailscaleStatusCmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	status := Status{}
	json.Unmarshal(output, &status)

	allExitNodes := exitNodes{}

	for _, peer := range status.Peer {
		if peer.ExitNodeOption {
			node := newNode(peer)
			if peer.ExitNode {
				allExitNodes = append(exitNodes{*node}, allExitNodes...)
			} else {
				allExitNodes = append(allExitNodes, *node)
			}
		}
	}

	nodeListString := allExitNodes.String()

	wofiCmd := exec.Command("wofi", "--show", "dmenu", "-W", "25%")
	wofiCmd.Stdin = strings.NewReader(nodeListString)
	wofiOuput, err := wofiCmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	var name string
	if strings.HasPrefix(string(wofiOuput), "Disable") {
		name = ""
	} else {
		splitWofiOuput := strings.Split(string(wofiOuput), " ")
		name = splitWofiOuput[len(splitWofiOuput)-1]
		name = strings.TrimSuffix(name, "\n")
	}

	tailscaleSetExitNodeCmd := exec.Command("tailscale", "set", "--exit-node", name)
	tailscaleSetExitNodeCmd.Stderr = os.Stderr
	err = tailscaleSetExitNodeCmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
