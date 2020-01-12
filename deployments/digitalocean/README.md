# Scripts to setup Kubernetes on DigitalOcean

## Scripts to run on Master

```
./install-kubernetes.sh
./create-user.sh
```
One this script is finished, copy an output and run the command on the node:
```kubeadm join ...```

## Scripts to run on Node's

```
./install-node.sh
./create-user.sh
```
