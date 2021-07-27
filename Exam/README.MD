# About the Exam #

We need to create a working infrastructure with automated building and deployment of an application by using Docker.  
The solution should consist of 2 **web** applications which return **"Hello world 1"** and **"Hello world 2"**. We need to choose 2 programming languages from the list below:  

* Python
* .NET Core
* php
* Java
* Go (gin framework)

We are free to choose any CI/CD, infrastructure, Registry and Control version tools. 

# Result #

Application is running on AWS platform and can be accessed by the links below:  
http://ec2-54-93-104-116.eu-central-1.compute.amazonaws.com:8001/  
http://ec2-54-93-104-116.eu-central-1.compute.amazonaws.com:8002/  

# How it works #

### Repository, CI/CD ###
For the git repository and CI/CD I choose GitLab solution.  
Below are the repositories I created for the applications:

* https://gitlab.com/andersen_devops_course/andersen_py_app - Python Web application
* https://gitlab.com/andersen_devops_course/andersen_go_app - Go Web application


### Infrastructure ###

The application is deployed on AWS EC2 instance. OS was chosen - Debian Linux.  
Infrastructure preparation and configuring is presented on separate repository - https://gitlab.com/andersen_devops_course/andersen_infra_deploy

### Additions ###

* Notifications are configured via GitLab built-in tool:

![Notifications](https://user-images.githubusercontent.com/75803382/127167421-0fb1577a-fc57-42c9-bb54-5017d5efe401.png)

* 
