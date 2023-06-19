# notation-cert-chain
This tool takes a COSE signature envelope following Notation's [spec]<https://github.com/notaryproject/notaryproject/blob/main/specs/signature-envelope-cose.md> and prints out SHA256 checksum of certificate chain in the signature envelope. The certificate containing the public key corresponding to the key used to digitally sign the COSE is the first certificate, followed by the intermediate and root certificates in the correct order. An example of the output:
`["303d73d0e9b4efe97e04bd58ad55423669f2a77c36e44408cc516388afbb18eb","9b1894f223d934cbd6575af3c6e1f6096b9221a7da132185f5a5cdc92235b5dc","23ffe2b8bdb9a1711515d4cffda04bc7f793d513c76c243f1020507d8669b7db"]` <br>
Usage: <br> 
`notation-cert-chain <file_path_of_COSE_signature_envelope>` <br>
