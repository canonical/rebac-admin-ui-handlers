#!/usr/bin/env bash

# Copyright 2024 Canonical Ltd.

set -e

_usage='Exercise running server

Options:
  --help           displays current doc.
  --reset          resets the server state at the beginning; default is false.
  --cleanup        cleans up created entities/relationships at the end; default is false.
  --bail-on-error  exits when a request fails (status code >= 400); default is false.
'

_host="localhost:9999"
_base="$_host/rebac/v1"
_names=(alpha bravo charlie delta echo foxtrot golf hotel india juliet kilo lima mike november oscar papa quebec romeo sierra tango uniform victor whiskey x yankee zulu)

_reset_at_start=""
_cleanup_at_end=""
_bail_on_error=""
_help=""
for option in "$@" ; do
    if [[ "$option" == "--reset" ]] ; then
        _reset_at_start="true"
    fi
    if [[ "$option" == "--cleanup" ]] ; then
        _cleanup_at_end="true"
    fi
    if [[ "$option" == "--bail-on-error" ]] ; then
        _bail_on_error="true"
    fi
    if [[ "$option" == "--help" || "$option" == "-h" ]] ; then
        _help="true"
        break
    fi
done

if [ "$_help" == "true" ]; then
    echo "$_usage"
    exit 0
fi

## Check if the server is running
if ! curl -s "$_host/health"; then
    function onexit {
        curl "$_host/shutdown"
        wait $_PID1
        echo "Server shut down"
    }
    trap onexit EXIT

    ## Run server in background
    go build -o server ./cmd
    bash -c 'sleep 1 && ./server' &
    _PID1=$!

    echo waiting for the server to be ready
    while ! curl -s "$_host/ready"
    do
        sleep 0.1
    done
fi

_opts='-w "\n"'
if [ "$_bail_on_error" == "true" ]; then
    _opts="--fail-with-body $_opts"
fi

## Reset state

if [ "$_reset_at_start" == "true" ]; then
    echo -n 'reset state'
    curl $_opts -X GET "$_host/reset"
fi

## Add single entities

for n in ${_names[@]}
do
    echo -n 'POST group: '
    curl $_opts -X POST "$_base/groups" -d "{\"name\":\"group-$n\"}"
    echo -n '  > GET group: '
    curl $_opts -X GET "$_base/groups/group-$n"
    echo -n '  > PUT group: '
    curl $_opts -X PUT "$_base/groups/group-$n" -d "{\"id\":\"group-$n\",\"name\":\"group-$n--updated\"}"
    echo -n '  > GET updated group: '
    curl $_opts -X GET "$_base/groups/group-$n"
done

for n in ${_names[@]}
do
    echo -n 'POST identity: '
    curl $_opts -X POST "$_base/identities" -d "{\"addedBy\":\"owner-of-$n\",\"email\":\"$n@host.com\",\"source\":\"some-idp\"}"
    echo -n '  > GET identity: '
    curl $_opts -X GET "$_base/identities/$n@host.com"
    echo -n '  > PUT identity: '
    curl $_opts -X PUT "$_base/identities/$n@host.com" -d "{\"id\":\"$n@host.com\",\"addedBy\":\"owner-of-$n--updated\",\"email\":\"$n@host.com\",\"source\":\"some-idp--updated\"}"
    echo -n '  > GET updated identity: '
    curl $_opts -X GET "$_base/identities/$n@host.com"
done

for n in ${_names[@]}
do
    echo -n 'POST role: '
    curl $_opts -X POST "$_base/roles" -d "{\"name\":\"role-$n\"}"
    echo -n '  > GET role: '
    curl $_opts -X GET "$_base/roles/role-$n"
    echo -n '  > PUT role: '
    curl $_opts -X PUT "$_base/roles/role-$n" -d "{\"id\":\"role-$n\",\"name\":\"role-$n--updated\"}"
    echo -n '  > GET updated role: '
    curl $_opts -X GET "$_base/roles/role-$n"
done


