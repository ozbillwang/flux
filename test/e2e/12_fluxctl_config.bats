#!/usr/bin/env bats

function setup() {
  load lib/env
  load lib/install
  load lib/poll

  kubectl create namespace "$FLUX_NAMESPACE"
  install_git_srv
  tmp_config=$(mktemp -d)
  defer rm -rf "$tmp_config"
  cat > "${tmp_config}/flux.yaml" << EOF
flux-config-version: v1
EOF
  install_flux_with_fluxctl '12_flux_config' '' --config-file="${tmp_config}/flux.yaml"
}

@test "'fluxctl install --config-file=...' smoke test" {
  # Test that the resources from https://github.com/fluxcd/flux-get-started are deployed
  poll_until_true 'namespace demo' 'kubectl describe ns/demo'
  poll_until_true 'workload podinfo' 'kubectl -n demo describe deployment/podinfo'
}

function teardown() {
  run_deferred
  # Although the namespace delete below takes care of removing most Flux
  # elements, the global resources will not be removed without this.
  uninstall_flux_with_fluxctl
  # Removing the namespace also takes care of removing Flux and gitsrv.
  kubectl delete namespace "$FLUX_NAMESPACE"
  # Only remove the demo workloads after Flux, so that they cannot be recreated.
  kubectl delete namespace "$DEMO_NAMESPACE"
}
