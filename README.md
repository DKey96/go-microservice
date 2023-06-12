# Go HTTPS Microservice

Microservice that returns a "Hello from the other side" message on `GET /`.

## How to use
- You need to create a SSL certificates first for the HTTPS
- Create a folder certs in the project root (`mkdir certs`)
- Navigate into it (`cd certs`) and run:
```bash
openssl genrsa -out localhost.localhost.key 2048
openssl req -new -sha256 -key localhost.localhost.key -out csr.csr
openssl req -x509 -sha256 -days 365 -key localhost.localhost.key -in csr.csr -out localhost.localhost.cert
openssl req -in csr.csr -text -noout | grep -i "Signature.*SHA256" && echo "All is well" || echo "This certificate will stop working in 2017! You must update OpenSSL to generate a widely-compatible certificate"
```
### Locally
- To run locally, you have to export some .env variables
  - **DK_CERT_FILE** - `export DK_CERT_FILE=<PATH_TO_CERTS_FOLDER>/localhost.localhost.cert`
  - **DK_KEY_FILE** - `export DK_KEY_FILE=<PATH_TO_CERTS_FOLDER>/localhost.localhost.key`
  - **DK_SERVICE_ADDR** - `export DK_SERVICE_ADDR=":8080"`
- Run 
```bash
go run .
```
- Open `https://localhost:8080` or use curl:
```bash
curl https://127.0.0.1:8000 -i -k 
```
### Docker
- First build the image (bear in mind the `dot` in the end)
```bash
docker build -t gomicroservice:microsrvice -f Dockerfile .
```
- Run the docker container with built image
```bash
docker run --publish 8080:8080 gomicroservice:microsrvice
```
