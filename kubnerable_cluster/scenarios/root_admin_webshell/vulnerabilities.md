# Cluster Administrator Shell
In this scenario, a cluster administrator has left a highly privileged workload running on the cluster.

## Vulnerabilities
* Workload running as root
* Privileged workload
* Privilege escalation inside the container
* Workload shares the host PID
* Workload share the host IPC
* Workload mounts sensitive folders from the hosts (host filesystem)

### MITRE Attack Techniques
* [Abuse Elevation Control Mechanism: Sudo and Sudo Caching](https://attack.mitre.org/techniques/T1548/003/)
* [Abuse Elevation Control Mechanism: Bypass User Account Control](https://attack.mitre.org/techniques/T1548/002/)
* [File and Directory Discovery](https://attack.mitre.org/techniques/T1083/)
* [Hijack Execution Flow: Services File Permissions Weakness](https://attack.mitre.org/techniques/T1574/010/)
* [Indicator Removal on Host: File Deletion](https://attack.mitre.org/techniques/T1070/004/)

## Local connection
First, check that the created service is running properly and has a cluster IP assigned:
```bash
kubectl get services
```
Once identified, forward the workload port to localhost:
```bash
kubectl port-forward service/root-admin-webshell-service 4200:4200
```
Web access should be enabled through `localhost:4200`.
