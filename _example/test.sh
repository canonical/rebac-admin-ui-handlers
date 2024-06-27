#!/usr/bin/env bash

# Copyright 2024 Canonical Ltd.

set -e

_host="localhost:9999"
_base="$_host/rebac/v1"
_names=(alpha bravo charlie delta echo foxtrot golf hotel india juliet kilo lima mike november oscar papa quebec romeo sierra tango uniform victor whiskey x yankee zulu)

function onexit {
    curl "$_host/shutdown"
    wait $_PID1
    echo "Server shut down"
}
trap onexit EXIT

## Run server in background
go build -o server ./cmd
./server &
_PID1=$!
echo $_PID1
sleep 1 # Make sure server is up and running

## Reset state

curl --fail-with-body -w "\n" -X GET "$_host/reset"

## Add single entities

for n in ${_names[@]}
do
    echo -n 'POST group: '
    curl --fail-with-body -w "\n" -X POST "$_base/groups" -d "{\"name\":\"group-$n\"}"
    echo -n '  > GET group: '
    curl --fail-with-body -w "\n" -X GET "$_base/groups/group-$n"
    echo -n '  > PUT group: '
    curl --fail-with-body -w "\n" -X PUT "$_base/groups/group-$n" -d "{\"id\":\"group-$n\",\"name\":\"group-$n--updated\"}"
    echo -n '  > GET updated group: '
    curl --fail-with-body -w "\n" -X GET "$_base/groups/group-$n"
done

for n in ${_names[@]}
do
    echo -n 'POST identity: '
    curl --fail-with-body -w "\n" -X POST "$_base/identities" -d "{\"addedBy\":\"owner-of-$n\",\"email\":\"$n@host.com\",\"source\":\"some-idp\"}"
    echo -n '  > GET identity: '
    curl --fail-with-body -w "\n" -X GET "$_base/identities/$n@host.com"
    echo -n '  > PUT identity: '
    curl --fail-with-body -w "\n" -X PUT "$_base/identities/$n@host.com" -d "{\"id\":\"$n@host.com\",\"addedBy\":\"owner-of-$n--updated\",\"email\":\"$n@host.com\",\"source\":\"some-idp--updated\"}"
    echo -n '  > GET updated identity: '
    curl --fail-with-body -w "\n" -X GET "$_base/identities/$n@host.com"
done

for n in ${_names[@]}
do
    echo -n 'POST role: '
    curl --fail-with-body -w "\n" -X POST "$_base/roles" -d "{\"name\":\"role-$n\"}"
    echo -n '  > GET role: '
    curl --fail-with-body -w "\n" -X GET "$_base/roles/role-$n"
    echo -n '  > PUT role: '
    curl --fail-with-body -w "\n" -X PUT "$_base/roles/role-$n" -d "{\"id\":\"role-$n\",\"name\":\"role-$n--updated\"}"
    echo -n '  > GET updated role: '
    curl --fail-with-body -w "\n" -X GET "$_base/roles/role-$n"
done


for n in ${_names[@]}
do
    echo -n 'POST idp: '
    curl --fail-with-body -w "\n" -X POST "$_base/authentication" -d "{\"name\":\"idp-$n\"}"
    echo -n '  > GET idp: '
    curl --fail-with-body -w "\n" -X GET "$_base/authentication/idp-$n"
    echo -n '  > PUT idp: '
    curl --fail-with-body -w "\n" -X PUT "$_base/authentication/idp-$n" -d "{\"id\":\"idp-$n\",\"name\":\"idp-$n--updated\"}"
    echo -n '  > GET updated idp: '
    curl --fail-with-body -w "\n" -X GET "$_base/authentication/idp-$n"
done

## Add relationships

for n in ${_names[@]}
do
    echo -n "PATCH group identities: group-$n"
    curl --fail-with-body -w "\n" -X PATCH "$_base/groups/group-$n/identities" -d "{\"patches\":[{\"op\":\"add\",\"identity\":\"$n@host.com\"}]}"
    echo -n '  > GET group identities: '
    curl --fail-with-body -w "\n" -X GET "$_base/groups/group-$n/identities"
done

