/*
	Generating ed25519 Private/Public Key Pair:
	Public Key: Ln081IZnjKInTLIOumuj3yRq30VvOU5mm9IwBgeGiXs=
	Private Key: OmwQUAAXyKArmg9Cmt5W6u79K0SvmpAk+JXUoqr55yQufTzUhmeMoidMsg66a6PfJGrfRW85Tmab0jAGB4aJew==

*/

package transactions

import (
	"testing"
)

func TestCalculatePublicKey(t *testing.T) {
	publicKeyBase64 := CalculatePublicKey("OmwQUAAXyKArmg9Cmt5W6u79K0SvmpAk+JXUoqr55yQufTzUhmeMoidMsg66a6PfJGrfRW85Tmab0jAGB4aJew==")
	if publicKeyBase64 != "Ln081IZnjKInTLIOumuj3yRq30VvOU5mm9IwBgeGiXs=" {
		t.Errorf("CalculatePublicKey was wrong, got: %s, want: %s.", publicKeyBase64, "Ln081IZnjKInTLIOumuj3yRq30VvOU5mm9IwBgeGiXs=")
	}
}
