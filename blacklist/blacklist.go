package blacklist

import (
	"sync"
	"time"
)

// Blacklist manages blacklisted JWT tokens.
type Blacklist struct {
	mu     sync.RWMutex
	tokens map[string]time.Time // token string -> expiry time
}

// NewBlacklist creates a new Blacklist instance.
func NewBlacklist() *Blacklist {
	return &Blacklist{
		tokens: make(map[string]time.Time),
	}
}

// Add adds a token to the blacklist until its expiry.
func (b *Blacklist) Add(token string, expiry time.Time) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.tokens[token] = expiry
}

// IsBlacklisted checks if a token is blacklisted and not yet expired.
func (b *Blacklist) IsBlacklisted(token string) bool {
	b.mu.RLock()
	expiry, exists := b.tokens[token]
	b.mu.RUnlock()
	if !exists {
		return false
	}
	// Remove expired tokens
	if time.Now().After(expiry) {
		b.mu.Lock()
		delete(b.tokens, token)
		b.mu.Unlock()
		return false
	}
	return true
}

// Cleanup removes expired tokens from the blacklist.
func (b *Blacklist) Cleanup() {
	b.mu.Lock()
	defer b.mu.Unlock()
	now := time.Now()
	for token, expiry := range b.tokens {
		if now.After(expiry) {
			delete(b.tokens, token)
		}
	}
}
