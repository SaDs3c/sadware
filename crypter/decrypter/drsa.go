package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"hash"
	"io"
	"os"
)

type RSA struct {
	priv       *rsa.PrivateKey
	err        error
	hash       hash.Hash
	random     io.Reader
	cipherText []byte
	plainText  []byte
}

func NewRSA() *RSA {
	rsaS := &RSA{hash: sha256.New(), random: rand.Reader}
	return rsaS
}

func (key *RSA) importedKey() {
	f, err := os.Open("privPemFile.pem")
	if err != nil {
		fmt.Println(err)
	}

	pemFileData, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	privPemBlock, _ := pem.Decode(pemFileData)
	privKeyImported, err := x509.ParsePKCS1PrivateKey(privPemBlock.Bytes)
	if err != nil {
		fmt.Println(err)
	}

	key.priv = privKeyImported

}

func (key *RSA) rsaDecrypt() []byte {
	key.importedKey()
	key.messageFile()
	key.plainText, key.err = rsa.DecryptOAEP(key.hash, key.random, key.priv, key.cipherText, nil)

	if key.err != nil {
		fmt.Println(key.err.Error())
	}

	return key.plainText

}

func (key *RSA) messageFile() {
	f, err := os.Open("../sad.json")
	if err != nil {
		fmt.Println(err)
	}

	secret, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	key.cipherText = secret

}
