* 使用`https`

  ```powershell
  # 创建 namespace 并添加 istio 标签
  kubectl create ns simple-web
  kubectl label ns simple-web istio-injection=enabled

  # 部署 simple-web 
  kubectl -nsimple-web apply -f simple-web.yaml
  # ---
  # deployment.apps/simple-web created
  # service/simple-web created

  # 创建证书
  openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=jrx Inc./CN=*.jrx.io' -keyout jrx.io.key -out jrx.io.crt
  kubectl create -n istio-system secret tls jrx-credential --key=jrx.io.key --cert=jrx.io.crt

  # 创建 gateway
  kubectl -nsimple-web apply -f simple-gw.yaml
  # ---
  # gateway.networking.istio.io/simple-gw created

  # 创建 virtual service
  kubectl -nsimple-web apply -f simple-vs.yaml 
  # ---
  # virtualservice.networking.istio.io/simple-vs created

  # 查看 istio service 信息
  # $INGRESS_IP = 10.100.103.155
  kubectl get svc -nistio-system

  # ---
  # NAME                   TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)                                                                      AGE
  # istio-egressgateway    ClusterIP      10.111.91.105    <none>        80/TCP,443/TCP                                                               3d21h
  # istio-ingressgateway   LoadBalancer   10.100.103.155   <pending>     15021:31954/TCP,80:31252/TCP,443:30443/TCP,31400:32716/TCP,15443:32556/TCP   3d21h
  # istiod                 ClusterIP      10.105.253.76    <none>        15010/TCP,15012/TCP,443/TCP,15014/TCP

  # 测试链接
  export INGRESS_IP=10.100.103.155
  curl --resolve simple.jrx.io:443:$INGRESS_IP https://simple.jrx.io/simple-web/time -v -k
  # ...
  # * Connection #0 to host simple.jrx.io left intact
  # "ok, TimeController"#
  ```