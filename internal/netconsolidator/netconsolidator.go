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

func Parsecidr(subnets []string) []*net.IPNet {
	var nets []*net.IPNet
	var isErrors = false

	for _, s := range subnets {
		_, subnet, err := net.ParseCIDR(s)
		if err != nil {
			fmt.Printf("Can't parse subnet: %v\n", err)
			isErrors = true
			continue
		}
		nets = append(nets, subnet)
	}
	if isErrors {
		return nil
	}
	return nets
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
	nets := Parsecidr(subnets)
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
