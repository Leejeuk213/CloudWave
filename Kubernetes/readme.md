### 요약 

  k8s 환경과 서비스를 공부 yaml 파일 직접 작성해가며 local 환경에서 cli를 통해 kubectl 명령어를 통해 실습  

  ide에서 terraform을 활용하여 aws eks Cluster 활성화 eks 환경에서 실습  

---

### 실습 내용
* docker compose 활용 kubectl 실습할 웹 ide 환경 구성
* k8s plane, worker들은 kind 네트워크로 구성   
* pod 생성 실습  
* liveness probe 실습
* Replication Controller 실습
* Replicaset 실습
* Deamon set 실습
* Deployment 실습
* terraform 활용 aws eks cluster 활성화 eks 환경 구축
* cluster ip portforwarding을 통해 연결 실습
* nodeport 활용 replica된 pod들 동시 연결 실습
* eks 환경에서 aws load balance 활용하여 접속 실습
* local kind 네트워크 환경에서 metallb 활용하여 접속 실습
* eks 환경에서 ingress 실습
* local kind 네트워크 환경에서 nginx 컨트롤러 설치한 후 ingress 실습
* aws volume 실습
* pv, pvc 직접 생성하고 pod 생성 실습
* pvc만으로 pod 생성 실습
* sc를 직접 만들고 pvc애 기재해 pod 생성 실습
* configmap, secret을 이용하여 pod 생성 실습
* 자바 스프링부트를 이용해 만든 웹사이트 깃 액션으로 배포하고 aws와 argocd 활용하여 cicd 환경 구축 실습
* helm 실습
