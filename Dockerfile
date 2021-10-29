FROM golang:1.16-stretch

#Install openssh-server w/ required credentials for Azure App Service
RUN apt-get update \
  && apt-get install -y --no-install-recommends openssh-server \
  && echo "root:Docker!" | chpasswd

#Copy SSH config file, startup script with 775 permissions
COPY sshd_config /etc/ssh/sshd_config
COPY deb_start.sh /usr/local/bin/
RUN chmod 755 /usr/local/bin/deb_start.sh
#Create enviorment variable for our sshd_config
ENV SSH_PORT 2222
#Expose ports needed for the sample app(80) & SSH on Azure(2222)
EXPOSE 80 2222

COPY main.go ./
#ENTRYPOINT ["go", "run", "main.go"]
ENTRYPOINT ["/usr/local/bin/deb_start.sh"]
