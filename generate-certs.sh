#!/bin/bash
#Courtesy of https://github.com/simplesteph/grpc-go-course/blob/master/ssl/instructions.sh

# Summary 
# Private files: ca.key, server.key, server.pem, server.crt
# "Share" files: ca.crt (needed by the client), server.csr (needed by the CA)

# Changes these CN's to match your hosts in your environment if needed.
SERVER_CN=$1
CERT_PREFIX=$2
CERT_PPHRASE=$3

# Step 1: Generate Certificate Authority + Trust Certificate (ca.crt)
openssl genrsa -passout ${CERT_PPHRASE} -des3 -out ${CERT_PREFIX}.ca.key 4096
openssl req -passin ${CERT_PPHRASE} -new -x509 -days 3650 -key ${CERT_PREFIX}.ca.key -out ${CERT_PREFIX}.ca.crt -subj "/CN=${SERVER_CN}"

# Step 2: Generate the Server Private Key (server.key)
openssl genrsa -passout ${CERT_PPHRASE} -des3 -out ${CERT_PREFIX}.server.key 4096

# Step 3: Get a certificate signing request from the CA (server.csr)
openssl req -passin ${CERT_PPHRASE} -new -key ${CERT_PREFIX}.server.key -out ${CERT_PREFIX}.server.csr -subj "/CN=${SERVER_CN}"

# Step 4: Sign the certificate with the CA we created (it's called self signing) - server.crt
openssl x509 -req -passin ${CERT_PPHRASE} -days 3650 -in ${CERT_PREFIX}.server.csr -CA ${CERT_PREFIX}.ca.crt -CAkey ${CERT_PREFIX}.ca.key -set_serial 01 -out ${CERT_PREFIX}.server.crt 

# Step 5: Convert the server certificate to .pem format (server.pem) - usable by gRPC
openssl pkcs8 -topk8 -nocrypt -passin ${CERT_PPHRASE} -in ${CERT_PREFIX}.server.key -out ${CERT_PREFIX}.server.pem