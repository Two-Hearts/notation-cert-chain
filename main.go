package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/notaryproject/notation-core-go/signature"
	"github.com/notaryproject/notation-core-go/signature/cose"
)

const SignatureMediaType = cose.MediaTypeEnvelope

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Number of input signature envelope has to be 1")
		os.Exit(1)
	}
	sigEnvPath := os.Args[1]
	sigEnvBytes, err := os.ReadFile(sigEnvPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read in signature file: %v\n", err)
		os.Exit(1)
	}
	sigEnv, err := signature.ParseEnvelope(SignatureMediaType, sigEnvBytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse signature envelope: %v\n", err)
		os.Exit(1)
	}
	envContent, err := sigEnv.Content()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get signature envelope content: %v\n", err)
		os.Exit(1)
	}
	var thumbprints []string
	for _, cert := range envContent.SignerInfo.CertificateChain {
		checkSum := sha256.Sum256(cert.Raw)
		thumbprints = append(thumbprints, hex.EncodeToString(checkSum[:]))
	}
	val, err := json.Marshal(thumbprints)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal thumbprints: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(val))
}
