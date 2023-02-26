package integration_test

import (
	"os"
	"testing"

	"github.com/letsencrypt/pebble/v2/test"
)

func TestCAALogChecker(t *testing.T) {
	os.Setenv("DIRECTORY", "https://localhost:14000/directory")
	c, err := makeClient()
	test.AssertNotError(t, err, "makeClient failed")

	domains := []string{random_domain()}
	t.Log("domains to issue: ", domains)
	result, err := authAndIssue(c, nil, domains)
	test.AssertNotError(t, err, "authAndIssue failed")
	test.AssertEquals(t, result.Order.Status, "valid")
	test.AssertEquals(t, len(result.Order.Authorizations), 1)

	// for i, cert := range result.certs {
	// 	var buf bytes.Buffer

	// 	err = pem.Encode(&buf, &pem.Block{
	// 		Type:  "RSA PRIVATE KEY",
	// 		Bytes: x509.MarshalPKCS1PrivateKey(key),
	// 	})
	// }
}
