FROM registry.cn-hangzhou.aliyuncs.com/bocloud/busybox:1.35
ARG TARGETARCH
ARG BIN_NAME=aios-train-job

WORKDIR /
COPY bin/${BIN_NAME}-${TARGETARCH} ./${BIN_NAME}
ENTRYPOINT ["/${BIN_NAME}"]