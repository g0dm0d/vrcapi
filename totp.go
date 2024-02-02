package vrcapi

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"time"
)

func generateTOTP(secret string) (string, error) {
	// Decode the secret from base32
	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(secret)
	if err != nil {
		return "", err
	}

	// Get the current time in seconds
	currentTime := time.Now().Unix()

	// TOTP time step duration (30 seconds)
	timeStep := int64(30)

	// Calculate the counter based on the current time
	counter := uint64(currentTime / timeStep)

	// Convert the counter to a byte slice
	counterBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(counterBytes, counter)

	// Calculate the HMAC-SHA1 using the key and counter
	h := hmac.New(sha1.New, key)
	h.Write(counterBytes)
	hash := h.Sum(nil)

	// Get the offset from the last 4 bits of the hash
	offset := hash[len(hash)-1] & 0xf

	// Get the 4 bytes at the offset
	truncatedHash := hash[offset : offset+4]

	// Set the most significant bit to 0
	truncatedHash[0] &= 0x7f

	// Convert the 4 bytes to a uint32
	otp := binary.BigEndian.Uint32(truncatedHash)

	// Calculate the final 6-digit OTP
	otp = otp % 1000000

	// Format the OTP as a string with leading zeros if necessary
	otpString := fmt.Sprintf("%06d", otp)

	return otpString, nil
}
