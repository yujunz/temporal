#!/usr/bin/env bash
set -euo pipefail
IFS=$' \n\t'

export TMPDIR=${TMPDIR:-`mktemp -d`}
my_dir="$(dirname "$BASH_SOURCE")"
# Note that we prefer the local Cadence cassandra, etc.
SEARCHPATH=$my_dir:$my_dir/..:$my_dir/../config:$my_dir/../cassandra/conf:$my_dir/../cassandra/bin:`pwd`:$PATH:/etc/cassandra:/usr/local/etc/cassandra:/usr/lib/jvm/java-8-openjdk-amd64:/usr/local/bin:/usr/sbin:/usr/local/lib:/usr/share/cadence

ffind() {
	local target=$1
	find ${SEARCHPATH//:/ } -maxdepth 1 -name $target -and -type f -print -quit 2>/dev/null || true
}
readonly -f ffind

CADENCE_SCRIPTS_DIR=$my_dir

if [ ! -z ${UBER_ENVIRONMENT:-} ]; then CADENCE_ENVIRONMENT=$UBER_ENVIRONMENT; fi
if [ -z ${CADENCE_ENVIRONMENT:-} ]; then CADENCE_ENVIRONMENT=jenkins; fi

if [ -z ${LISTEN_ADDRESS:-} ]; then LISTEN_ADDRESS=127.0.0.1; fi

CADENCE_CASSANDRA_DIR=$(ffind cassandra)
if [ -z ${CADENCE_CASSANDRA_DIR} ]; then echo "Couldn't find Cassandra"; else CADENCE_CASSANDRA_DIR=`dirname $CADENCE_CASSANDRA_DIR`; fi

CADENCE_CQLSH_DIR=$(ffind cqlsh)
if [ -z ${CADENCE_CQLSH_DIR} ]; then echo "Couldn't find CQLSH"; else CADENCE_CQLSH_DIR=`dirname $CADENCE_CQLSH_DIR`; fi

if [ -z ${CASSANDRA_CONFIG_DIR:-} ]; then
	CASSANDRA_CONFIG_DIR=$(ffind cassandra.yaml)
	if [ -z ${CASSANDRA_CONFIG_DIR} ]; then echo "Couldn't find Cassandra config"; else CASSANDRA_CONFIG_DIR=`dirname $CASSANDRA_CONFIG_DIR`; fi
fi

if [ -z ${CADENCE_SCHEMA_DIR:-} ]; then
	CADENCE_SCHEMA_DIR=$(ffind workflow_test.cql)
	if [ -z ${CADENCE_SCHEMA_DIR} ]; then echo "Couldn't find Cherami workflow schema directory"; else CADENCE_SCHEMA_DIR=$(dirname $CADENCE_SCHEMA_DIR); fi
fi

#Rewrite the Cassandra config and base.yaml with the listen address if it is a loopback address, but only if we haven't done it already
if [[ ${LISTEN_ADDRESS:-} == 127.* ]] && [ ! -f $TMPDIR/base.yaml ]; then
	cp -fv $CASSANDRA_CONFIG_DIR/*.yaml $TMPDIR
	sed -i.bak "s/localhost/$LISTEN_ADDRESS/;s/127.0.0.1/$LISTEN_ADDRESS/;s/ListenAddress: *\"\"/ListenAddress: \"$LISTEN_ADDRESS\"/;" $TMPDIR/{base,cassandra}.yaml
	CASSANDRA_CONFIG_DIR=$TMPDIR
	# Can't just set CASSANDRA_HOME=$TMPDIR, Cassandra can't find its dependencies

	echo -e "\ndata_file_directories:\n    - $TMPDIR\n" >> $TMPDIR/cassandra.yaml
	for d in commitlog_directory saved_caches_directory; do
		echo "$d: $TMPDIR/$d" >> $TMPDIR/cassandra.yaml
	done
	echo >> $TMPDIR/cassandra.yaml
	tail $TMPDIR/cassandra.yaml

	#sed "s#/var/log/cassandra#$TMPDIR#" $CADENCE_CASSANDRA_DIR/cassandra > $TMPDIR/cassandra
	#chmod u+x $TMPDIR/cassandra
	#CADENCE_CASSANDRA_DIR=$TMPDIR
fi