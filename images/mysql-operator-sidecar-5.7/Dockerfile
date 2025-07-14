###############################################################################
#  Docker image for sidecar containers
###############################################################################




FROM debian:bullseye-slim  as sidecar

RUN groupadd -g 999 mysql
RUN useradd -u 999 -r -g 999 -s /sbin/nologin \
    -c "Default Application User" mysql

COPY rootfs/ /

RUN apt-get update -y \
    && apt-get install -y --no-install-recommends \
    apt-transport-https ca-certificates unzip pigz wget gnupg \
    && rm -rf /var/lib/apt/lists/*


RUN export RCLONE_VERSION=1.57.0 \
    && gpg --import /etc/gpg-keys/rclone.gpg \
    && wget -nv https://github.com/ncw/rclone/releases/download/v${RCLONE_VERSION}/rclone-v${RCLONE_VERSION}-linux-amd64.zip \
    && wget -nv https://github.com/ncw/rclone/releases/download/v${RCLONE_VERSION}/SHA256SUMS \
    && gpg --verify --output=- SHA256SUMS > sums \
    && sha256sum -c --ignore-missing sums \
    && unzip rclone-*-linux-amd64.zip \
    && mv rclone-*-linux-amd64/rclone /usr/local/bin/ \
    && chmod 755 /usr/local/bin/rclone \
    && rm -r rclone-*-linux-amd64 rclone-*-linux-amd64.zip

# https://repo.percona.com/apt/dists/buster/ 看样子是有，但为什么安装不了
#RUN echo 'deb https://repo.percona.com/apt buster main' > /etc/apt/sources.list.d/percona.list

ARG XTRABACKUP_PKG=percona-xtrabackup-24
RUN apt-get update -y \
    && apt install curl -y \
    && curl -O https://repo.percona.com/apt/percona-release_latest.generic_all.deb \
    && apt install -y gnupg2 lsb-release ./percona-release_latest.generic_all.deb
RUN apt update -y \
    && percona-release setup pxb-80 -y \
    && apt install default-mysql-client -y

RUN apt install ${XTRABACKUP_PKG} percona-toolkit unzip -y \
    && rm -rf ./percona-release_latest.generic_all.deb \
    && rm -rf /var/lib/apt/lists/*

USER mysql
ENTRYPOINT ["/usr/local/bin/docker-entrypoint.sh"]
