// Package clustersrv holds the handlers and related functionality for running a single, in-memory riak cluster coordination server.
//
// The server has two endpoints: /start and /end/:lock_id. Essentially, it's a network-attached in-memory lock server.
// Currently its only use is to coordinate 'riak-admin cluster plan' and 'riak-admin cluster commit' calls so that no more than one is in flight at once
package clustersrv
