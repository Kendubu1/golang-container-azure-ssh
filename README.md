# golang Container w/ SSH Enabled on Azure App Service
This sample leverages a startup script as the entrypoint. This script allows us to start the SSH service & golang container.

### Repository Modifications 

| Files             |  Content                                   |
|----------------------|--------------------------------------------|
| `Dockerfile`           | Includes files for an alpine or debian based image           |
| `sshd_config_*`       | OpenSSH SSH daemon configuration file. Requires SSH_PORT ENV in dockerfile                      |
| `*_start.sh`               | Initalization Script to parse port number, pass env variables, start SSH & the app.                                 |


SSH Reference: https://docs.microsoft.com/bs-latn-ba/Azure/app-service/containers/configure-custom-container#enable-ssh
Configure Custom Containers: https://docs.microsoft.com/en-us/azure/app-service/configure-custom-container?pivots=container-linux


