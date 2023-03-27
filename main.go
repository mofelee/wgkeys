package main

import (
    "fmt"
    "golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func main() {
    // Generate a new private key
    privateKey, err := wgtypes.GeneratePrivateKey()
    if err != nil {
        panic(err)
    }

    // Derive the corresponding public key
    publicKey := privateKey.PublicKey()

    // Print the private and public keys in hexadecimal format
    fmt.Printf("private: %s\n", privateKey.String())
    fmt.Printf("public: %s\n", publicKey.String())
}
