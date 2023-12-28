package netconsolidator

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func isSubset(subnet1 *net.IPNet, subnet2 *net.IPNet) bool {
	return subnet2.Contains(subnet1.IP) && subnet2.Mask.String() <= subnet1.Mask.String()
}

func Parsecidr(subnets []string) ([]*net.IPNet, []string) {
	var nets []*net.IPNet
	var netsWithErrors []string

	for _, s := range subnets {
		_, subnet, err := net.ParseCIDR(s)
		if err != nil {
			netsWithErrors = append(netsWithErrors, s)
			continue
		}
		nets = append(nets, subnet)
	}

	return nets, netsWithErrors
}

func ConsolidateSubnets(nets []*net.IPNet) []string {
	var result []string

	for i, net1 := range nets {
		subset := false
		for j, net2 := range nets {
			if i != j && isSubset(net1, net2) {
				subset = true
				break
			}
		}
		if !subset {
			result = append(result, net1.String())
		}
	}

	return result
}

func Consolidate(subnets []string) []string {
	nets, netsWithErrors := Parsecidr(subnets)
	if len(netsWithErrors) > 0 {
		for s := range netsWithErrors {
			fmt.Fprintf(os.Stderr, "invalid CIDR address: %s\n", s)
		}
	}
	consolidated := ConsolidateSubnets(nets)
	return consolidated
}

func Main() {
	scanner := bufio.NewScanner(os.Stdin)
	var subnets []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			subnets = append(subnets, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	consolidatedNets := Consolidate(subnets)
	for _, subnet := range consolidatedNets {
		fmt.Println(subnet)
	}
}
