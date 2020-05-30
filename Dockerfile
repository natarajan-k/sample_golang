FROM openshift/base-centos7
RUN yum install -y golang && \
    yum clean all
ENV GOLANG_VERSION=1.9 \
    GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
COPY my_app /usr/local/bin/
RUN my_app
EXPOSE 8081
USER 1001
CMD ["my_app", "run"]
