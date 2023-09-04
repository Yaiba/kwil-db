package crypto

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cometbft/cometbft/crypto/secp256k1"
	"github.com/stretchr/testify/assert"
)

func TestComebftSecp256k1Signer_Sign(t *testing.T) {
	msg := []byte("foo")
	pvKeyHex := "f1aa5a7966c3863ccde3047f6a1e266cdc0c76b399e256b8fede92b1c69e4f4e"
	expectSignatureHex := "19a4aced02d5b9142b4f622b06442b1904445e16bd25409e6b0ff357ccc021d001d0e7824654b695b4b6e0991cb7507f487b82be4b2ed713d1e3e2cbc3d2518d"
	expectSignatureBytes, _ := hex.DecodeString(expectSignatureHex)
	expectSig := &Signature{
		Signature: expectSignatureBytes,
		Type:      SignatureTypeSecp256k1Cometbft,
	}
	require.Equal(t, SignatureSecp256k1CometbftLength, len(expectSignatureBytes), "invalid signature length")

	// comebft secp256k1 private key, and signature
	pvKeyBytes, _ := hex.DecodeString(pvKeyHex)
	cometBftSecp256k1Key := secp256k1.PrivKey(pvKeyBytes)
	cometBfgSecp256k1Sig, err := cometBftSecp256k1Key.Sign(msg) // sha256 is done in `Sign`
	assert.NoError(t, err, "error signing message")
	assert.Equal(t, expectSignatureHex, hex.EncodeToString(cometBfgSecp256k1Sig), "signature mismatch")

	// use the kwil secp256k1 private key and cometbft signer to sign the message
	kwilCometBftKey, err := Secp256k1PrivateKeyFromHex(pvKeyHex)
	require.NoError(t, err, "error parse private pvKeyHex")

	cometBfgSigner := NewCometbftSecp256k1Signer(kwilCometBftKey)
	sig, err := cometBfgSigner.Sign(msg)
	assert.NoError(t, err, "error signing message")
	require.EqualValues(t, expectSig, sig, "signature mismatch")

	err = sig.Verify(kwilCometBftKey.PubKey(), msg)
	assert.NoError(t, err, "error verifying signature")
}

func TestEthPersonalSecp256k1Signer_Sign(t *testing.T) {
	msg := []byte("foo")
	pvKeyHex := "f1aa5a7966c3863ccde3047f6a1e266cdc0c76b399e256b8fede92b1c69e4f4e"
	expectSignatureHex := "cb3fed7f6ff36e59054c04a831b215e514052753ee353e6fe31d4b4ef736acd6155127db555d3006ba14fcb4c79bbad56c8e63b81a9896319bb053a9e253475800"
	expectSignatureBytes, _ := hex.DecodeString(expectSignatureHex)
	expectSig := &Signature{
		Signature: expectSignatureBytes,
		Type:      SignatureTypeSecp256k1Personal,
	}
	require.Equal(t, SignatureSecp256k1PersonalLength, len(expectSignatureBytes), "invalid signature length")

	pk, err := Secp256k1PrivateKeyFromHex(pvKeyHex)
	require.NoError(t, err, "error parse private pvKeyHex")

	ethSigner := NewEthPersonalSecp256k1Signer(pk)

	sig, err := ethSigner.Sign(msg)
	require.NoError(t, err, "error signing msg")
	assert.EqualValues(t, expectSig, sig, "signature mismatch")

	err = sig.Verify(pk.PubKey(), msg)
	assert.NoError(t, err, "error verifying signature")
}

func TestStdEd25519Signer_Sign(t *testing.T) {
	msg := []byte("foo")
	pvKeyHex := "7c67e60fce0c403ff40193a3128e5f3d8c2139aed36d76d7b5f1e70ec19c43f00aa611bf555596912bc6f9a9f169f8785918e7bab9924001895798ff13f05842"
	expectSignature := "59b2db2d1e4ce6f8771453cfc78d1f943723528f00fa14adf574600f15c601d591fa2ba29c94d9ed694db324f9e8671bdfbcba4b8e10f6a8733682fa3d115f0c"
	expectSignatureBytes, _ := hex.DecodeString(expectSignature)
	expectSig := &Signature{
		Signature: expectSignatureBytes,
		Type:      SignatureTypeEd25519,
	}
	require.Equal(t, SignatureEd25519Length, len(expectSignatureBytes), "invalid signature length")

	pk, err := Ed25519PrivateKeyFromHex(pvKeyHex)
	require.NoError(t, err, "error parse private key")

	edSigner := NewStdEd25519Signer(pk)

	sig, err := edSigner.Sign(msg)
	require.NoError(t, err, "error sign")
	assert.EqualValues(t, expectSig, sig, "signature mismatch")

	err = sig.Verify(pk.PubKey(), msg)
	assert.NoError(t, err, "error verifying signature")
}

func TestNearEd25519Signer_Sign(t *testing.T) {
	msg := []byte("foo")
	pvKeyHex := "7c67e60fce0c403ff40193a3128e5f3d8c2139aed36d76d7b5f1e70ec19c43f00aa611bf555596912bc6f9a9f169f8785918e7bab9924001895798ff13f05842"
	expectSignature := "089bcf52220dad77abc2cfcb1639bcb2944fdf64e0b173f40cd0d144bdbf7808f4eff3716eb3e98ed40be3ab126e1449d5f57efbe5626673059edc90e9cd9801"
	expectSignatureBytes, _ := hex.DecodeString(expectSignature)
	expectSig := &Signature{
		Signature: expectSignatureBytes,
		Type:      SignatureTypeEd25519Near,
	}
	require.Equal(t, SignatureEd25519Length, len(expectSignatureBytes), "invalid signature length")

	pk, err := Ed25519PrivateKeyFromHex(pvKeyHex)
	require.NoError(t, err, "error parse private key")

	nearSigner := NewNearSigner(pk)

	sig, err := nearSigner.Sign(msg)
	require.NoError(t, err, "error sign")
	assert.EqualValues(t, expectSig, sig, "signature mismatch")
	fmt.Println(hex.EncodeToString(sig.Signature))

	err = sig.Verify(pk.PubKey(), msg)
	assert.NoError(t, err, "error verifying signature")
}

func TestTrivialSigner_Sign(t *testing.T) {
	signer := NewTrivialSigner(nil)
	_, err := signer.Sign([]byte("foo"))
	require.Error(t, err, "suppose to error")
}