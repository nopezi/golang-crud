docker network create \
  --driver=bridge \
  --subnet=162.28.0.0/16 \
  --ip-range=162.28.5.0/24 \
  --gateway=162.28.5.254 \
  mysql_default

  docker network rm my-network