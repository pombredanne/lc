// SPDX-License-Identifier: GPL-3.0-only

package parsers

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

func getMd5Hash(content []byte) string {
	hasher := md5.New()
	hasher.Write(content)
	return hex.EncodeToString(hasher.Sum(nil))
}

func getSha1Hash(content []byte) string {
	hasher := sha1.New()
	hasher.Write(content)
	return hex.EncodeToString(hasher.Sum(nil))
}

func getSha256Hash(content []byte) string {
	hasher := sha256.New()
	hasher.Write(content)
	return hex.EncodeToString(hasher.Sum(nil))
}

func readFile(filepath string) []byte {
	// TODO only read as deep into the file as we need
	bytes, err := ioutil.ReadFile(filepath)

	if err != nil {
		fmt.Print(err)
	}

	return bytes
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// Given a list of license matches return a new list containing no duplicates
func uniqLicenseMatch(l []LicenseMatch) []LicenseMatch {
	m := make(map[LicenseMatch]bool)
	for _, s := range l {
		m[s] = true
	}
	result := make([]LicenseMatch, 0, len(m))
	for s := range m {
		result = append(result, s)
	}
	return result
}

// Returns true if a license list contains the license
func licenceListHasLicense(license LicenseMatch, licenseList []LicenseMatch) bool {
	for _, v := range licenseList {
		if v.LicenseId == license.LicenseId {
			return true
		}
	}

	return false
}

// Borrowed from https://github.com/cloudfoundry/bytefmt
// Apache-2.0 is compatible with GPL-3.0-only
// See https://apache.org/licenses/GPL-compatibility.html
// SPDX-License-Identifier: Apache-2.0
func bytesToHuman(bytes int) string {

	const TERABYTE = 1099511627776
	const GIGABYTE = 1073741824
	const MEGABYTE = 1048576
	const KILOBYTE = 1024

	unit := ""
	value := float64(bytes)

	switch {
	case bytes >= TERABYTE:
		unit = "T"
		value = value / TERABYTE
	case bytes >= GIGABYTE:
		unit = "G"
		value = value / GIGABYTE
	case bytes >= MEGABYTE:
		unit = "M"
		value = value / MEGABYTE
	case bytes >= KILOBYTE:
		unit = "K"
		value = value / KILOBYTE
	case bytes >= 1:
		unit = "B"
	case bytes == 0:
		return "0"
	}

	stringValue := fmt.Sprintf("%.1f", value)
	stringValue = strings.TrimSuffix(stringValue, ".0")
	return fmt.Sprintf("%s%s", stringValue, unit)
}
