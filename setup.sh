#!/usr/bin/bash
set -e -o pipefail

gh extension install meiji163/gh-mail
gh repo create inbox\
  --public --clone\
  --template meiji163/inbox 

PUB_KEY=$(gh mail keygen | grep "public.pem")
cat $PUB_KEY > ./inbox/public.pem

echo "Adding public key to inbox"
cd inbox
git add public.pem 
git commit -m "adding public key"
git push origin HEAD
cd -
