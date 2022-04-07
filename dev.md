# kill port
sudo kill -9 `sudo lsof -t -i:5000`

git push origin dev-dik && git checkout development && git merge dev-dik && git push origin development 