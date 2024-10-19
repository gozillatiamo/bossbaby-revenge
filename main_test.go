package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBossBabyRevenge(t *testing.T) {
	t.Run("SRSSRRR: All shorts is revenged", func(t *testing.T) {
		result, err := CheckBossBaby("SRSSRRR")
		assert.NoError(t, err)
		assert.Equal(t, "Good boy", result)
	})
	t.Run("RSSRR: Initiate with revenged", func(t *testing.T) {
		result, err := CheckBossBaby("RSSRR")
		assert.NoError(t, err)
		assert.Equal(t, "Bad boy", result)
	})
	t.Run("SSSRRRRS: Remain a shot to revenge", func(t *testing.T) {
		result, err := CheckBossBaby("SSSRRRRS")
		assert.NoError(t, err)
		assert.Equal(t, "Bad boy", result)
	})
	t.Run("SRRSSR: Remain a shot to revenge", func(t *testing.T) {
		result, err := CheckBossBaby("SRRSSR")
		assert.NoError(t, err)
		assert.Equal(t, "Bad boy", result)
	})
	t.Run("SSRSRRR: All shorts is revenged but uncontinued", func(t *testing.T) {
		result, err := CheckBossBaby("SSRSRRR")
		assert.NoError(t, err)
		assert.Equal(t, "Good boy", result)
	})

	t.Run("SRSSRSSSRRSRRRRRSR: Ping with more Pong case", func(t *testing.T) {
		result, err := CheckBossBaby("SRSSRSSSRRSRRRRRSR")
		assert.NoError(t, err)
		assert.Equal(t, "Good boy", result)
	})
}
