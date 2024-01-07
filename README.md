# K8s regional scheduler plugin
 This Kubernetes scheduler plugin extends the filter interface to geographically align pods with specific regions, ensuring efficient node assignment within clusters.

 ## Install

### Build an image

I’ve already built an image `casapuioana/k8s:region-filter-plugin`. 
If you want to build your own, run the following command:

```
docker build -t casapuioana/k8s:region-filter-plugin ./
```
### Create a Kubernetes cluster

If you do not have a cluster yet, create one by using tools like kind or minikube.

Here is a configuration file for a kind cluster with 3 nodes and labels

```
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
name: k8s
nodes:
  - role: control-plane
    image: kindest/node:v1.23.10
  - role: worker
    image: kindest/node:v1.23.10
    labels:
      custom-label: us-west
  - role: worker
    image: kindest/node:v1.23.10
    labels:
      custom-label: us-south

```

#### Log into the control plane

```
docker exec -it $(sudo docker ps | grep control-plane | awk '{print $1}') bash
```

#### Backup `kube-scheduler.yaml`
```
cp /etc/kubernetes/manifests/kube-scheduler.yaml /etc/kubernetes/kube-scheduler.yaml
```
#### Modify the content of `/etc/kubernetes/manifests/kube-scheduler.yaml` according to `region-scheduler/config/kube-scheduler.yaml`

#### Create `/etc/kubernetes/scheduler-config.yaml` (see `region-scheduler/config/scheduler-config.yaml`)

## Test

### You can use the test pods from the `test` folder

```
$ kubectl apply -f test/pod-south.yaml
``` 

In order to test the scheduler status run this command

```
kubectl logs -n kube-system -l component=kube-scheduler
```

And in order to check the result run this

```
$ kubectl get pods -o wide --show-labels
NAME         READY   STATUS    RESTARTS   AGE    LABELS
nginx        1/1     Running   0          2d7h   <none>
test-pod-1   1/1     Running   0          30h    custom-label=us-south
test-pod-2   0/1     Pending   0          30h    custom-label=us-north
```

And if we want to further check the status of the `test-pod-2`

```
$ kubectl describe pod test-pod-2
# content
Events:
  Type     Reason            Age                   From               Message
  ----     ------            ----                  ----               -------
  Warning  FailedScheduling  10m (x1547 over 30h)  default-scheduler  0/3 nodes are available: 1 node(s) had taint {node-role.kubernetes.io/master: }, that the pod didn't tolerate, 2 Node is in a different region.
```
We notice that the pod with the us-north label is in a pending state, thus it has not been scheduled.
