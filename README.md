### Template Generator
Given a manifest folder path, it generate chart to deploy into K8s.

#### Running App
Load a configuration file with details from config.yml
```
go mod vendor
go build
./template-gen -d sample-manifest -e test -n account -r Release-1
```

Based on the release manifest it generates the following for the each application listed in the release manifest

```
- service
- service-account
- deployment
```