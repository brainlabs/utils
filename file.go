// Package utils
package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

// ReadFileAsBytes function read file and return bytes
func ReadFileAsBytes(filePath string) ([]byte, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return content, nil
}

// WriteFile function write file
func WriteFile(content []byte, filename string) (err error) {
	filepath := fmt.Sprintf("%s", filename)

	err = ioutil.WriteFile(filepath, content, 0644)
	if err != nil {
		return err
	}

	return nil
}


// ReadPublicKeyByte function rsa read public key from file
func ReadPublicKeyByte(publicKeyData []byte) (publicKey *rsa.PublicKey, err error) {
	block, _ := pem.Decode(publicKeyData)
	if block == nil {
		err = fmt.Errorf("failed to parse PEM block containing the key")
		return
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		publicKey = pub
	default:
		err = fmt.Errorf("failed to parse the file to Public Key")
	}


	return
}

// ReadPublicKeyFile function read file rsa public key
func ReadPublicKeyFile(filePath string) (pubKey *rsa.PublicKey, err error) {
	pemByte, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	pubKey, err = ReadPublicKeyByte(pemByte)

	return
}

// ReadPrivateKey function RSA Read private key from byte
func ReadPrivateKey(privateKeyData []byte) (privateKey *rsa.PrivateKey, err error) {
	block, _ := pem.Decode(privateKeyData)
	if block == nil {
		err = fmt.Errorf("failed to parse PEM block containing the key")
		return
	}

	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}

	return
}

// ReadPrivateKeyFile function RSA Read private key from file
func ReadPrivateKeyFile(filepath string) (privateKey *rsa.PrivateKey, err error) {
	privPEM, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}
	privateKey, err = ReadPrivateKey(privPEM)
	return
}



