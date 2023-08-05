package controller

import (
	nodev1 "github.com/log/api/v1"
	v1 "k8s.io/api/core/v1"
)

// GetSaleIPs 找出 IP 在 nodesInCR 而不在  nodesInCluster，表示需要扩容

func getNodesInternalIP(nodes *v1.NodeList) []string {
	var ips []string
	for _, node := range nodes.Items {
		for _, addr := range node.Status.Addresses {
			if addr.Type == "InternalIP" {
				ips = append(ips, addr.Address)
			}
		}
	}
	return ips
}

func GetSaleIPs(nodesInCR [] nodev1.Node, nodesInCluster *v1.NodeList)  []string {
	var delta []string
	in := false
	ipsInCLuster := getNodesInternalIP(nodesInCluster)
	for _, nodeInCR := range nodesInCR {
		in = false
		for _, ipInCluster := range ipsInCLuster {
			if nodeInCR.Ip == ipInCluster {
				in = true
				break
			}
		}
		if ! in {
			delta = append(delta, nodeInCR.Ip)
		}
	}
	return delta
}


// GetShrinkIPs 找出 IP 在 nodesInCluster 而不在 nodesInCR，表示需要缩容

func GetShrinkIPs(nodesInCR [] nodev1.Node, nodesInCluster *v1.NodeList)  []string {
	var delta []string
	in := false
	ipsInCLuster := getNodesInternalIP(nodesInCluster)
	for _, ipInCLuster := range ipsInCLuster {
		in = false
		for _, nodeInCR := range nodesInCR {
			if nodeInCR.Ip == ipInCLuster {
				in = true
				break
			}
		}
		if ! in {
			delta = append(delta, ipInCLuster)
		}
	}
	return delta
}