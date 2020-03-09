# Run Vault server on Kubernetes Engine cluster

## Steps

1. ```helm install vault ./vault-helm```
2. ```k exec -it vault-0 -- /bin/sh```
3.  k port-forward service/vault 8200:8200
4. ```vault operator init -key-shares 1 -key-threshold 1```
5. ```vault operator unseal LhLHI7860=```
6. ```vault login s.vgukMfp```
7. ```vault secrets enable -path=secret kv-v2```
8. ```vault kv put secret/exampleapp/config username=“helmchart” password="secrets"```
9. ```vault kv get secret/exampleapp/config```
10. ```vault auth enable kubernetes```
11. ```vault write auth/kubernetes/config \
        		token_reviewer_jwt="$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" \
        		kubernetes_host="https://$KUBERNETES_PORT_443_TCP_ADDR:443" \
        		kubernetes_ca_cert=@/var/run/secrets/kubernetes.io/serviceaccount/ca.crt```
10. ```vault policy write exampleapp - <<EOH
		path "secret/data/exampleapp/config" {
		  capabilities = ["read"]
		}
	EOH```
11. ```vault write auth/kubernetes/role/exampleapp bound_service_account_names=vault bound_service_account_namespaces=default policies=exampleapp ttl=24h```

**NOTE:** Refer to the [Kubernetes Engine](https://cloud.google.com/kubernetes-engine/docs/quickstart) documentation for more details.
