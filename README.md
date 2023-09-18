<h3 align="center">crypty - golang cryptography library</h3>

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
</div>

###

## features/functionality

- **library that provides cryptographic functions and algorithms, 
including encryption, decryption, hashing, and digital signatures, generation**
  
- **encryption**:
  
	- **EncryptAES(key []byte, plaintext []byte) ([]byte, error)**
   
	    - encrypts the provided plaintext using the AES encryption algorithm
	    - utilizes AES in CBC mode with a random initialization vector (IV)
       
	- **EncryptDES(key []byte, plaintext []byte) ([]byte, error)**
   
		- encrypts the provided plaintext using the DES encryption algorithm
		- supports DES encryption for data blocks of the correct size
    
	- **Encrypt3DES(key []byte, plaintext []byte) ([]byte, error**)
   
	    - encrypts the provided plaintext using the 3DES (Triple DES) encryption algorithm
	    - requires a 24-byte (192-bit) key

- **decryption**:
  
	- **DecryptAES(key []byte, ciphertext []byte) ([]byte, error)**
   
		- decrypts AES-encrypted ciphertext using the provided key
    
	- **DecryptDES(key []byte, ciphertext []byte) ([]byte, error)**
   
		- decrypts DES-encrypted ciphertext using the provided key
    
	- **Decrypt3DES(key []byte, ciphertext []byte) ([]byte, error)**
   
		- decrypts 3DES (Triple DES)-encrypted ciphertext using the provided key

- **hashing**:
  
	- **CalculateSHA256(data []byte) []byte**
   
		- calculates the SHA-256 hash of the input data
    
	- **CalculateSHA224(data []byte) []byte**
   
		- calculates the SHA-224 hash of the input data
    
	- **CalculateSHA3_256(data []byte) []byte**
   
		- calculates the SHA3-256 hash of the input data
    
	- **CalculateSHA1(data []byte) []byte**
   
		- calculates the SHA-1 hash of the input data
    
	- **CalculateMD5(data []byte) []byte**
   
		- calculates the MD5 hash of the input data
    
	- **CalculateSHA512(data []byte) []byte**
   
		- calculates the SHA-512 hash of the input data
    
	- **CalculateSHA384(data []byte) []byte**
   
		- calculates the SHA-384 hash of the input data

- **randomization**:
  
    - **GenerateRandomBytes(length int) ([]byte, error)**
      
		- generates a specified number of random bytes
		- uses the crypto/rand package for secure random byte generation

	- **GenerateRandomHex(length int) (string, error)**
   
		- generates a random hexadecimal string of the specified length
		- utilizes GenerateRandomBytes internally and encodes the result in hexadecimal

- **digital_signatures**:

	- **GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error)****
 
	    - generates an RSA key pair for digital signatures with the specified number of bits
	    - returns a private key and its corresponding public key
       
	- **SignDataWithRSA(data []byte, privateKey *rsa.PrivateKey) ([]byte, error)***
   
		- signs a byte array of data using an RSA private key
		- uses the PKCS1v15 signature scheme with SHA-256 hash
    
	- **VerifySignatureWithRSA(data, signature []byte, publicKey *rsa.PublicKey) error***
   
	    - verifies a digital signature using an RSA public key
	    - validates the signature against the data using the PKCS1v15 signature scheme and SHA-256 hash
       
	- **ExportRSAPrivateKeyToPEM(privateKey *rsa.PrivateKey) (string, error)***
   
		- exports an RSA private key to a PEM-encoded string
    
	 - **ExportRSAPublicKeyToPEM(publicKey *rsa.PublicKey) (string, error)***
    
		- exports an RSA public key to a PEM-encoded string
    
	 - **ImportRSAPrivateKeyFromPEM(keyPEM string) (*rsa.PrivateKey, error)***
    
		- imports an RSA private key from a PEM-encoded string
    
	- **ImportRSAPublicKeyFromPEM(keyPEM string) (*rsa.PublicKey, error)***
   
		- imports an RSA public key from a PEM-encoded string
 
## project structure:

```go
.
├── decrypt
│   └── decrypt.go
├── digital_signatures
│   └── digital_signatures.go
├── encrypt
│   └── encrypt.go
├── go.mod
├── go.sum
├── hashing
│   └── hashing.go
├── random
│   └── random.go
└── README.md
```

## installation

```shell
go get github.com/kenjitheman/crypty
```

## contributing

- pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change

## license

- [MIT](https://choosealicense.com/licenses/mit/)
