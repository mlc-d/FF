package keys

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"log"
	"os"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

// JwkSet holds the Public and Private rsa keys.
type JwkSet struct {
	Public  jwk.Key
	Private jwk.Key
}

var (
	keys = new(JwkSet)

	bitSize = flag.Int("bits", 3072, "size in bits of the rsa key")
)

// GetKeys singleton function to initialize the [keys] variable, or to return its
// current value.
func GetKeys() *JwkSet {
	if keys.Private != nil || keys.Public != nil {
		return keys
	}
	// create a newJWTRepo RSA Key
	rawRSAPrivateKey, err := rsa.GenerateKey(rand.Reader, *bitSize)
	if err != nil {
		log.Fatalln(err)
	}
	rawRSAPublicKey := rawRSAPrivateKey.PublicKey

	{ // set fields for Private key
		keys.Private, err = jwk.FromRaw(rawRSAPrivateKey)
		if err != nil {
			log.Fatalln(err)
		}
		err = keys.Private.Set(jwk.AlgorithmKey, jwa.RS256)
		if err != nil {
			log.Fatalln(err)
		}
		err = keys.Private.Set(jwk.KeyIDKey, "Private-key")
	}

	{ // set fields for Public key
		keys.Public, err = jwk.FromRaw(rawRSAPublicKey)
		if err != nil {
			log.Fatalln(err)
		}
		err = keys.Public.Set(jwk.AlgorithmKey, jwa.RS256)
		if err != nil {
			log.Fatalln(err)
		}
		err = keys.Public.Set(jwk.KeyIDKey, "Public-key")
		if err != nil {
			log.Fatalln(err)
		}
	}
	err = saveKeysToDisk(rawRSAPrivateKey, &rawRSAPublicKey)
	if err != nil {
		log.Fatalln(err)
	}
	return keys
}

const (
	privateKeyFilename = "rsa_jwt"
	publicKeyFilename  = "rsa_jwt.pub"
)

// saveKeysToDisk creates two files to store the Private and Public part of an RSA key.
func saveKeysToDisk(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) error {
	// TODO: handle in case files already exists
	var privateKeyFile, publicKeyFile *os.File
	privateKeyFile, err := os.OpenFile(privateKeyFilename, os.O_RDWR|os.O_CREATE, 0600) // (rw-------)
	if err != nil {
		return err
	}
	publicKeyFile, err = os.OpenFile(publicKeyFilename, os.O_RDWR|os.O_CREATE, 0644) // (rw-r--r--)
	if err != nil {
		return err
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privateKeyBytes,
		},
	)
	publicKeyBytes := x509.MarshalPKCS1PublicKey(publicKey)
	pubKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	_, err = privateKeyFile.Write(privateKeyPem)
	if err != nil {
		return err
	}
	_, err = publicKeyFile.Write(pubKeyPem)
	return err
}
