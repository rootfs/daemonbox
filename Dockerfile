From centos
Maintainer Huamin Chen hchen@redhat.com

ADD daemonbox /daemonbox
EXPOSE 5000
ENV SERVICE_PORT=5000
ENTRYPOINT /daemonbox