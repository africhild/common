# common package for Go (Gin + Gorm) APIs
## Common

### Introduction
A collection of reusable utilities for building APIs 

### Usage
```bash
# Published version
go get github.com/africhild/common

# Local development
git clone github.com/africhild/common
go mod edit -replace=github.com/africhild/common@v0.0.0-unpublished=../common
go get github.com/africhild/common@v0.0.0-unpublished
```

### Structure
1. `pkg` - reusables for config, dbs, repository controller etc

### TODO
1. Add unit tests