#!/usr/bin/bash
set -e -o pipefail

GH_VERSION=""
gh_version(){
  if command -v gh &> /dev/null; then
    GH_VERSION=$(gh version | grep -o "v[0-9]\+\.[0-9]\+\.[0-9]")
  else
    echo "gh not found"; exit 1
  fi
}

setup_inbox(){
  gh extension install meiji163/gh-mail
  gh repo create inbox\
    --public --clone\
    --template meiji163/inbox 

  PUB_KEY=$(gh mail keygen | grep "public.pem")
  cat $PUB_KEY > ./inbox/public.pem
}

gh_version
if [[ $GH_VERSION < "v2.0.0" ]]; then
  echo "gh version v2.0.0 or later required"
fi

setup_inbox

echo "Adding public key to inbox"
cd inbox
git add public.pem 
git commit -m "adding public key"
git push origin HEAD
cd -
