FROM ubuntu:focal

ENV HOME=/home/cudo

USER root
RUN  apt-get update -y && apt-get install -y curl wget net-tools sudo jq 
RUN useradd -m cudo && echo "cudo:cudo" | chpasswd && adduser cudo sudo

COPY ./build/cudo-linux-amd64 $HOME/build/cudos-noded
COPY ./scripts/run-node.sh $HOME/run-node.sh

RUN echo PATH="$HOME/bin:$PATH" >> $HOME/.bashrc

USER cudo
WORKDIR $HOME
CMD /bin/bash