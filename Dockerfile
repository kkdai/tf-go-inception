FROM ubuntu
RUN apt-get update
#RUN apt-get install libcupti-dev
RUN apt-get install python-pip python-dev -y
RUN pip install --upgrade tensorflow

