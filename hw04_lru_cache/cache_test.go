package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("check expel logic when going over capacity", func(t *testing.T) {
		c := NewCache(3)
		for i := range 7 {
			key := Key(strconv.Itoa(i))
			wasInCache := c.Set(key, i)
			require.False(t, wasInCache)
		}

		for i := range 4 {
			key := Key(strconv.Itoa(i))
			_, ok := c.Get(key)
			require.False(t, ok)
		}
	})

	t.Run("check expel logic when items have been swapped", func(t *testing.T) {
		c := NewCache(3)
		for i := range 3 {
			key := Key(strconv.Itoa(i))
			wasInCache := c.Set(key, i)
			require.False(t, wasInCache)
		}

		val, ok := c.Get("2")
		require.True(t, ok)
		require.Equal(t, 2, val)

		val, ok = c.Get("3")
		require.False(t, ok)
		require.Equal(t, nil, val)

		val, ok = c.Get("0")
		require.True(t, ok)
		require.Equal(t, 0, val)

		wasInCache := c.Set("10", 10)
		require.False(t, wasInCache)

		_, ok = c.Get("0")
		require.True(t, ok)

		_, ok = c.Get("2")
		require.True(t, ok)

		_, ok = c.Get("1")
		require.False(t, ok)

		_, ok = c.Get("10")
		require.True(t, ok)
	})

	// Check whether Clear func removes all elements

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)
		for i := range 3 {
			key := Key(strconv.Itoa(i))
			wasInCache := c.Set(key, i)
			require.False(t, wasInCache)
		}

		c.Clear()

		for i := range 3 {
			key := Key(strconv.Itoa(i))
			_, ok := c.Get(key)
			require.False(t, ok)
		}
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
