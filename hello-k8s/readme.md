## K3s
1. k3d 全名是 k3s in docker，k3s 是輕量級的 Kubernetes，由 Rancher Labs 發佈及維護。k3s 運行所需的資源較少，與 kubernetes 相比刪除了部分功能 (5個項目)，因此才稱為 k3s。
2. 其中有server、agent的差異


| 角色  | 功能 | 運行的組件 | 是否能管理 Cluster？ |
|-----|----|-------|----------------|
| Server (主節點) | 控制整個 Kubernetes Cluster，負責 API Server、Scheduler、Controller Manager | etcd、kube-apiserver、kube-controller-manager、kube-scheduler |  123 |
| Agent (工作節點) | 執行 Pod 和應用程式，負責計算資源 | kubelet、kube-proxy、container runtime (Docker, containerd) | ❌ 否 (不管理 Cluster，只執行 Workload) |

## 在Window中該怎麼安裝k3s
1. curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
2. 建立一個 K3s 叢集 `k3d cluster create mycluster`
3. 檢查 K3s 狀態 `kubectl get nodes`
4. 撰寫好對應的yaml後，啟動配置 `kubectl apply -f nginx-deployment.yaml`
5. 確認是否正確運行(svc = service)
    ```markdown
    kubectl get pods
    kubectl get svc
    ```
6. 確保 kubectl 連接到正確的 K3d 叢集 `kubectl config current-context`
7. 如果顯示的不是 k3d-mycluster（或你的叢集名稱），可以手動設定：
    ```markdown
    kubectl config use-context k3d-mycluster
    ```
8. 綁定外部IP 8080 => 內部的30080，提供localhost能使用8080來訪問服務
    ```markdown
    k3d cluster edit mycluster --port-add "8080:30080@server:0"
    ```
9. 可以用 `k3s cluster list` 、`docker ps` 來查看是否有綁定正確8080:30080

### 其他語法
```markdown
kubectl get pods
kubectl get pods -o wide

kubectl get node
kubectl get node -o wide

# svc就是service
kubectl get svc
kubectl get svc -o wide

# rs就是replicaset
kubectl get replicaset
kubectl get replicaset -o wide

# ns就是namespace
kubectl get ns


kubectl get deployment

kubectl get all

k3d cluster list

netstat -an | grep 8080
```

### 參考
https://www.youtube.com/watch?v=SL83f7Nzxr0&t=1272s


