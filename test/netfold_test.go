package test

import (
	"github.com/stretchr/testify/require"
	"netfold/internal/netconsolidator"
	"testing"
)

func Test_complete(t *testing.T) {
	sourceSubnets := []string{
		"1011:4d2::/32",
		"1013:4836::/48",
		"1013:4836:1::/48",
		"1013:4836:2::/48",
		"1013:4836:3::/48",
		"1013:4836:10::/46",
		"101a:a836:0:100::/56",
		"101a:a836:0:4100::/56",
		"101a:a836:0:8100::/56",
		"1013:4836:0:f100::/56",
		"1013:4836:10:0:9011::/80",
		"1013:4836:11:0:9011::/80",
		"1013:4836:12:0:9011::/80",
		"101a:a836:1:0:9000:5a::/96",
		"1013:4836:10:0:9010:5a::/96",
		"1013:4836:11:0:9010:5a::/96",
		"1013:4836:12:0:9010:5a::/96",
		"1013:4836:10:0:9010:5a::/96",
		"1013:4836:11:0:9010:5a::/96",
		"1013:4836:12:0:9010:5a::/96",
		"1013:4836:10:0:9010:b1::/96",
		"1013:4836:11:0:9010:b1::/96",
		"1013:4836:12:0:9010:b1::/96",
		"1013:4836:10:0:9010:6f::/112",
		"1013:4836:11:0:9010:6f::/112",
		"1013:4836:12:0:9010:6f::/112",
		"101a:a836:1:4000:9000:5a::/96",
		"101a:a836:1:8000:9000:5a::/96",
		"1013:4836:0:1001:aaaa:aaaa:bbbb:cccc/128",
	}

	targetSubnets := []string{
		"1011:4d2::/32",
		"1013:4836::/48",
		"1013:4836:1::/48",
		"1013:4836:2::/48",
		"1013:4836:3::/48",
		"1013:4836:10::/46",
		"101a:a836:0:100::/56",
		"101a:a836:0:4100::/56",
		"101a:a836:0:8100::/56",
		"101a:a836:1:0:9000:5a::/96",
		"101a:a836:1:4000:9000:5a::/96",
		"101a:a836:1:8000:9000:5a::/96",
	}

	consolidated := netconsolidator.Consolidate(sourceSubnets)

	require.EqualValues(t, targetSubnets, consolidated)
}

func Test_NetsParsing(t *testing.T) {
	sourceSubnets := []string{
		"1011:4d2::/256",
		"1011:4d2::",
		"1011:4d2:0",
		"",
		"VRED&^ &F!*O&HPU",
		"123.123.123.123",
		"1123123123",
	}
	nets, netsWithError := netconsolidator.Parsecidr(sourceSubnets)

	require.NotEmpty(t, netsWithError)
	require.EqualValues(t, sourceSubnets, netsWithError)
	require.Empty(t, nets)
}
