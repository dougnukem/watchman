package watchman

// FlushSubscriptions is the return object of the FlushSubscriptions call
type FlushSubscriptions struct {
}

// FlushSubscriptionsOptions are optional arguments to pass into FlushSubscriptions
type FlushSubscriptionsOptions struct {
	SyncTimeout   int      `bser:"sync_timeout"`
	Subscriptions []string `bser:"subscriptions"`
}

// FlushSubscriptions flushes buffered updates to subscriptions associated with the current session, guaranteeing that they are up-to-date as of the time Watchman received the flush-subscriptions command.
// https://facebook.github.io/watchman/docs/cmd/flush-subscriptions.html
func (c *Client) FlushSubscriptions(path string, opts FlushSubscriptionsOptions) (*FlushSubscriptions, error) {
	return nil, nil
}
