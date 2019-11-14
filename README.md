```
./app -p /mnt/extra/hk 
```
or 
```
docker run -d -v /mnt/extra/hk:/data -p 8089:80 matosama/receive:arm64v8 -p /data
```
