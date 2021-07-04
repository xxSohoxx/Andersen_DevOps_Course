# About the task #
In this task we need to package our app from the [Ansible task](https://github.com/xxSohoxx/Andersen_DevOps_Course/tree/main/ansible_task) to Docker container.  
The lighter image provides us more poitns. Image must be less than 100 Mb.  
Application needs to listen port 8080 and HTTP only.

### How it works ###  
In order to run it on your working station you need to have Docker [installed](https://docs.docker.com/get-docker/).  
Then you need to copy only the Dockerfile from this repository and build an Docker image by running the command below from the directory where the Dockerfile is located:  
~~~
docker build --tag YOUR_IMAGE_NAME:TAG .
~~~
You should put your image name and tag instead "YOUR_IMAGE_NAME:TAG".  
Size of the image will be **54 MB** here.  
To start container you just need to run the command below:  
~~~
docker run YOUR_IMAGE_NAME:TAG  
~~~
