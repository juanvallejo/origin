#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

OS_ROOT=$(dirname "${BASH_SOURCE}")/../..
source "${OS_ROOT}/hack/lib/init.sh"
os::log::stacktrace::install
trap os::test::junit::reconcile_output EXIT

# Cleanup cluster resources created by this test
(
  set +e
  oc delete all,templates,secrets --all
  exit 0
) &>/dev/null

# check to make sure that "get"ting a resource with no config file present
# still returns error indicating that no config-file is set
os::test::junit::declare_suite_start "cmd/configuration"
os::cmd::expect_success_and_not_text 'oc get buildconfig' 'does not exist'
(
  export HOME=/tmp
  unset KUBECONFIG
  unset KUBERNETES_MASTER

  os::cmd::expect_failure_and_text 'env -u KUBERNETES_SERVICE_HOST oc get buildconfig --user=""' 'Missing or incomplete configuration info'
  os::cmd::expect_failure_and_text 'env -u KUBERNETES_SERVICE_HOST oc get buildconfig --context=""' 'Missing or incomplete configuration info'
  os::cmd::expect_failure_and_text 'env -u KUBERNETES_SERVICE_HOST oc get buildconfig --cluster=""' 'Missing or incomplete configuration info'

  os::cmd::expect_failure_and_text 'env -u KUBERNETES_SERVICE_HOST oc get buildconfig --context="test"' 'context was not found for specified context\: test'
  os::cmd::expect_failure_and_text 'env -u KUBERNETES_SERVICE_HOST oc get buildconfig --user="test"' 'auth info "test" does not exist'
  os::cmd::expect_failure_and_text 'env -u KUBERNETES_SERVICE_HOST oc get buildconfig --cluster="test"' 'no server found for cluster \"test\"'

  os::cmd::expect_failure_and_text 'oc get buildconfig --config=missing' 'missing: no such file or directory'

  # define temp location for new config
  NEW_CONFIG_LOC="${BASETMPDIR}/new-config.yaml"

  # make sure non-existing --cluster and --user can still be set
  os::cmd::expect_success_and_text "oc config set-context new-context-name --cluster=missing-cluster --user=missing-user --namespace=default --config='${NEW_CONFIG_LOC}'" 'Context "new-context-name" '
  os::cmd::expect_failure_and_text "env -u KUBERNETES_SERVICE_HOST -u KUBECONFIG -u KUBERNETES_MASTER oc get buildconfig --config='${NEW_CONFIG_LOC}'" 'Missing or incomplete configuration info'
  os::cmd::expect_failure_and_text "env -u KUBERNETES_SERVICE_HOST oc get buildconfig --config='${NEW_CONFIG_LOC}'" 'Missing or incomplete configuration info'
)
echo "config error handling: ok"
os::test::junit::declare_suite_end
