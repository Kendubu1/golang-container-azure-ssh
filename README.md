# Golang Container w/ SSH Enabled on Azure App Service | Proof of concept
This sample demonstrate the installation differences when installing SSH on Debian & Alpine base Linux distributions. A seperate startup script, sshd_config & dockerfile is used for each build & this sample leverages a startup script as the entrypoint

### Repository Modifications 

| Files             |  Content                                   |
|----------------------|--------------------------------------------|
| `Dockerfile`           | Includes files for an alpine or debian based image           |
| `sshd_config_*`       | OpenSSH SSH daemon configuration file. Requires SSH_PORT ENV in dockerfile                      |
| `*_start.sh`               | Initalization Script to parse port number, pass env variables, start SSH & the app.                                 |

SSH Reference: https://docs.microsoft.com/bs-latn-ba/Azure/app-service/containers/configure-custom-container#enable-ssh
Configure Custom Containers: https://docs.microsoft.com/en-us/azure/app-service/configure-custom-container?pivots=container-linux


## Build Alpine Container Locally | localhost:4000
⚠ Make sure the start.sh shell script is saved with LF line-endings on your local machine before building in docker.
```
$ git clone https://github.com/Kendubu1/golang-container-azure-ssh.git
$ cd golang-container-azure-ssh
$ mv Dockerfile_alp Dockerfile
$ docker build -t <tag-name> .
$ docker run -d -p 4000:80 <tag-name>
```

## Build Debian Container Locally | localhost:4000
⚠ Make sure the start.sh shell script is saved with LF line-endings on your local machine before building in docker.
```
$ git clone https://github.com/Kendubu1/golang-container-azure-ssh.git
$ cd golang-container-azure-ssh
$ mv Dockerfile_deb Dockerfile
$ docker build -t <tag-name> .
$ docker run -d -p 4000:80 <tag-name>
```
