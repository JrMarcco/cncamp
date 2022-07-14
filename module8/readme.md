# readme

* simple-web-env-config-map.yaml

  ```yaml
  apiVersion: v1
  kind: ConfigMap
  metadata:
    name: simple-web-env
  data:
    VERSION: "DEFAULT_VALUE"
    APP_NAME: "simple-web"
    APP_VERSION: "latest"
  ```

* simple-web-deployment.yaml

  ```yaml
  apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: simple-web-deployment
  spec:
    replicas: 3
    selector:
      matchLabels:
        app: simple-web
    template:
      metadata:
        labels:
          app: simple-web
      spec:
        containers:
          - name: simple-web
            image: jrmarcco/simple-web
            # 将 configMap : simple-web-env 的数据定义为容器环境变量
            args:
              - /bin/sh
              - -c
              - env
            envFrom:
              - configMapRef:
                  name: simple-web-env
            resources:
              limits:
                memory: 256Mi
                cpu: 500m
              requests:
                memory: 128Mi
                cpu: 100m
            readinessProbe:
              failureThreshold: 3
              httpGet:
                path: /healthz
                port: 8080
                scheme: HTTP
              periodSeconds: 5
              timeoutSeconds: 15
  ```

* 部署 simple-web

  ```powershell
  # 创建 configMap 并查看确认
  kubectl apply -f simple-web-env-config-map.yaml

  kubectl get configmap simple-web-env -oyaml
  # ---
  # apiVersion: v1
  # data:
  #   APP_NAME: simple-web
  #   APP_VERSION: latest
  #   VERSION: DEFAULT_VALUE
  # kind: ConfigMap
  # ...

  # 部署 simple-web 并查看确认
  kubectl apply -f simple-web-deployment.yaml
  kubectl get po -owide
  # NAME                                     READY   STATUS    RESTARTS   AGE     IP              NODE      NOMINATED NODE   READINESS GATES
  # ...
  # simple-web-deployment-6c5b7f4d7d-6jtdq   1/1     Running   0          32s    192.168.64.60   jrx-gcp   <none>           <none>
  # simple-web-deployment-6c5b7f4d7d-9jx87   1/1     Running   0          32s    192.168.64.62   jrx-gcp   <none>           <none>
  # simple-web-deployment-6c5b7f4d7d-v2v8q   1/1     Running   0          32s    192.168.64.61   jrx-gcp   <none>           <none>

  kubectl exec -it simple-web-deployment-6c5b7f4d7d-9dj7n -- bash
  env 
  # ...
  # APP_NAME=simple-web
  # APP_VERSION=latest
  # VERSION=DEFAULT_VALUE
  # ...
  # 环境变量设置正确

  # 测试访问，获取系统变量正确
  curl http://192.168.64.61:8080/envVar?APP_NAME
  # simple-web
  ```
* simple-web-service.yaml

  ```yaml
  apiVersion: v1
  kind: Service
  metadata:
    name: simple-web-svc
  spec:
    type: ClusterIP
    ports:
      - port: 8080
        protocol: TCP
        name: http
    selector:
      app: simple-web
  ```

* simple-web-ingress.yaml

  ```yaml
  apiVersion: networking.k8s.io/v1
  kind: Ingress
  metadata:
    name: simple-web-ingress
    annotations:
      kubernetes.io/ingress.class: "nginx"
  spec:
    # 证书
    tls:
      - hosts:
          - jrx.com
        secretName: jrx-tls
    rules:
      - host: jrx.com
        http:
          paths:
            - path: "/envVar"
              pathType: Exact
              backend:
                service:
                  name: simple-web-svc
                  port:
                    number: 8080
  ```
* 部署高可用

  ```powershell
  # 创建 Service 并查看
  kubectl apply -f simple-web-service.yaml
  kubectl get svc -owide
  # NAME             TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE     SELECTOR
  # kubernetes       ClusterIP   10.96.0.1        <none>        443/TCP        6d6h    <none>
  # nginx-basic      NodePort    10.106.57.213    <none>        80:32320/TCP   4d      app=nginx
  # simple-web-svc   ClusterIP   10.109.215.144   <none>        8080/TCP       3m24s   app=simple-web

  # 测试连接
  curl 10.109.215.144:8080/envVar

  # 生成证书
  openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
      -keyout tls.key \
      -out tls.crt \
      -subj "/CN=jrx.com/O=jrx" \
      -addext "subjectAltName = DNS:jrx.com"

  # 创建 secret
  kubectl create secret tls jrx-tls --cert=./tls.crt --key=./tls.key

  # 创建 ingress
  kubectl apply -f simple-web-ingress.yaml

  # 查看 ingress svc 地址
  kubectl get svc -n ingress-nginx
  # NAME                                 TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                      AGE
  # ingress-nginx-controller             NodePort    10.109.51.27     <none>        80:30772/TCP,443:31266/TCP   40m
  # ingress-nginx-controller-admission   ClusterIP   10.100.227.169   <none>        443/TCP                      40mku

  # 测试连接，访问正常
  curl -H "Host: jrx.com" https://10.109.51.27/envVar?APP_NAME -k
  ```

---
