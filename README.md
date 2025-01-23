Goal: Be a beast outperforming everything that comes in its way

# TODO:

## Core engine changes:
- Optimize for speed

## New feats:
- key expiry event notifs(tricky)

Problem:

- we are relying on usedCapacity of SM to trigger upgrades and downgrades, but it is not reliable at all

solution: make the totalCapacity and usedCapacity a *int from int

- we need a non blocking EventQueue so core engine connection reads don't get blocked, such a massively buffered channel is being created for a global queue, takes a SHIT LOAD of memory on initialization, also in some case like a connection disconnecting, their secondary write channel just keeps getting bigger and bigger without getting drained, could be handled separately when they disconnect, but still stays a big issue for the core engine

### commands supported rn:
```
SET <key> <value>
SET <key> <value> EX <expiry_time_in_seconds>

GET <key>

SUBSCRIBE <key_name>
```