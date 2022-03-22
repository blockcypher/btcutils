Usage
-----

Utility to sign some provided data (usually a signature transaction hash) with the provided private key using the Bitcoin ECDSA curve. Both the data and the private key are expected to be hex-encoded.

Usage: signer [datahex] [privatehex]

Example:

```shell
$ ./signer 646b5cc387cef8ced58d861c2ddae75568b4936ccb2971371a0f9d2321460381 b3bd48cbfc88ddcaa9b600c90a62a07fe6c27503ae5b4872c041e5aecf2e723d
304402201fce11a7b612f9bc7d446054b1e661836f65cfecee0581700320a4c37caa9c2e0220540ea2d2edbe182e28a24afd0ffe29175eadea7c5ba69c5ed3ab0745b49cb9b0
```

As a friendly reminder, if you run this example locally, your resulting signature may differ---[ECDSA](https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm) does not furnish deterministic signatures, but they are no less valid. 

Building
--------

Install Go (http://golang.org/doc/install) and run:

```shell
git clone https://github.com/blockcypher/btcutils.git

cd btcutils/signer
go build
```

This will generate the signer binary. 
