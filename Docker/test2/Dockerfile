FROM ubuntu:22.04
RUN mkdir -p /scripts
# COPY install_docker_engine.sh /scripts
ADD https://raw.githubusercontent.com/matenduel/cloudwave/main/docker-practice/docker_install_script.sh  scripts/docker_install_script.sh
WORKDIR /scripts
# RUN chmod +x install_docker_engine.sh
RUN chmod +x ./docker_install_script.sh
# RUN ./install_docker_engine.sh
RUN ./docker_install_script.sh
RUN apt-get install -y docker-ce docker-ce-cli