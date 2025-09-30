/*
Copyright © 2024 Savely Krasovsky <savely@krasovs.ky>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// gendidCmd represents the gendid command
var gendidCmd = &cobra.Command{
	Use:   "gendid",
	Short: "Generates did.json for to serve",
	RunE:  generateDid,
}

var (
	handle   string
	diddomain string
	pubkey   string
	hostname string
)

func init() {
	rootCmd.AddCommand(gendidCmd)

	gendidCmd.Flags().StringVar(&handle, "handle", "", "Pass your handle here")
	gendidCmd.MarkFlagRequired("handle")
	gendidCmd.Flags().StringVar(&diddomain, "diddomain", "", "Pass the domain of your did.")
	gendidCmd.MarkFlagRequired("diddomain")
	gendidCmd.Flags().StringVar(&pubkey, "pubkey", "", "Pass public key from pubkey command here")
	gendidCmd.MarkFlagRequired("pubkey")
	gendidCmd.Flags().StringVar(&hostname, "hostname", "", "Pass hostname of PDS you want to register")
	gendidCmd.MarkFlagRequired("hostname")
}

func generateDid(cmd *cobra.Command, args []string) error {
	did := &DID{
		Context: []string{
			"https://www.w3.org/ns/did/v1",
			"https://w3id.org/security/multikey/v1",
			"https://w3id.org/security/suites/secp256k1-2019/v1",
		},
		Id:          "did:web:" + diddomain,
		AlsoKnownAs: []string{"at://" + handle},
		VerificationMethod: []*VerificationMethod{{
			ID:                 "did:web:" + diddomain + "#atproto",
			Type:               "Multikey",
			Controller:         "did:web:" + diddomain,
			PublicKeyMultibase: pubkey,
		}},
		Service: []*Service{{
			ID:              "#atproto_pds",
			Type:            "AtprotoPersonalDataServer",
			ServiceEndpoint: "https://" + hostname,
		}},
	}

	b, err := json.MarshalIndent(did, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

type DID struct {
	Context            []string              `json:"@context"`
	Id                 string                `json:"id"`
	AlsoKnownAs        []string              `json:"alsoKnownAs"`
	VerificationMethod []*VerificationMethod `json:"verificationMethod"`
	Service            []*Service            `json:"service"`
}

type VerificationMethod struct {
	ID                 string `json:"id"`
	Type               string `json:"type"`
	Controller         string `json:"controller"`
	PublicKeyMultibase string `json:"publicKeyMultibase"`
}

type Service struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	ServiceEndpoint string `json:"serviceEndpoint"`
}
