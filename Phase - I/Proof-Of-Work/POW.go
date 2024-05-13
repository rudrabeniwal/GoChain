package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const targetBits = 20 //In Bitcoin, “target bits” is the block header storing the difficulty at which the block was mined.
//Increasing the targetBits increase the difficulty in mining

type ProofOfWork struct {
	block  *Block
	target *big.Int //24 is an arbitrary number, our goal is to have a target that takes less than 256 bits in memory. And we want the difference to be significant enough, but not too big, because the bigger the difference the more difficult it’s to find a proper hash.
}

//You can think of a target as the upper boundary of a range: if a number (a hash) is lower than the boundary, it’s valid, and vice versa. Lowering the boundary will result in fewer valid numbers, and thus, more difficult work required to find a valid one.

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, 256-targetBits) //Lsh sets z = x << n and returns z.
	//Lsh(x *big.Int, n uint) *big.Int

	pow := &ProofOfWork{b, target}
	return pow
}

// Now we need data to hash
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.block.PrevBlockHash,
		pow.block.Data,
		IntToHex(pow.block.Timestamp),
		IntToHex(int64(targetBits)),
		IntToHex(int64(nonce)),
	}, []byte{})

	return data

}

func (pow *ProofOfWork) Run() (int, []byte) {

	//First, we initialize variables: hashInt is the integer representation of hash; nonce is the counter. Next, we run an “infinite” loop: it’s limited by maxNonce, which equals to math.MaxInt64; this is done to avoid a possible overflow of nonce
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	maxNonce := math.MaxInt64
	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

	/*In the loop we:
	1.Prepare data.
	2.Hash it with SHA-256.
	3.Convert the hash to a big integer.
	4.Compare the integer with the target. */
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

// Validate POW
func (pow *ProofOfWork) validate() bool {
	var hashInt big.Int

	data := pow.prepareData(int(pow.block.Nonce))
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
