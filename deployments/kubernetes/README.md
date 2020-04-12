# commands to run app

## Steps

1. Create Storage

```gcloud compute disks create --size=1GiB --zone=europe-west4-a mongodb```

k port-forward pod/mongodb  27017:27017

```
db
use examples
db.foo.insert({bar:'aa'})
db.foo.find({})
```