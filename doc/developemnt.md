# how to development mysql-operator

## add tool for deveoping

```
make build.tools
```

## add new crd 

```
 bin/kubebuilder create api --group mysql --version v1alpha1 --kind MysqlRestore --controller=true --resource=true --make=true
```

## generate code and generate manifest 

```
make kubebuilder.code;make kubebuild.manifests
```
