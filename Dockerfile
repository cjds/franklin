ARG SOURCE_STAGE=git
FROM golang:latest as src_git
ARG GIT_RANDOMIZER
# Authorize ssh private key
RUN mkdir -p /root/.ssh/
ARG REMOTE=cjds
ARG REPOSITORY=franklin
ARG SSH_PRIVATE_KEY
RUN echo "$SSH_PRIVATE_KEY" > /root/.ssh/id_rsa
ARG SSH_KNOWN_HOSTS
RUN echo "$SSH_KNOWN_HOSTS" > /root/.ssh/known_hosts
RUN chmod 400 /root/.ssh/id_rsa
WORKDIR /
RUN git clone git@github.com:$REMOTE/$REPOSITORY.git

FROM golang:latest as src_local
ARG REPOSITORY=franklin
COPY . /$REPOSITORY

FROM src_$SOURCE_STAGE AS stageforcopy


FROM golang:latest as dev_env_stage
RUN DEBIAN_FRONTEND=noninteractive  apt-get update -y && apt-get install vim byobu -y

FROM dev_env_stage
ARG REPOSITORY=franklin
COPY --from=stageforcopy /$REPOSITORY /go/src/$REPOSITORY
RUN go get -u github.com/golang/dep/cmd/dep
RUN curl -sL https://deb.nodesource.com/setup_10.x | bash -
RUN DEBIAN_FRONTEND=noninteractive  apt-get update -y && apt-get install npm build-essential libgl1-mesa-dev xorg-dev -y

WORKDIR /go/src/$REPOSITORY
RUN cd ui && npm install
RUN pip install -r requirements.txt
RUN cd boston && dep ensure
