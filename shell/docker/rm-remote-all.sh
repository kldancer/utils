#! /bin/bash

image_name=$1


echo "image_name:     $image_name"
echo "********"

for tag in $(curl -sX GET "http://deploy.bocloud.k8s:40443/v2/$image_name/tags/list" | jq -r '.tags[]'); do
  echo " "
  echo "--------------------------"
  digest=$(curl -s --header "Accept: application/vnd.docker.distribution.manifest.v2+json" -I -X GET "http://deploy.bocloud.k8s:40443/v2/$image_name/manifests/$tag" | grep Docker-Content-Digest | awk -F: '{print $2}' | tr -d '[:space:]');
  url="http://deploy.bocloud.k8s:40443/v2/$image_name/manifests/$digest"
  echo "tag:            $tag"
  echo "digest:         $digest"
  echo "url:            $url"
  echo "curl -X DELETE  $url"
  curl -X DELETE  "$url"
done

echo "********"
echo "finish rm $image_name "