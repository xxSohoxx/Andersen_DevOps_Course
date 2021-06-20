# About the task and script #

We need to write a script which do the following:
- Check if there are open Pull Requests for a repository. An url in format https://github.com/$user/$repo will be passed to the script
- Print the list of the most productive contributors (authors of more than 1 open PR)
- Print the number of PRs each contributor has created with the labels

## How it works ##

You need to run a script /Andersen_DevOps_Course/git_task/git_script  
Then it will ask you for the link with github repository. You need to provide the url in format https://github.com/$user/$repo.  
There is a limit of API calls. In order to be able to get the info you need authenticate yourself at Github. In this script it is implemented via personal access tokens.  

You can get more information here:  
https://docs.github.com/en/rest/overview/other-authentication-methods#via-oauth-and-personal-access-tokens  
And how to create access tokens here:  
https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token  

## Below is the example of script execute: ##  

*Andersen_DevOps_Course/git_task$ ./git_script*
*Please enter the GitHub repository link: %REPO_URL%*  
*You also need to provide you GitHub Authorization token in order to send API requests: %YOUR_GITHUB_ACCESS_TOKEN%*    
*--------------------------------------------------------------------------------------------------------------*  
*We have found openned Pull Requests*  
*Below are the contributors who have created more than 1 PR:*  
*--------------------------------------------------------------------------------------------------------------*    
*Here will be the list of contributors*  
*--------------------------------------------------------------------------------------------------------------*  
  
*Here is the contributors who created PR with labels in format:*  
*%Number of PRs with Labels% and %Contributor%*  
*--------------------------------------------------------------------------------------------------------------*  
*Here will be the list of contributors who created PR with labels*  
*--------------------------------------------------------------------------------------------------------------*  



