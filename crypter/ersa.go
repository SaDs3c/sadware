package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"hash"
	"io"
	"os"
)

type RSA struct {
	pub        *rsa.PublicKey
	err        error
	hash       hash.Hash
	random     io.Reader
	cipherText []byte
	plainText  []byte
}

func (key *RSA) importedKey() {
	f, err := os.Open("pubPemFile.pem")
	if err != nil {
		fmt.Println(err)
	}

	pemFileData, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	pubPemBlock, _ := pem.Decode(pemFileData)
	pubKeyImported, err := x509.ParsePKCS1PublicKey(pubPemBlock.Bytes)
	if err != nil {
		fmt.Println(err)
	}

	key.pub = pubKeyImported

}

func (key *RSA) rsaEncrypt() []byte {
	key.cipherText, key.err = rsa.EncryptOAEP(key.hash, key.random, key.pub, key.plainText, nil)

	if key.err != nil {
		fmt.Println(key.err.Error())
	}

	return key.cipherText

}
