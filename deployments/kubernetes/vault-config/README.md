# Run Vault server on Kubernetes Engine cluster

## Steps

1. ```helm install vault ./vault-helm```
2. Expose helm service:
```k port-forward service/vault 8200:8200```
3. ```k exec -it vault-0 -- /bin/sh```
4. ```vault operator init -key-shares 1 -key-threshold 1```
5. ```vault operator unseal LhLHI7860=```
6. ```vault login s.vgukMfp```
7. ```vault secrets enable -path=secret kv-v2```
8. ```vault kv put secret/exampleapp/config username=“helmchart” password="secrets"```
9. ```vault kv get secret/exampleapp/config```
10. ```vault auth enable kubernetes```
11. 
```
vault write auth/kubernetes/config \
    token_reviewer_jwt="$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" \
    kubernetes_host="https://$KUBERNETES_PORT_443_TCP_ADDR:443" \
    kubernetes_ca_cert=@/var/run/secrets/kubernetes.io/serviceaccount/ca.crt
```

10. 
```
vault policy write webapp - <<EOH
path "secret/data/webapp/config" {
  capabilities = ["read"]
}
EOH
```
11. 
```
vault write auth/kubernetes/role/webapp \
    bound_service_account_names=vault \
    bound_service_account_namespaces=default \
    policies=webapp \
    ttl=24h
```

kubectl exec -it vault-0 -- vault operator init -key-shares=1 -key-threshold=1 -format=json > cluster-keys.json

kubectl exec -it vault-0 -- vault operator unseal YOzhtj77eNd56I9LxFjtHDoI=
kubectl exec -it vault-1 -- vault operator unseal YOzhtj77eNd56I9LxFjtHDoI=
kubectl exec -it vault-2 -- vault operator unseal YOzhtj77eNd56I9LxFjtHDoI=

kubectl exec -it vault-0 -- /bin/sh



vault kv put secret/webapp/config username="static-user" password="static-password"

vault auth enable kubernetes

**NOTE:** Refer to the [Vault on Kubernetes](https://learn.hashicorp.com/vault/getting-started-k8s/k8s-intro) documentation for more details.
