package models

import (
    "log"
    "fmt"
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "github.com/btcsuite/btcutil/base58"
    "golang.org/x/crypto/ripemd160"
)

type Wallet struct {
    UserID        uint      `gorm:"primaryKey;autoIncrement:false;" json:"user_id"`
    BaseTimeField
    PrivateKey    string `gorm:"size:256;not null;unique" json:"private_key"`
    PublicKey     string  `gorm:"size:256;not null;unique" json:"public_key"`
    BlockchainAddress  string    `gorm:"size:256;not null;unique" json:"blockchain_address"`
    Amount float32 `gorm:"default:0.0" json"amount"`
}

func NewWallet() *Wallet {
    w := &Wallet{}

    privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        log.Println("Failed to generate a private key", err)
    }
    publicKey := &privateKey.PublicKey
    w.PrivateKey = makePriKeyStr(privateKey)
    w.PublicKey = makePubKeyStr(publicKey)
    w.BlockchainAddress = generateBlockChainAddress(publicKey)

    return w
}

func makePriKeyStr(priKey *ecdsa.PrivateKey) string {
    return fmt.Sprintf("%x", priKey.D.Bytes())
}

func makePubKeyStr(pubKey *ecdsa.PublicKey) string {
    return fmt.Sprintf("%064x%064x", pubKey.X.Bytes(), pubKey.Y.Bytes())
}

func generateBlockChainAddress(publicKey *ecdsa.PublicKey) string {
    // ref: How to create Bitcoin Address
    // https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses

    // Perform SHA-256 hashing on the public key
    sha256Hash := sha256.New()
    sha256Hash.Write(publicKey.X.Bytes())
    sha256Hash.Write(publicKey.Y.Bytes())
    sha256Digest := sha256Hash.Sum(nil)

    // Perform RIPEMD-160 hashing on the result of SHA-256
    ripemd160Hash := ripemd160.New()
    ripemd160Hash.Write(sha256Digest)
    ripemd160Digest := ripemd160Hash.Sum(nil)

    // Add version byte in front of RIPEMD-160 hash (0x00 for Main Network)
    vb := make([]byte, 21)
    vb[0] = 0x00
    copy(vb[1:], ripemd160Digest[:])

    // Perform SHA-256 hash on the extended RIPEMD-160 result
    vbSha256Hash := sha256.New()
    vbSha256Hash.Write(vb)
    vbSha356HashDigest := vbSha256Hash.Sum(nil)

    // Perform SHA-256 hash on the result of the previous SHA-256 hash
    lastSha256Hash := sha256.New()
    lastSha256Hash.Write(vbSha356HashDigest)
    digest := lastSha256Hash.Sum(nil)

    // Take the first 4 bytes of the second SHA-256 hash. This is the address checksum
    checksum := digest[:4]

    // Add the 4 checksum bytes at the end of extended RIPEMD-160.
    binaryAddress := make([]byte, 25)
    copy(binaryAddress[:21], vb[:])
    copy(binaryAddress[21:], checksum[:])

    // Convert the result from a byte string into a base58 string using Base58Check encoding.
    return base58.Encode(binaryAddress)
}
