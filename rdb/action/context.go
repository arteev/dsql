package action

import "sync"

//A Context - this is a context when action executing
type Context struct {
	mu    *sync.Mutex
	Snap  actionSnap
	addon map[string]interface{}
}

//NewContext - Create new Context
func NewContext() *Context {
	result := &Context{
		Snap: newSnap(),
		mu:   &sync.Mutex{},
	}
	result.addon = make(map[string]interface{})
	return result
}

//Set - set  the pair: key and value
func (c *Context) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.addon[key] = value
}

//Get - returns value(or nil) for the key
func (c *Context) Get(key string) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()
	result, ok := c.addon[key]
	if ok {
		return result
	}
	return nil
}

//GetDef - returns value or default for the key
func (c *Context) GetDef(key string, def interface{}) interface{} {
	val := c.Get(key)
	if val != nil {
		return val
	}
	return def
}

func (c *Context) inc(key string, delta interface{}) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.addon[key]
	if ok {
		switch delta.(type) {
		case int:
			val = val.(int) + delta.(int)
		case int64:
			val = val.(int64) + delta.(int64)
		}
	} else {
		val = delta
	}
	c.addon[key] = val
	return val
}

//IncInt - increment value for a key and retrun it
func (c *Context) IncInt(key string, delta int) int {
	return c.inc(key, delta).(int)
}

//IncInt64 - increment value for a key and retrun it
func (c *Context) IncInt64(key string, delta int64) int64 {
	return c.inc(key, delta).(int64)
}
