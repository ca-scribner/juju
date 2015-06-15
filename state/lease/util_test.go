// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package lease_test

import (
	"time"

	jujutxn "github.com/juju/txn"
	"gopkg.in/mgo.v2"

	"github.com/juju/juju/mongo"
)

// Clock exposes time via Now, and can be controlled via Advance. It can be
// configured to Advance automatically whenever Now is called.
type Clock struct {
	now  time.Time
	step time.Duration
}

func NewClock(now time.Time, step time.Duration) *Clock {
	return &Clock{now, step}
}

// Now is part of the lease.Clock interface.
func (clock *Clock) Now() time.Time {
	defer clock.Advance(clock.step)
	return clock.now
}

func (clock *Clock) Set(now time.Time) {
	clock.now = now
}

func (clock *Clock) Advance(duration time.Duration) {
	clock.now = clock.now.Add(duration)
}

// Mongo exposes database operations. It uses a real database -- we can't mock
// mongo out, we need to check it really actually works -- but it's good to
// have the runner accessible for adversarial transaction tests.
type Mongo struct {
	database *mgo.Database
	runner   jujutxn.Runner
}

func NewMongo(database *mgo.Database) *Mongo {
	return &Mongo{
		database: database,
		runner: jujutxn.NewRunner(jujutxn.RunnerParams{
			Database: database,
		}),
	}
}

// GetCollection is part of the lease.Mongo interface.
func (m *Mongo) GetCollection(name string) (*mgo.Collection, func()) {
	return mongo.CollectionFromName(m.database, name)
}

// RunTransaction is part of the lease.Mongo interface.
func (m *Mongo) RunTransaction(getTxn jujutxn.TransactionSource) error {
	return m.runner.Run(getTxn)
}