for n in ${_names[@]}
do
    echo -n "PATCH group roles: group-$n"
    curl --fail-with-body -w "\n" -X PATCH "$_base/groups/group-$n/roles" -d "{\"patches\":[{\"op\":\"add\",\"role\":\"role-$n\"}]}"
    echo -n '  > GET group roles: '
    curl --fail-with-body -w "\n" -X GET "$_base/groups/group-$n/roles"
done

for n in ${_names[@]}
do
    echo -n "PATCH group entitlements: group-$n"
    curl --fail-with-body -w "\n" -X PATCH "$_base/groups/group-$n/entitlements" -d "{\"patches\":[{\"op\":\"add\",\"entitlement\":{\"entitlement_type\":\"entitlement-type-$n\",\"entity_name\":\"entity-$n\",\"entity_type\":\"entity-type-$n\"}}]}"
    echo -n '  > GET group entitlements: '
    curl --fail-with-body -w "\n" -X GET "$_base/groups/group-$n/entitlements"
done

for n in ${_names[@]}
do
    echo -n "PATCH identity groups: $n@host.com"
    curl --fail-with-body -w "\n" -X PATCH "$_base/identities/$n@host.com/groups" -d "{\"patches\":[{\"op\":\"add\",\"group\":\"group-$n\"}]}"
    echo -n '  > GET identity groups: '
    curl --fail-with-body -w "\n" -X GET "$_base/identities/$n@host.com/groups"
done

for n in ${_names[@]}
do
    echo -n "PATCH identity roles: $n@host.com"
    curl --fail-with-body -w "\n" -X PATCH "$_base/identities/$n@host.com/roles" -d "{\"patches\":[{\"op\":\"add\",\"role\":\"role-$n\"}]}"
    echo -n '  > GET identity roles: '
    curl --fail-with-body -w "\n" -X GET "$_base/identities/$n@host.com/roles"
done

for n in ${_names[@]}
do
    echo -n "PATCH identity entitlements: $n@host.com"
    curl --fail-with-body -w "\n" -X PATCH "$_base/identities/$n@host.com/entitlements" -d "{\"patches\":[{\"op\":\"add\",\"entitlement\":{\"entitlement_type\":\"entitlement-type-$n\",\"entity_name\":\"entity-$n\",\"entity_type\":\"entity-type-$n\"}}]}"
    echo -n '  > GET identity entitlements: '
    curl --fail-with-body -w "\n" -X GET "$_base/identities/$n@host.com/entitlements"
done

for n in ${_names[@]}
do
    echo -n "PATCH role entitlements: role-$n"
    curl --fail-with-body -w "\n" -X PATCH "$_base/roles/role-$n/entitlements" -d "{\"patches\":[{\"op\":\"add\",\"entitlement\":{\"entitlement_type\":\"entitlement-type-$n\",\"entity_name\":\"entity-$n\",\"entity_type\":\"entity-type-$n\"}}]}"
    echo -n '  > GET role entitlements: '
    curl --fail-with-body -w "\n" -X GET "$_base/roles/role-$n/entitlements"
done

## Invoke other endpoints

echo -n 'GET resources: '
curl --fail-with-body -w "\n" -X GET "$_base/resources"
echo -n 'GET entitlements: '
curl --fail-with-body -w "\n" -X GET "$_base/entitlements"
echo -n 'GET entitlements/raw: '
curl --fail-with-body -w "\n" -X GET "$_base/entitlements/raw"
echo -n 'GET swagger.json (first few chars): '
curl --fail-with-body -X GET "$_base/swagger.json" 2>/dev/null | cut -c1-40
echo -n 'GET capabilities'
curl --fail-with-body -X GET "$_base/capabilities"

## Here are some more endpoints that you can try

# curl -X GET "$_base/swagger.json"
# curl -X GET "$_base/groups?page=1&size=10"
# curl -X GET "$_base/identities?page=1&size=10"

if [[ "$1" != "--cleanup" ]]; then
    exit 0
fi

## Clean up by calling DELETE/PATCH endpoints.

