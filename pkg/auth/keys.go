package auth

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

// jwkSet holds the public and private rsa keys.
type jwkSet struct {
	public  jwk.Key
	private jwk.Key
}

var (
	keys *jwkSet

	bitSize = flag.Int("bits", 3072, "size in bits of the rsa key")
)

// getKeys singleton function to initialize the [keys] variable, or to return its
// current value.
func getKeys() *jwkSet {
	if keys.private != nil && keys.public != nil {
		return keys
	}
	// create a new RSA Key
	rawRSAPrivateKey, err := rsa.GenerateKey(rand.Reader, *bitSize)
	if err != nil {
		log.Fatalln(err)
	}
	rawRSAPublicKey := rawRSAPrivateKey.PublicKey

	{ // set fields for private key
		keys.private, err = jwk.FromRaw(rawRSAPrivateKey)
		if err != nil {
			log.Fatalln(err)
		}
		err = keys.private.Set(jwk.AlgorithmKey, jwa.RS256)
		if err != nil {
			log.Fatalln(err)
		}
		err = keys.private.Set(jwk.KeyIDKey, "private-key")
	}

	{ // set fields for public key
		keys.public, err = jwk.FromRaw(rawRSAPublicKey)
		if err != nil {
			log.Fatalln(err)
		}
		err = keys.public.Set(jwk.AlgorithmKey, jwa.RS256)
		if err != nil {
			log.Fatalln(err)
		}
		err = keys.public.Set(jwk.KeyIDKey, "public-key")
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

// saveKeysToDisk creates two files to store the private and public part of a RSA key.
func saveKeysToDisk(privKey *rsa.PrivateKey, pubKey *rsa.PublicKey) error {
	// TODO: handle in case files already exists
	priv, err := os.OpenFile(privateKeyFilename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600) // (rw-------)
	if err != nil {
		return err
	}
	pub, err := os.OpenFile(publicKeyFilename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644) // (rw-r--r--)
	if err != nil {
		return err
	}

	privKeyBytes := x509.MarshalPKCS1PrivateKey(privKey)
	privKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privKeyBytes,
		},
	)
	publicKeyBytes := x509.MarshalPKCS1PublicKey(pubKey)
	pubKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	_, err = priv.Write(privKeyPem)
	_, err = pub.Write(pubKeyPem)
	return err
}
