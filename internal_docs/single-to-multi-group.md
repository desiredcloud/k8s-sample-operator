https://kubebuilder.io/migration/multi-group.html

```shell
mkdir -p apis/image
mv api/* apis/image
rm -rf api
mkdir controllers/image
mv controllers/* controllers/image/
```