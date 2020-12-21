# go-bazel

# Run application
```
$ bazel run //service
```

# Run test
```
$ bazel test //...
```

# Build binary
```
$ bazel build //service
```

# Build docker image
```
$ bazel build //service:service_image.tar
$ docker load -i ./bazel-bin/service/service_image.tar
```
