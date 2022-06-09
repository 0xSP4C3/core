package utils

import (
	"fmt"

	"github.com/0xsp4c3/core/pkg/repository"
)

// GetCredentialsByRole func for getting credentials from a role name.
func GetCredentialsByRole(role string) ([]string, error) {
	// Define credentials variable.
	var credentials []string

	// Switch given role.
	switch role {
	case repository.AdminRoleName:
		// Admin credentials (all access).
		credentials = []string{
            repository.CoinCreateCredential,
            repository.CoinUpdateCredential,
            repository.CoinDeleteCredential,
            repository.ExchangeCreateCredential,
            repository.ExchangeUpdateCredential,
            repository.ExchangeDeleteCredential,
		}
	case repository.ModeratorRoleName:
		// Moderator credentials (only book creation and update).
		credentials = []string{
            repository.CoinCreateCredential,
            repository.CoinUpdateCredential,
            repository.ExchangeCreateCredential,
            repository.ExchangeUpdateCredential,
		}
	case repository.UserRoleName:
		// Simple user credentials (only book creation).
		credentials = []string{
            repository.CoinCreateCredential,
            repository.ExchangeCreateCredential,
		}
	default:
		// Return error message.
		return nil, fmt.Errorf("role '%v' does not exist", role)
	}

	return credentials, nil
}
