FROM python:2.7-stretch
MAINTAINER jkouassi
RUN groupadd -r jojo -g 436 && useradd -u 435 -r -g jojo -s /sbin/nologin -c "utlisateur dockerfile" jojo
RUN apt-get update -y && apt-get install gcc python3-dev libsasl2-dev libldap2-dev libssl-dev git -y
RUN pip install flask==1.1.2 flask_httpauth==4.1.0 flask_simpleldap python-dotenv==0.14.0
EXPOSE 5000
VOLUME ["/data"]
ADD student_age.py /
CMD [ "python", "./student_age.py" ]
USER jojo