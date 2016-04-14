^(｀（ ● ●）´ )^

A tool for managing simultaneously several git repositories.

Some typical usages:

       got status
          Run git status on all Git repos in the working dir

       got status @REPO1
          Run git status on all Git repos in the working dir except REPO1

       got @REPO1 status
          Run git status only on REPO1

       got checkout A_branch @REPO1 checkout Another_branch
          Run git checkout A_branch on all Git repos except REPO01 and git checkout Another_branch on REPO1
