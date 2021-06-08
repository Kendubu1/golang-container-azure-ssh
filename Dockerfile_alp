FROM golang:1.15-alpine

#Install openssh-server w/ required credentials for Azure App Service
RUN apk add --update openssh-server tzdata openrc openssl \
        && echo "root:Docker!" | chpasswd \
        && rm -rf /var/cache/apk/* \
        # Remove unnecessary services
        && rm -f /etc/init.d/hwdrivers \
                 /etc/init.d/hwclock \
                 /etc/init.d/mtab \
                 /etc/init.d/bootmisc \
                 /etc/init.d/modules \
                 /etc/init.d/modules-load \
                 /etc/init.d/modloop \
        # Can't do cgroups
        && sed -i 's/\tcgroup_add_service/\t#cgroup_add_service/g' /lib/rc/sh/openrc-run.sh \
        && apk add --no-cache bash; 

#Copy SSH config file, startup script with 775 permissions
COPY sshd_config_alp /etc/ssh/sshd_config_alp
COPY alp_start.sh /usr/local/bin/
RUN chmod 755 /usr/local/bin/alp_start.sh
#Create enviorment variable for our sshd_config
ENV SSH_PORT 2222
#Expose ports needed for the sample app(80) & SSH on Azure(2222)
EXPOSE 80 2222

COPY main.go ./
#ENTRYPOINT ["go", "run", "main.go"]
ENTRYPOINT ["/usr/local/bin/alp_start.sh"]
