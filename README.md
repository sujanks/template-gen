### Template Generator
Given a manifest folder path, it generate chart to deploy into K8s.

#### Running App
Load a configuration file with details from config.yml
```
go mod vendor
go build
./template-gen -d /manifest -e test -n account -r releaseName
```