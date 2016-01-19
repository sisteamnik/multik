package multik

type Filter func(c *Controller, filterChain []Filter)
