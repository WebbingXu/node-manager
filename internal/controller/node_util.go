package controller

import (
	nodev1 "github.com/log/api/v1"
	v1 "k8s.io/api/core/v1"
)

func NodeDiff(nodesInCR [] nodev1.Node, nodesInCluster *v1.NodeList)  []string {
	var delta []string
	in := false
	for _, nodesInCR := range nodesInCR {
		in = false
		for _, nodeInCLuster := range nodesInCluster.Items {
			for _, address := range nodeInCLuster.Status.Addresses {
				if nodesInCR.Ip == address.Address {
					in = true
					break
				}
			}
		}
		if ! in {
			delta = append(delta, nodesInCR.Ip)
		}
	}
	return delta
}