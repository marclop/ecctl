## ecctl deployment elasticsearch instances

Manages elasticsearch at the instance level

### Synopsis

Manages elasticsearch at the instance level

```
ecctl deployment elasticsearch instances [flags]
```

### Options

```
  -h, --help   help for instances
```

### Options inherited from parent commands

```
      --apikey string      API key to use to authenticate (If empty will look for EC_APIKEY environment variable)
      --config string      Config name, used to have multiple configs in $HOME/.ecctl/<env> (default "config")
      --force              Do not ask for confirmation
      --format string      Formats the output using a Go template
      --host string        Base URL to use (default "https://api.elastic-cloud.com")
      --insecure           Skips all TLS validation
      --message string     A message to set on cluster operation
      --output string      Output format [text|json] (default "text")
      --pass string        Password to use to authenticate (If empty will look for EC_PASS environment variable)
      --pprof              Enables pprofing and saves the profile to pprof-20060102150405
  -q, --quiet              Suppresses the configuration file used for the run, if any
      --region string      Elastic Cloud region
      --timeout duration   Timeout to use on all HTTP calls (default 30s)
      --trace              Enables tracing saves the trace to trace-20060102150405
      --user string        Username to use to authenticate (If empty will look for EC_USER environment variable)
      --verbose            Enable verbose mode
```

### SEE ALSO

* [ecctl deployment elasticsearch](ecctl_deployment_elasticsearch.md)	 - Manages Elasticsearch clusters
* [ecctl deployment elasticsearch instances override-capacity](ecctl_deployment_elasticsearch_instances_override-capacity.md)	 - Sets overrides at the instance level, use the --all flag to target all instances
* [ecctl deployment elasticsearch instances pause](ecctl_deployment_elasticsearch_instances_pause.md)	 - Pauses (stops) specific Elasticsearch instances, use the --all flag to target all instances
* [ecctl deployment elasticsearch instances resume](ecctl_deployment_elasticsearch_instances_resume.md)	 - Resumes (starts) specific Elasticsearch instances, use the --all flag to target all instances
* [ecctl deployment elasticsearch instances start-routing](ecctl_deployment_elasticsearch_instances_start-routing.md)	 - Resumes routing on specific Elasticsearch instances, use the --all flag to target all instances
* [ecctl deployment elasticsearch instances stop-routing](ecctl_deployment_elasticsearch_instances_stop-routing.md)	 - Stops routing on specific Elasticsearch instances, use the --all flag to target all instances

