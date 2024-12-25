// Description: This file contains the cache implementation for the search results. Please check documentation for more details.

package cache

import (
	// packages
	"IGI_API/internal/models"
	"IGI_API/internal/utils"

	// modules
	"fmt"
	"strings"
	"sync"
	"time"
)

// Represents a single cache item with data and expiration time.
type CacheItem struct {
	Data      models.SearchResults
	ExpiresAt time.Time
}

// Represents the cache itself. It contains a mutex for synchronization and a map to store cache items.
// Map key is a string built by search query values and the value it holds is a CacheItem.
type Cache struct {
	mutex sync.Mutex
	store map[string]CacheItem
}

// singleton instance of the cache.
var (
	instance *Cache
	lock     sync.Mutex
)

// Init a new Cache instance by sişngleton pattern.
func NewCache() *Cache {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Cache{
				store: make(map[string]CacheItem),
			}
		}
	}
	return instance
}

// Write to cache
func (cache *Cache) Set(key string, data models.SearchResults, ttl time.Duration) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	// store the data in the cache with the key and expiration time. (15 minutes)
	cache.store[key] = CacheItem{
		Data:      data,
		ExpiresAt: time.Now().Add(ttl),
	}
	utils.Logger("info", "Cache", 0, fmt.Sprintf("key %s is cached!", key))
}

// Read from cache
func (cache *Cache) Get(key string) (CacheItem, bool) {
	utils.Logger("info", "Cache", 0, fmt.Sprintf("Reading %s from cache...", key))

	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	// Check if the key exists in the map.
	item, exists := cache.store[key]

	// key does not exist or the item has expired. Inform invoker and clear the cache.
	if !exists || time.Now().After(item.ExpiresAt) {
		delete(cache.store, key)
		return CacheItem{}, false
	}

	// Data found in cache. Return the data and inform invoker.
	return item, true
}

// Generates a key with query word and selected sourced to identify the cache item
func (cache *Cache) KeyGen(keyword string, sources []string) string {
	sourceKey := strings.Join(sources, ",")
	return fmt.Sprintf("search:%s:%s", keyword, sourceKey)
}
