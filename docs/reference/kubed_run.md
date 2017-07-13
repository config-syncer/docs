## kubed run

Run daemon

### Synopsis


Run daemon

```
kubed run [flags]
```

### Options

```
      --address string         The address of the Kubed API Server (default ":8081")
      --analytics              Send analytical events to Google Analytics (default true)
      --clusterconfig string   Path to cluster config file (default "/home/tamal/go/src/github.com/appscode/kubed/hack/config/clusterconfig.yaml")
      --enable-reverse-index   Reverse indexing of pods to service and others (default true)
  -h, --help                   help for run
      --indexer string         Reverse indexing of pods to service and others (default "indexers.bleve")
      --kubeconfig string      Path to kubeconfig file with authorization information (the master location is set by the master flag).
      --master string          The address of the Kubernetes API server (overrides any value in kubeconfig)
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO
* [kubed](kubed.md)	 - Kubed by AppsCode - Kubernetes Daemon


