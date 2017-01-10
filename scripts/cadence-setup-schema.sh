#!/bin/bash
set -uxeo pipefail
# this script is used to setup the "cadence" keyspace
# and load all the tables in the .cql file within this keyspace,
# if cassandra is running
#
# this script is only intended to be used in test environments

my_dir="$(dirname "$0")"
source "$my_dir/cadence-wait-cassandra"
. $my_dir/cadence-environment

# the default cqlsh listen port is 9042
port=9042
address=${LISTEN_ADDRESS:-127.0.0.1}

# the default keyspace is cherami
# TODO: probably allow getting this from command line
keyspace="cadence"

wait_for_cassandra $port
res=$?
if [ $res -eq 0 ]; then
    $CADENCE_CQLSH_DIR/cqlsh -f $CADENCE_SCHEMA_DIR/keyspace_test.cql $address
    $CADENCE_CQLSH_DIR/cqlsh -k $workflow_keyspace -f $CADENCE_SCHEMA_DIR/workflow_test.cql $address
fi