for n in ${_names[@]}
do
    echo -n "PATCH group identities: group-$n"
    curl --fail-with-body -w "\n" -X PATCH "$_base/groups/group-$n/identities" -d "{\"patches\":[{\"op\":\"remove\",\"identity\":\"$n@host.com\"}]}"
    echo -n '  > GET group identities: '
    curl --fail-with-body -w "\n" -X GET "$_base/groups/group-$n/identities"
done

for n in ${_names[@]}
do
    echo -n "PATCH group roles: group-$n"
    curl --fail-with-body -w "\n" -X PATCH "$_base/groups/group-$n/roles" -d "{\"patches\":[{\"op\":\"remove\",\"role\":\"role-$n\"}]}"
    echo -n '  > GET group roles: '
    curl --fail-with-body -w "\n" -X GET "$_base/groups/group-$n/roles"
done

for n in ${_names[@]}
do
    echo -n "PATCH group entitlements: group-$n"
    curl --fail-with-body -w "\n" -X PATCH "$_base/groups/group-$n/entitlements" -d "{\"patches\":[{\"op\":\"remove\",\"entitlement\":{\"entitlement_type\":\"entitlement-type-$n\",\"entity_name\":\"entity-$n\",\"entity_type\":\"entity-type-$n\"}}]}"
    echo -n '  > GET group entitlements: '
    curl --fail-with-body -w "\n" -X GET "$_base/groups/group-$n/entitlements"
done

for n in ${_names[@]}
do
    echo -n "PATCH identity groups: $n@host.com"
    curl --fail-with-body -w "\n" -X PATCH "$_base/identities/$n@host.com/groups" -d "{\"patches\":[{\"op\":\"remove\",\"group\":\"group-$n\"}]}"
    echo -n '  > GET identity groups: '
    curl --fail-with-body -w "\n" -X GET "$_base/identities/$n@host.com/groups"
done

for n in ${_names[@]}
do
    echo -n "PATCH identity roles: $n@host.com"
    curl --fail-with-body -w "\n" -X PATCH "$_base/identities/$n@host.com/roles" -d "{\"patches\":[{\"op\":\"remove\",\"role\":\"role-$n\"}]}"
    echo -n '  > GET identity roles: '
    curl --fail-with-body -w "\n" -X GET "$_base/identities/$n@host.com/roles"
done

for n in ${_names[@]}
do
    echo -n "PATCH identity entitlements: $n@host.com"
    curl --fail-with-body -w "\n" -X PATCH "$_base/identities/$n@host.com/entitlements" -d "{\"patches\":[{\"op\":\"remove\",\"entitlement\":{\"entitlement_type\":\"entitlement-type-$n\",\"entity_name\":\"entity-$n\",\"entity_type\":\"entity-type-$n\"}}]}"
    echo -n '  > GET identity entitlements: '
    curl --fail-with-body -w "\n" -X GET "$_base/identities/$n@host.com/entitlements"
done

for n in ${_names[@]}
do
    echo -n "PATCH role entitlements: role-$n"
    curl --fail-with-body -w "\n" -X PATCH "$_base/roles/role-$n/entitlements" -d "{\"patches\":[{\"op\":\"remove\",\"entitlement\":{\"entitlement_type\":\"entitlement-type-$n\",\"entity_name\":\"entity-$n\",\"entity_type\":\"entity-type-$n\"}}]}"
    echo -n '  > GET role entitlements: '
    curl --fail-with-body -w "\n" -X GET "$_base/roles/role-$n/entitlements"
done

for n in ${_names[@]}
do
    echo -n "DELETE group: group-$n"
    curl --fail-with-body -w "\n" -X DELETE "$_base/groups/group-$n"
done 

for n in ${_names[@]}
do
    echo -n "DELETE identity: $n@host.com"
    curl --fail-with-body -w "\n" -X DELETE "$_base/identities/$n@host.com"
done 

for n in ${_names[@]}
do
    echo -n "DELETE role: role-$n"
    curl --fail-with-body -w "\n" -X DELETE "$_base/roles/role-$n"
done 

for n in ${_names[@]}
do
    echo -n "DELETE idp: idp-$n"
    curl --fail-with-body -w "\n" -X DELETE "$_base/authentication/idp-$n"
done 

## At the end, `state.json` and `state.zero.json` must be equal.
diff --ignore-trailing-space state.zero.json state.json
