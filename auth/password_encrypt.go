package auth

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
)

/*Our passwords are encrypted under AES cypher: https://gist.github.com/jpillora/cb46d183eca0710d909a*/

const key = "1234567 1234567 1234567 1234567k"

//PassEncrypt is used to encrypt the user password before it goes to de db.
func PassEncrypt(pass string) (string, error) {
	ciphertext, err := encrypt([]byte(pass), key)
	if err != nil {
		return "", err
	}
	return string(ciphertext), nil
}

//DecryptPassword it is the routine to decrypt passwords
func DecryptPassword(encryptedPass string) (string, error) {
	plaintext, err := decrypt([]byte(encryptedPass), key)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

/*ComparePasswords receives an encrypted password and a unencrypted password and compares them,
if they match returns true and nil, if there's an error decrypting returns false and the error, and if they not match, returns false and nil*/
func ComparePasswords(passNotEncrypted, passEncrypted []byte) (bool, error) {
	pass, err := decrypt(passEncrypted, key)

	if err != nil {
		return false, err
	}

	if bytes.Compare([]byte(passNotEncrypted), pass) == 0 {
		return true, nil
	}

	return false, nil
}

//createHash receives a passphrase, key or any string, hash it, then return the hash as a hexadecimal value
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) ([]byte, error) {
	//First we create a new block cipher based on the hashed passphrase(our key)
	block, err := aes.NewCipher([]byte(createHash(passphrase)))
	if err != nil {
		return nil, err
	}
	//Then we want to wrap it in Galois Counter Mode (GCM) with a standard nonce length (nonce = number that can be only used once)
	//Galois/Counter Mode (GCM) is a mode of operation for symmetric-key cryptographic block ciphers
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize()) //the nonce used for decryption must be the same nonce used for encryption.

	_, errorIo := io.ReadFull(rand.Reader, nonce)

	if errorIo != nil {
		return nil, errorIo
	}

	//to make sure our decryption nonce matches the encryption nonce, we will prepend the nonce
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func decrypt(data []byte, passphrase string) ([]byte, error) {
	key := []byte(createHash(passphrase))
	//we create a new block cipher using a hashed passphrase.
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}
	//we wrap the block cipher in Galois Counter Mode
	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return nil, err
	}
	//then get the nonce size
	nonceSize := gcm.NonceSize()
	// we need to separate the nonce and the encrypted data.
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	//we can decrypt the data and return it as plaintext
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)

	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
