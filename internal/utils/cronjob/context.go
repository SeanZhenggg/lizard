package cronjob

import (
	"context"
	"sync"
)

type FuncJob func(ctx *Context)
type handlerChain []FuncJob

type Context struct {
	context.Context
	mu           sync.RWMutex
	param        map[string]any
	index        int
	handlerChain handlerChain
}

func (c *Context) reset(handlers handlerChain) {
	c.param = nil
	c.handlerChain = handlers
	c.index = -1
	c.Context = context.Background()
}
func (c *Context) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.param == nil {
		c.param = make(map[string]any)
	}

	c.param[key] = value
}

// Get returns the value for the given key, ie: (value, true).
// If the value does not exist it returns (nil, false)
func (c *Context) Get(key string) (value any, exists bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists = c.param[key]
	return
}

func (c *Context) Next() {
	c.index++
	for c.index < len(c.handlerChain) {
		c.handlerChain[c.index](c)
		c.index++
	}
}
