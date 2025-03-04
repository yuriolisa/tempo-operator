Changes by Version
==================

<!-- next version -->

## 0.8.0

### 💡 Enhancements 💡

- `operator`: Make Tempo-Query forwarding on gateway optional (#628)
- `operator`: Support monolithic deployment mode (#710)

  The operator exposes a new CRD `TempoMonolithic`, which manages a Tempo instance in monolithic mode.
  The monolithic mode supports the following additional storage backends: in-memory and file system (persistent volume).
  

### 🧰 Bug fixes 🧰

- `operator`: Fix the cluster-monitoring-view RBAC when operator is deployed in arbitrary namespace (#741)
- `operator`: NIL pointer dereference when OIDC not specified for tenants in static mode (#647)

### Components
- Tempo: [v2.3.1](https://github.com/grafana/tempo/releases/tag/v2.3.1)

## 0.7.0

### 💡 Enhancements 💡

- `operator`: Divide assigned limits with replicas (#721)
- `operator`: Allow override arbitrary tempo configurations (#629)
- `operator`: Create Grafana Tempo Operator datasource (#423)
- `operator`: Add .spec.hashRing.memberlist.enableIPv6 option to enable IPv6 support (#704)
- `operator`: Propagating proxy env vars to component containers (#700)
- `operator`: Upgrade tempo to v2.3.1 (#729)

### 🧰 Bug fixes 🧰

- `operator`: Configure the number of replicas for compactor, querier and query-frontend according to the CR (#712)

### Components
- Tempo: [v2.3.1](https://github.com/grafana/tempo/releases/tag/v2.3.1)

## 0.6.0

### 🛑 Breaking changes 🛑

- `operator`: Move default images from operator configuration to environment variable (#591)
- `operator`: Unset (default) images in TempoStack CR (#674)
  This upgrade reverts any change to the `spec.images` fields of any TempoStack instance.
  Beginning with version 0.6.0, the image location is not stored in the TempoStack instance unless it is changed manually.
  

### 💡 Enhancements 💡

- `operator`: Support configuration of TLS in receiver settings (#527)
- `operator`: Exposing the Tempo API through the gateway (#672)
- `operator`: Reduce log level of certrotation messages (#623)
- `operator`: Upgrade tempo to v2.3.0 (#688)

### 🧰 Bug fixes 🧰

- `gateway`: fix CVE-2023-45142 tempo-gateway-container: opentelemetry: DoS vulnerability in otelhttp (#691)

### Components
- Tempo: [v2.3.0](https://github.com/grafana/tempo/releases/tag/v2.3.0)

## 0.5.0

### 🛑 Breaking changes 🛑

- `operator`: Install operator in tempo-operator-system namespace by default when installed with OLM or manifests of the OpenShift variant (#538)

### 💡 Enhancements 💡

- `operator`: Bump tempo version to 2.2.3 (#646)
- `operands`: Bump operands to fix CVE-2023-39325 (#650)
- `operator`: Expose the OTLP HTTP port in the distributor service. (#610)
- `operator`: Add pprof flag to optionally expose pprof data (#242)
- `operator`: Use tempo service account to query metrics from OpenShift monitoring stack. (#526)
  On OpenShift tempo service account is used to query metrics from OpenShift monitoring stack for the monitor tab.
- `operator`: Support setting a custom CA certificate for S3 object storage (#545)
- `operator`: Enable ingress (or route) in samples, add MinLength validation to .spec.storage.secret.name of the TempoStack CR (#541)
- `operator`: Support monitor tab in Jaeger console (#470)
- `operator`: Explicitly specify log level for all components. (#550)
- `operator`: Support Tempo 2.2.0 (#525)

### 🧰 Bug fixes 🧰

- `operator`: Fix ingester StatefulSet reconciliation if ingester is in an unhealthy state (#597)
- `operator`: Enable mTLS for all components except query-frontend. (#561)
  Only enable mTLS for query-frontend when the gateway is enabled.
- `operator`: Fix for Http2 reset vulnerability CVE-2023-39325 (#642)
- `operator`: Upgrade TempoStack instances once they are switched back from Unmanaged to Managed (#478)

### Components
- Tempo: [v2.2.3](https://github.com/grafana/tempo/releases/tag/v2.2.3)

## 0.4.0

### 💡 Enhancements 💡

- `operator`: Remove operator ServiceMonitor and PrometheusRule when operator deployment is removed (#536)

### 🧰 Bug fixes 🧰

- `operator`: Disable mTLS by default, to allow connections from Grafana to the query-frontend component (#552)
- `apis/tempo/v1alpha1`: provide correct mode name via operator-sdk annotation for oc console (#556)

### Components
- Tempo: [v2.1.1](https://github.com/grafana/tempo/releases/tag/v2.1.1)

## 0.3.0

### 🛑 Breaking changes 🛑

- `operator`: Upgrade tempo to v2.1.1 (#408)

  The `maxSearchBytesPerTrace` global and per-tenant limit is deprecated.
  A new `maxSearchDuration` global and per-tenant limit is available.
  Some metrics got renamed or deleted, see the [Tempo v2.1.0 release notes](https://github.com/grafana/tempo/releases/tag/v2.1.0) for details.
  

### 💡 Enhancements 💡

- `operator`: Add .spec.managedState to TempoStack. It allows disabling reconciliation of the TempoStack CR. The validation and defaulting webhooks remain enabled in unmanaged state. (#411)
- `operator`: Enable mTLS by default for internal components (#505)
- `operator`: Expose the Jaeger GRPC query port (#513)
- `operator`: Expose supported protocols on the distributor (#436)

  The following protocols are now exposed: 
    - Jaeger: thrift-http port 14268, thrift-binary port 6832, thrift-compact port 6831, Grpc port 14250
    - Zipkin port 9411
  
- `operator`: Use internal certificate for internal HTTP server of gateway (#480)
- `operator`: Add ability to create and configure route or ingress for the gateway (#265)
- `operator`: The operator is now at Level 4 - Deep Insights (#504)

  The operator optionally exposes metrics and alerts for the operator and the operand.
  
- `operator`: Enable mTLS on receivers when gateway is enabled (#535)
- `operator`: Enable multitenancy without need the gateway (#224)
- `operator`: Add operator alerts and runbook. (#309)
- `operator`: Add new operator configuration options to enable or disable the creation of ServiceMonitor and PrometheusRule for the operator itself (#491)
- `operator`: Probe webhook server in operator health checks (#459)
- `operator`: Rename Degraded condition to ConfigurationError and expose reconcile errors via a new FailedReconciliation status (#400, #422)
- `operator`: Use consistent log format, specify logger names and update log severity levels of reconcile logs (#430)
- `operator`: Implement operator upgrade (#296)
- `operator`: Validate if createServiceMonitors is enabled when enabling createPrometheusRules in the webhook (#510)
- `operator`: Set tempo version in the status field based on the default tempo version of the operator (#400, #422)

### 🧰 Bug fixes 🧰

- `operator`: Fix a panic when an invalid tenant configuration is provided to the operator. If the authentication is provided but the authorization is not, the validator panics (#494)
- `operator`: Fix TLS configuration of ServiceMonitors (#481)
- `operator`: Always set all status condition values in the tempostack_status_condition metric (#452)
  Additionally, deprecate the `status` label of the tempostack_status_condition metric.
  
- `operator`: Update operator container image location in bundle (#443)
- `operator`: Scope PrometheusRule to a specific TempoStack instance (#485)

### Components
- Tempo: [v2.1.1](https://github.com/grafana/tempo/releases/tag/v2.1.1)

## 0.2.0

### 🛑 Breaking changes 🛑

- `operator`: Rename operator deployment to enable upgrading from 0.1.0 (#432)

  If you have installed the operator via Kubernetes manifests, please run `kubectl -n tempo-operator-system delete deployment tempo-operator-controller-manager` to prune the old deployment.
  If you have installed the operator via OLM, no action is required.


### 💡 Enhancements 💡

- `operator`: Add support for Kubernetes 1.26 and 1.27. (#385, #365)
- `operator`: Configure logging (#217)
- `tests`: Add a smoketest for tempo + opentelemetry-collector + multitenancy (OpenShift) (#202)
- `operator`: Add mTLS support to the communication between gateway and internal components. (#240)
- `operator`: Create ServiceMonitors for Tempo components (#298, #333)
- `operator`: Add operator metrics (#308, #334)
- `operator`: Recover the resource.requests field for the operator manager as the OpenShift guidelines recommend (#426)
- `operator`: add tempo gateway to resource pool, when is enable it will take into account the gateway in the resource calculation. (#201)
- `operator`: Sanitize generated manifest names (#223)
- `operator`: Create one TLS cert/key per component/service instead of having different certs for HTTP and GRPC (#383)
- `operator`: Introducing alerts for operands (#307)

### Components
- Tempo: [v2.0.1](https://github.com/grafana/tempo/releases/tag/v2.0.1)

## 0.1.0

### 🚀 New components 🚀

- `operator`: Initial release of tempo operator
  - Supports [Tempo - v2.0.1](https://github.com/grafana/tempo/releases/tag/v2.0.1)
  
