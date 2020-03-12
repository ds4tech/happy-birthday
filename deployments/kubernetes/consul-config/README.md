# Run Consul server on Kubernetes Engine cluster

## Steps

1. 
```
helm install -f helm-consul-values.yaml hashicorp ./consul-helm

k port-forward hashicorp-consul-server-0  8500:8500
```
or
```k port-forward dashboard 9002:9002```

Manage consul rules:
```
consul intention create dashboard counting
consul intention create -deny dashboard counting
consul intention delete dashboard counting
```

**NOTE:** Refer to the [Consul on Kubernetes](https://learn.hashicorp.com/consul/kubernetes/minikube) documentation for more details.
