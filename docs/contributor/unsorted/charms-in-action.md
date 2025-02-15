(charms-in-action)=
Charms in action
================

This document describes the behaviour of the go implementation of the unit
agent, whose behaviour differs in some respects from that of the python
implementation. This information is largely relevant to charm authors, and
potentially to developers interested in the unit agent.

Hooks
-----

A service unit's direct action is entirely defined by its charm's hooks. Hooks
are executable files in a charm's hooks directory; hooks with particular names
will be invoked by the juju unit agent at particular times, and thereby cause
changes to the world.

Whenever a hook-worthy event takes place, the unit agent tries to run the hook
with the appropriate name. If the hook doesn't exist, the agent continues
without complaint; if it does, it is invoked without arguments in a specific
environment, and its output is written to the unit agent's log. If it returns
a non-zero exit code, the agent enters an error state and awaits resolution;
otherwise it continues to process model changes as before.

In general, a unit will run hooks in a clear sequence, about which a number of
useful guarantees are made. All such guarantees come with the caveat that there
is [TODO: will be: `remove-unit --force`] a mechanism for forcible termination
of a unit, and that a unit so terminated will just stop, dead, and completely
fail to run anything else ever again. This shouldn't actually be a big deal in
practice.

Errors in hooks
---------------

Hooks should ideally be idempotent, so that they can fail and be re-executed
from scratch without trouble. As a hook author, you don't have complete control
over the times your hook might be stopped: if the unit agent process is killed
for any reason while running a hook, then when it recovers it will treat that
hook as having failed -- just as if it had returned a non-zero exit code -- and
request user intervention.

It is unrealistic to expect great sophistication on the part of the average user,
and as a charm author you should expect that users will attempt to re-execute
failed hooks before attempting to investigate or understand the situation. You
should therefore make every effort to ensure your hooks are idempotent when
aborted and restarted.

[TODO: I have a vague feeling that `juju resolved` actually defaults to "just
pretend the hook ran successfully" mode. I'm not sure that's really the best
default, but I'm also not sure we're in a position to change the UI that much.]

The most sophisticated charms will consider the nature of their operations with
care, and will be prepared to internally retry any operations they suspect of
having failed transiently, to ensure that they only request user intervention in
the most trying circumstances; and will also be careful to log any relevant
information or advice before signalling the error.

[TODO: I just thought; it would be really nice to have a juju-fail hook tool,
which would allow charm authors to explicity set the unit's error status to
something a bit more sophisticated than "X hook failed". Wishlist, really.]

Charm deployment
----------------

  * A charm is deployed into a directory that is entirely owned and controlled
    by juju.
  * At certain times, control of the directory is ceded to the charm (by
    running a hook) or to the user (by entering an error state).
  * At these times, and only at these times, should the charm directory be
    used by anything other than juju itself.

The most important consequence of this is that it is a mistake to conflate the
state of the charm with the state of the software deployed by the charm: it's
fine to store *charm* state in the charm directory, but the charm must deploy
its actual software elsewhere on the system.

To put it another way: deleting the charm directory should not impact the
software deployed by the charm in any way; and there is currently no mechanism
by which deployed software can safely feed information back into the charm
and/or expect that it will be acted upon in a timely way.

Execution environment
---------------------

