FROM debian:10
RUN apt update && apt install -y git
RUN git clone https://github.com/xxSohoxx/Andersen_DevOps_Course.git


FROM python:3.7.11-alpine3.13
RUN pip3 install markupsafe werkzeug jinja2 click itsdangerous Flask
WORKDIR /opt/animals_app/
COPY --from=0 ./Andersen_DevOps_Course/docker_task/app.py /opt/animals_app/app.py

CMD ["python3", "./app.py"] 

