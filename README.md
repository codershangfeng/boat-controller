# boat-controller

## How to curl Local Docker Kubernetes api-server

The mTLS is enabled on Kubernetes api-server.
There are **2** certificate and **1** key file needed:

1. CA certificate belonging to the CA that signed the server’s certificate
1. Client certificate
1. Client private key

We could find these information in the `${HOME}/.kube/config`, in which these information are encoded with Base64.

For example, 

1. CA cert -> `.clusters[0].cluster.certificate-authority-data`
1. Client cert -> `.users[0].user.client-certificate-data`
1. Client private key -> `.users[0].user.client-key-data`

You can use `yq` and `base64` to decode it, i.e.:
```bash
yq e '.users[0].user.client-certificate-data' ~/.kube/config | base64 -d > client.crt
```

After then, you can send watch request to the target kubernetes api-server by:
```bash
curl -i --cacert ca.crt --key client.key --cert client.crt https://localhost:6443/api/v1/watch/pods\?watch\=yes
```
With response:
```bash
HTTP/2 200
cache-control: no-cache, private
content-type: application/json
date: Wed, 23 Jun 2021 09:16:17 GMT
```

More details: 
1. [how to curl an endpoint protected by mutual tls (mtls)](https://downey.io/notes/dev/curl-using-mutual-tls/)