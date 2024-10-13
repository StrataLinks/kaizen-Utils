package drivers

import (
	"log"

	scylla "github.com/gocql/gocql"
)

func ProvideScyllaDatabaseClient(address, keyspace, username, password string) (*scylla.Session, error) {
	return providerScyllaDatabaseClient(address, keyspace, username, password)
}

func providerScyllaDatabaseClient(address, keyspace, username, password string) (*scylla.Session, error) {
	cluster := scylla.NewCluster(address) // example IP, use your Scylla cluster IPs
	cluster.Keyspace = keyspace
	cluster.Consistency = scylla.Quorum
	cluster.ProtoVersion = 4
	cluster.Authenticator = scylla.PasswordAuthenticator{
		Username: username,
		Password: password,
	}

	// Create a session
	session, err := cluster.CreateSession()
	if err != nil {
		log.Printf("\nUnable to connect to Scylla: %v\n", err)
		log.Fatalf("Unable to connect to Scylla: %v", err)
	}

	return session, nil
}
