package genburn

import (
	"fmt"
	"strings"
	"testing"
	"encoding/hex"

	"github.com/btcsuite/btcd/chaincfg"
)

func TestFindwhcAddress(t *testing.T) {
	for v := 0; v < 10000000; v++ {
		x := fmt.Sprintf("%02x", v)
		len1 := len(x)
		for i := len1; i <= 6; i++ {
			x = fmt.Sprintf("0%s", x)
		}

		result := fmt.Sprintf("000000000000000000000000000000000%s", x)

		hash160, err := hex.DecodeString(result)
		if err != nil {
			t.Error(err) // encoding/hex: odd length hex string
		}
		address, err1 := NewCashAddressPubKeyHash(hash160, &chaincfg.MainNetParams)
		if err1 != nil {
			t.Error(err1)
			return
		}

		if strings.HasSuffix(address.String(), "whc") {
			addr1 := fmt.Sprintf("%v", address)
			if addr1[len(addr1)-4] >= '0' && addr1[len(addr1)-4] <= '9' && strings.HasSuffix(addr1, "whc") {
				addr2 := fmt.Sprintf("%v", addr1)
				if strings.HasSuffix(addr2, "8whc") {
					fmt.Println("===========find burn address last 4 8whc=========")
					addr3 := fmt.Sprintf("%v", addr2)
					fmt.Printf("the public hash value is:%v\nthe burn addr value is:%v\n", result, addr2)
					if addr3[len(addr3)-5] >= '0' && addr3[len(addr3)-5] <= '9' && strings.HasSuffix(addr3, "whc") {
						fmt.Println("===========find burn address last 5 [0-9]8whc=========")
						fmt.Printf("the public hash value is:%v\nthe burn addr value is:%v\n", result, addr3)
					}
				}
			}
		}
	}
}
