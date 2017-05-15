tf-go-inception: A Go API server to serve Tensorflow Inception Model
==============

[![Join the chat at https://gitter.im/kkdai/tf-go-inception](https://badges.gitter.im/kkdai/tf-go-inception)](https://gitter.im/kkdai/tf-go-inception?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

 [![GoDoc](https://godoc.org/github.com/kkdai/tf-go-inception.svg?status.svg)](https://godoc.org/github.com/kkdai/tf-go-inception)  [![Build Status](https://travis-ci.org/kkdai/tf-go-inception.svg?branch=master)](https://travis-ci.org/kkdai/tf-go-inception.svg)

[![goreportcard.com](https://goreportcard.com/badge/github.com/kkdai/tf-go-inception)](https://goreportcard.com/report/github.com/kkdai/tf-go-inception)


![](https://github.com/kkdai/LineBotAnimal/blob/master/images/how_use.PNG?raw=true)

## How to build it

- Install Go 1.8
- Download tensorflow prebuild library

```
TF_TYPE="cpu" # Set to "gpu" for GPU support
curl -L \
  "https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-${TF_TYPE}-$(go env GOOS)-x86_64-1.1.0.tar.gz" |
sudo tar -C /usr/local -xz
```

- Change your LD_LIBRARY path `export LD_LIBRARY_PATH=/usr/local/lib/`
- Clone this repo
- Just build it `go build`
- Run this web server it will default port `3000`


## API List:

- `/api/v1/tf-image`: POST API to upload image in multipart "upload".
- `/api/v1/foo`:  Testing API and just response "bar".

## How to use it

TBC

License
---------------

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