Every hook is run in the deployed charm directory, in an environment with the
following characteristics:

  * $PATH is prefixed by a directory containing command line tools through
    which the hooks can interact with juju.
  * $CHARM_DIR holds the path to the charm directory.
  * $JUJU_UNIT_NAME holds the name of the local unit.
  * $JUJU_CONTEXT_ID, $JUJU_AGENT_SOCKET_NETWORK and $JUJU_AGENT_SOCKET_ADDRESS 
    are set (but should not be messed with: the command line tools won't work without them).
  * $JUJU_API_ADDRESSES holds a space separated list of juju API addresses.
  * $JUJU_MODEL_NAME holds the human friendly name of the current model.
  * $JUJU_PRINCIPAL_UNIT holds the name of the principal unit if the current unit is a subordinate.

Hook tools
----------

All hooks can directly use the following tools:

  * juju-log (write arguments direct to juju's log (potentially redundant, hook
    output is all logged anyway, but --debug may remain useful))
  * unit-get (returns the local unit's private-address or public-address)
  * open-port (marks the supplied port/protocol as ready to open when the
    application is exposed)
  * close-port (reverses the effect of open-port)
  * config-get (get current service configuration values)
  * relation-get (get the settings of some related unit)
  * relation-set (write the local unit's relation settings)
  * relation-ids (list all relations using a given charm relation)
  * relation-list (list all units of a related application)
  * relation-model-get (get details about the model hosting a related application)
  * storage-add (add storage instances)
  * storage-get (get storage instance values)
  * status-get (get unit workload status information)
  * status-set (set unit workload status information)

Within the context of a single hook execution, the above tools present a
sandboxed view of the system with the following properties:

  * Any data retrieved corresponds to the real value of the underlying state at
    some point in time.
  * Once state data has been observed within a given hook execution, further
    requests for the same data will produce the same results, unless that data
    has been explicitly changed with relation-set.
  * Data changed by relation-set is only written to global state when the hook
    completes without error; changes made by a failing hook will be discarded
    and never observed by any other part of the system.
  * Not actually sandboxed: open-port and close-port operate directly on state.
    [TODO: lp:1089304 - might be a little tricky.]

Hook kinds
----------

There are 5 `unit hooks` with predefined names that can be implemented by any
charm:

  * install
  * config-changed
  * start
  * upgrade-charm
  * stop

For every relation defined by a charm, an additional 4 `relation hooks` can be
implemented, named after the charm relation:

  * <name>-relation-joined
  * <name>-relation-changed
  * <name>-relation-departed
  * <name>-relation-broken

Unit hooks
----------

The `install` hook always runs once, and only once, before any other hook.

The `config-changed` hook always runs once immediately after the install hook,
and likewise after the upgrade-charm hook. It also runs whenever the service
configuration changes, and when recovering from transient unit agent errors.

The `start` hook always runs once immediately after the first config-changed
 hook; there are currently no other circumstances in which it will be called,
but this may change in the future.

The `upgrade-charm` hook always runs once immediately after the charm directory
contents have been changed by an unforced charm upgrade operation, and *may* do
so after a forced upgrade; but will *not* be run after a forced upgrade from an
existing error state. (Consequently, neither will the config-changed hook that
would ordinarily follow the upgrade-charm.)

The `stop` hook is the last hook to be run before the unit is destroyed. In the
future, it may be called in other situations.

In normal operation, a unit will run at least the install, start, config-changed
and stop hooks over the course of its lifetime.

It should be noted that, while all hook tools are available to all hooks, the
relation-* tools are not useful to the install, start, and stop hooks; this is
because the first two are run before the unit has any opportunity to participate
in any relations, and the stop hooks will not be run while the unit is still
participating in one.

Relation hooks
--------------

For each charm relation, any or all of the 4 relation hooks can be implemented.
Relation hooks operate in an environment slightly different to that of unit
hooks, in the following ways:

  * JUJU_RELATION is set to the name of the charm relation. This is of limited
    value, because every relation hook already "knows" what charm relation it
    was written for; that is, in the "foo-relation-joined" hook, JUJU_RELATION
    is "foo".
  * JUJU_RELATION_ID is more useful, because it serves as unique identifier for
    a particular relation, and thereby allows the charm to handle distinct
    relations over a single endpoint. In hooks for the "foo" charm relation,
    JUJU_RELATION_ID always has the form "foo:<id>", where id uniquely but
    opaquely identifies the runtime relation currently in play.
  * The relation-* hook tools, which ordinarily require that a relation be
    specified, assume they're being called with respect to the current
    relation. The default can of course be overridden as usual.

Furthermore, all relation hooks except relation-broken are notifications about
some specific unit of a related service, and operate in an environment with the
following additional properties:

  * JUJU_REMOTE_UNIT is set to the name of the current related unit.
  * The relation-get hook tool, which ordinarily requires that a related unit
    be specified, assumes that it is being called with respect to the current
    related unit. The default can of course be overridden as usual.

For every relation in which a unit partcipates, hooks for the appropriate charm
relation are run according to the following rules.

The "relation-joined" hook always runs once when a related unit is first seen.

The "relation-changed" hook for a given unit always runs once immediately
following the relation-joined hook for that unit, and subsequently whenever
the related unit changes its settings (by calling relation-set and exiting
without error). Note that "immediately" only applies within the context of
this particular runtime relation -- that is, when "foo-relation-joined" is
run for unit "bar/99" in relation id "foo:123", the only guarantee is that
the next hook to be run *in relation id "foo:123"* will be "foo-relation-changed"
for "bar/99". Unit hooks may intervene, as may hooks for other relations,
and even for other "foo" relations.

The "relation-departed" hook for a given unit always runs once when a related
unit is no longer related. After the "relation-departed" hook has run, no
further notifications will be received from that unit; however, its settings
will remain accessible via relation-get for the complete lifetime of the
relation.

The "relation-broken" hook is not specific to any unit, and always runs once
when the local unit is ready to depart the relation itself. Before this hook
is run, a relation-departed hook will be executed for every unit known to be
related; it will never run while the relation appears to have members, but it
may be the first and only hook to run for a given relation. The stop hook will
not run while relations remain to be broken.

Relations in depth
------------------

A unit's `scope` consists of the group of units that are transitively connected
to that unit within a particular relation. So, for a globally-scoped relation,
that means every unit of each service in the relation; for a locally-scoped
relation, it means only those sets of units which are deployed alongside one
another.  That is to say: a globally-scoped relation has a single unit scope,
whilst a locally-scoped relation has one for each principal unit.

When a unit becomes aware that it is a member of a relation, its only self-
directed action is to `join` its scope within that relation. This involves two
steps:

  * Write initial relation settings (just one value, "private-address"), to
    ensure that they will be available to observers before they're triggered
    by the next step;
  * Signal its existence, and role in the relation, to the rest of the system.

The unit then starts observing and reacting to any other units in its scope
which are playing a role in which it is interested. To be specific:

  * Each provider unit observes every requirer unit
  * Each requirer unit observes every provider unit
  * Each peer unit observes every other peer unit

Now, suppose that some unit as the very first unit to join the relation; and
let's say it's a requirer. No provider units are present, so no hooks will fire.
But, when a provider unit joins the relation, the requirer and provider become
aware of each other almost simultaneously. (Similarly, the first two units in a
peer relation become aware of each other almost simultaneously.)

So, concurrently, the units on each side of the relation run their relation-joined
and relation-changed hooks with respect to their counterpart. The intent is that
they communicate appropriate information to each other to set up some sort of
connection, by using the relation-set and relation-get hook tools; but neither
unit is safe to assume that any particular setting has yet been set by its
counterpart.

This sounds kinda tricky to deal with, but merely requires suitable respect for
the relation-get tool: it is important to realise that relation-get is *never*
guaranteed to contain any values at all, because we have decided that it's
perfectly legitimate for a unit to delete its own private-address value.

[TODO: There is a school of thought that maintains that we should add an
independent "juju-private-address" setting that *is* guaranteed, but for now
the reality is that relation-get can *always* fail to produce any given value.
However, in the name of sanity, it's probably reasonable to treat a missing
private-address as an error, and assume that `relation-get private-address` is
always safe. For all other values, we must operate with the understanding that
relation-get can always fail.]

In one specific kind of hook, this is easy to deal with. A relation-changed hook
can always exit without error when the current remote unit is missing data,
because the hook is guaranteed to be run again when that data changes -- and,
assuming the remote unit is running a charm that agrees on how to implement the
interface, the data *will* change and the hook *will* be run again.

In *all* other cases -- unit hooks, relation hooks for a different relation,
relation hooks for a different remote unit in the same relation, and even
relation hooks other than -changed for the *same* remote unit -- there is no
such guarantee. These hooks all run on their own schedule, and there is no
reason to expect them to be re-run on a predictable schedule, or in some cases
ever again.

This means that all such hooks need to be able to handle missing relation data,
and to complete successfully; they mustn't fail, because the user is powerless
to resolve the situation, and they can't even wait for state to change, because
they all see their own sandboxed composite snapshot of fairly-recent state,
which never changes.

So, outside a vey narrow range of circumstances, relation-get should be treated
with particular care. The corresponding advice for relation-set is very simple
by comparison: relation-set should be called early and often. Because the unit
agent serializes hook execution, there is never any danger of concurrent changes
to the data, and so a null setting change can be safely ignored, and will not
cause other units to react.

Departing relations
-------------------

A unit will depart a relation when either the relation or the unit itself is
marked for termination. In either case, it follows the same sequence:

  * For every known related unit -- those which have joined and not yet
    departed -- run the relation-departed hook.
  * Run the relation-broken hook.
  * `depart` from its scope in the relation.

The unit's departure from its scope will in turn be detected by units of the
related service, and cause them to run relation-departed hooks. A unit's
relation settings persist beyond its own departure from the relation; the
final unit to depart a relation marked for termination is responsible for
destroying the relation and all associated data.

Debugging charms
----------------

Facilities are currently not good.

  * juju ssh
  * juju debug-hooks [TODO: not implemented]
  * juju debug-log [TODO: not implemented]