for n in ${_names[@]}
do
    echo -n 'POST idp: '
    curl $_opts -X POST "$_base/authentication" -d "{\"name\":\"idp-$n\"}"
    echo -n '  > GET idp: '
    curl $_opts -X GET "$_base/authentication/idp-$n"
    echo -n '  > PUT idp: '
    curl $_opts -X PUT "$_base/authentication/idp-$n" -d "{\"id\":\"idp-$n\",\"name\":\"idp-$n--updated\"}"
    echo -n '  > GET updated idp: '
    curl $_opts -X GET "$_base/authentication/idp-$n"
done

## Add relationships

for n in ${_names[@]}
do
    echo -n "PATCH group identities: group-$n"
    curl $_opts -X PATCH "$_base/groups/group-$n/identities" -d "{\"patches\":[{\"op\":\"add\",\"identity\":\"$n@host.com\"}]}"
    echo -n '  > GET group identities: '
    curl $_opts -X GET "$_base/groups/group-$n/identities"
done

for n in ${_names[@]}
do
    echo -n "PATCH group roles: group-$n"
    curl $_opts -X PATCH "$_base/groups/group-$n/roles" -d "{\"patches\":[{\"op\":\"add\",\"role\":\"role-$n\"}]}"
    echo -n '  > GET group roles: '
    curl $_opts -X GET "$_base/groups/group-$n/roles"
done

for n in ${_names[@]}
do
    echo -n "PATCH group entitlements: group-$n"
    curl $_opts -X PATCH "$_base/groups/group-$n/entitlements" -d "{\"patches\":[{\"op\":\"add\",\"entitlement\":{\"entitlement_type\":\"entitlement-type-$n\",\"entity_name\":\"entity-$n\",\"entity_type\":\"entity-type-$n\"}}]}"
    echo -n '  > GET group entitlements: '
    curl $_opts -X GET "$_base/groups/group-$n/entitlements"
done

for n in ${_names[@]}
do
    echo -n "PATCH identity groups: $n@host.com"
    curl $_opts -X PATCH "$_base/identities/$n@host.com/groups" -d "{\"patches\":[{\"op\":\"add\",\"group\":\"group-$n\"}]}"
    echo -n '  > GET identity groups: '
    curl $_opts -X GET "$_base/identities/$n@host.com/groups"
done

for n in ${_names[@]}
do
    echo -n "PATCH identity roles: $n@host.com"
    curl $_opts -X PATCH "$_base/identities/$n@host.com/roles" -d "{\"patches\":[{\"op\":\"add\",\"role\":\"role-$n\"}]}"
    echo -n '  > GET identity roles: '
    curl $_opts -X GET "$_base/identities/$n@host.com/roles"
done

for n in ${_names[@]}
do
    echo -n "PATCH identity entitlements: $n@host.com"
    curl $_opts -X PATCH "$_base/identities/$n@host.com/entitlements" -d "{\"patches\":[{\"op\":\"add\",\"entitlement\":{\"entitlement_type\":\"entitlement-type-$n\",\"entity_name\":\"entity-$n\",\"entity_type\":\"entity-type-$n\"}}]}"
    echo -n '  > GET identity entitlements: '
    curl $_opts -X GET "$_base/identities/$n@host.com/entitlements"
done

for n in ${_names[@]}
do
    echo -n "PATCH role entitlements: role-$n"
    curl $_opts -X PATCH "$_base/roles/role-$n/entitlements" -d "{\"patches\":[{\"op\":\"add\",\"entitlement\":{\"entitlement_type\":\"entitlement-type-$n\",\"entity_name\":\"entity-$n\",\"entity_type\":\"entity-type-$n\"}}]}"
    echo -n '  > GET role entitlements: '
    curl $_opts -X GET "$_base/roles/role-$n/entitlements"
done

## Invoke other endpoints

echo -n 'GET resources: '
curl $_opts -X GET "$_base/resources"
echo -n 'GET entitlements: '
curl $_opts -X GET "$_base/entitlements"
echo -n 'GET entitlements/raw: '
curl $_opts -X GET "$_base/entitlements/raw"
echo -n 'GET swagger.json (first few chars): '
curl $_opts -X GET "$_base/swagger.json" 2>/dev/null | cut -c1-40
echo -n 'GET capabilities'
curl $_opts -X GET "$_base/capabilities"

## Here are some more endpoints that you can try

