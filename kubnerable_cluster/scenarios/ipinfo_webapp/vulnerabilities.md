# IP Info Webapp
In this scenario, a web application to display your IP has high Linux capabilities. Specially the `CAP_NET_RAW`, which can
let an attacker perform ARP or DNS spoofing through the workload.

## Vulnerabilities
* Containers with CAP_NET_RAW capability

### MITRE Attack Techniques
* [Exploitation of Remote Services](https://attack.mitre.org/techniques/T1210/)
* [Network Sniffing](https://attack.mitre.org/techniques/T1040/)

## Local connection
First, check that the created service is running properly and has a cluster IP assigned:
```bash
kubectl get services
```
Once identified, forward the workload port to localhost:
```bash
kubectl port-forward service/ipinfo-webapp-service 8080:80
```
Web access should be enabled through `localhost:8080`.
