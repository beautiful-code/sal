### Setup host machine
- To install minikube follow the instructions from this page https://github.com/kubernetes/minikube/releases/latest
- Run `eval $(minikube docker-env)` to switch your docker node to the one that is running on minikube cluster. You need to do this only if you are planning to build docker images locally and use them for creating containers within the minikube cluster.
- Run `kubectl config use-context minikube` to switch context to minikube cluster that is setup on your host machine.

### Creating k8s deployments
0. Create Persistent Storage volume for mysql server
  - Run `kubectl create -f mysql-pv-volume.yml`
  - Run `kubectl create -f mysql-pv-claim.yml`

1. Create the mysql-server deployments on the cluster.
  - Run `kubectl create -f sal-mysql-deployment.yaml`
  - Create a k8s service to expose the mysql-server
    - Run `kubectl expose deployment mysql-server --type=LoadBalancer`.
    - Run `kubectl get service mysql-server` and note down the internal/cluster IP.

2. Log on to the running `sal-mysql-*` container.
   - Run `kubectl exec -it sal-mysql-* bash` and run the below commands within the container.
   - $ mysql -uroot -p # Password: secret
   - $ mysql > `CREATE USER 'sal-user'@'localhost' IDENTIFIED BY 'sal-secret';`
   - $ mysql > `GRANT ALL PRIVILEGES ON *.* TO 'sal-user'@'localhost' WITH GRANT OPTION;`
   - $ mysql > `CREATE USER 'sal-user'@'%' IDENTIFIED BY 'sal-secret';`
   - $ mysql > `GRANT ALL PRIVILEGES ON *.* TO 'sal-user'@'%' WITH GRANT OPTION;`
   - $ mysql > exit
   - $ mysql -u sal-user -p
   - $ mysql > `create database sal_staging;`

3. Update `services/user/config.json.staging` with the above used credentials and
build the docker image. Update the k8s deployment yaml file if necessary.

4. Create the sal-user deployments on the cluster.
  - Run `kubectl create -f sal-user-deployment.yaml`
  - Create a k8s service to expose the sal-user app
    - Run `kubectl expose deployment sal-user --type=LoadBalancer`.
    - Run `kubectl get service sal-user` and note down the internal/cluster IP.

5. Update `services/application/config.json.staging` with the above used credentials and
build the docker image. Update the k8s deployment yaml file if necessary.

6. Create the sal-application deployments on the cluster.
  - Run `kubectl create -f sal-user-deployment.yaml`
  - Create a k8s service to expose the sal-application app
    - Run `kubectl expose deployment sal-application --type=LoadBalancer`.

7. Accessing the services from your host machine.
  - Run `minikube ip` to get the k8s cluster IP on your host machine.
  - Run `kubectl describe service sal-user sal-application` and note down the NodePort values for each service.
  - Open the browser and navigate to http://<minikube-ip-here>:<NodePort>
