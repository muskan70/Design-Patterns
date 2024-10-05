package main

type Context struct {
	hashmap map[string]int
}

func (c *Context) put(key string, val int) {
	c.hashmap[key] = val
}

func (c *Context) get(key string) int {
	return c.hashmap[key]
}
