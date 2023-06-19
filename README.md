# notation-cert-chain
## Input
a COSE signature envelope following Notation's [specs](https://github.com/notaryproject/notaryproject/blob/main/specs/signature-envelope-cose.md).
## Output
write to standard output SHA256 checksums of certificate chain in the signature envelope. The certificate containing the public key corresponding to the key used to digitally sign the COSE is the first certificate, followed by intermediate and root certificates in the correct order. <br> 
An example of the output (the whole array is in string format):
`["303d73d0e9b4efe97e04bd58ad55423669f2a77c36e44408cc516388afbb18eb","9b1894f223d934cbd6575af3c6e1f6096b9221a7da132185f5a5cdc92235b5dc","23ffe2b8bdb9a1711515d4cffda04bc7f793d513c76c243f1020507d8669b7db"]`
## Usage 
`notation-cert-chain <file_path_to_COSE_signature_envelope>`