# curl -X GET "$_base/swagger.json"
# curl -X GET "$_base/groups?page=1&size=10"
# curl -X GET "$_base/identities?page=1&size=10"

if [[ "$_cleanup_at_end" != "true" ]]; then
    exit 0
fi

## Clean up by calling DELETE/PATCH endpoints.

for n in ${_names[@]}
do
    echo -n "PATCH group identities: group-$n"
    curl $_opts -X PATCH "$_base/groups/group-$n/identities" -d "{\"patches\":[{\"op\":\"remove\",\"identity\":\"$n@host.com\"}]}"
    echo -n '  > GET group identities: '
    curl $_opts -X GET "$_base/groups/group-$n/identities"
done

for n in ${_names[@]}
do
    echo -n "PATCH group roles: group-$n"
    curl $_opts -X PATCH "$_base/groups/group-$n/roles" -d "{\"patches\":[{\"op\":\"remove\",\"role\":\"role-$n\"}]}"
    echo -n '  > GET group roles: '
    curl $_opts -X GET "$_base/groups/group-$n/roles"
done

for n in ${_names[@]}
do
    echo -n "PATCH group entitlements: group-$n"
    curl $_opts -X PATCH "$_base/groups/group-$n/entitlements" -d "{\"patches\":[{\"op\":\"remove\",\"entitlement\":{\"entitlement_type\":\"entitlement-type-$n\",\"entity_name\":\"entity-$n\",\"entity_type\":\"entity-type-$n\"}}]}"
    echo -n '  > GET group entitlements: '
    curl $_opts -X GET "$_base/groups/group-$n/entitlements"
done

for n in ${_names[@]}
do
    echo -n "PATCH identity groups: $n@host.com"
    curl $_opts -X PATCH "$_base/identities/$n@host.com/groups" -d "{\"patches\":[{\"op\":\"remove\",\"group\":\"group-$n\"}]}"
    echo -n '  > GET identity groups: '
    curl $_opts -X GET "$_base/identities/$n@host.com/groups"
done

for n in ${_names[@]}
do
    echo -n "PATCH identity roles: $n@host.com"
    curl $_opts -X PATCH "$_base/identities/$n@host.com/roles" -d "{\"patches\":[{\"op\":\"remove\",\"role\":\"role-$n\"}]}"
    echo -n '  > GET identity roles: '
    curl $_opts -X GET "$_base/identities/$n@host.com/roles"
done

for n in ${_names[@]}
do
    echo -n "PATCH identity entitlements: $n@host.com"
    curl $_opts -X PATCH "$_base/identities/$n@host.com/entitlements" -d "{\"patches\":[{\"op\":\"remove\",\"entitlement\":{\"entitlement_type\":\"entitlement-type-$n\",\"entity_name\":\"entity-$n\",\"entity_type\":\"entity-type-$n\"}}]}"
    echo -n '  > GET identity entitlements: '
    curl $_opts -X GET "$_base/identities/$n@host.com/entitlements"
done

for n in ${_names[@]}
do
    echo -n "PATCH role entitlements: role-$n"
    curl $_opts -X PATCH "$_base/roles/role-$n/entitlements" -d "{\"patches\":[{\"op\":\"remove\",\"entitlement\":{\"entitlement_type\":\"entitlement-type-$n\",\"entity_name\":\"entity-$n\",\"entity_type\":\"entity-type-$n\"}}]}"
    echo -n '  > GET role entitlements: '
    curl $_opts -X GET "$_base/roles/role-$n/entitlements"
done

for n in ${_names[@]}
do
    echo -n "DELETE group: group-$n"
    curl $_opts -X DELETE "$_base/groups/group-$n"
done

for n in ${_names[@]}
do
    echo -n "DELETE identity: $n@host.com"
    curl $_opts -X DELETE "$_base/identities/$n@host.com"
done

for n in ${_names[@]}
do
    echo -n "DELETE role: role-$n"
    curl $_opts -X DELETE "$_base/roles/role-$n"
done

for n in ${_names[@]}
do
    echo -n "DELETE idp: idp-$n"
    curl $_opts -X DELETE "$_base/authentication/idp-$n"
done
