# go-sample-api

```
# build linux/amd64 binary
./ci/build.sh

# deploy to PCF
cf push go-sample-api --random-route -b binary_buildpack --no-manifest -c './dist/go-sample-api-linux'
```
