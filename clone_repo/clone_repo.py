'''
import git
def CloneRepo(url, _branch = 'master'):
     
    path_to_clone = input("Give Path For Repo to be Clones to")
    permission_to_clone = input("Give permission to Clone Y/N")
    if permission_to_clone == 'Y' or permission_to_clone == 'Yes' or permission_to_clone == 'YES' or permission_to_clone == 'yes':
        repo = git.Repo.clone_from(url,
                           path_to_clone,
                           branch=_branch)
'''