# GoHTML2PDF
[![Publish Docker image](https://github.com/Binozo/GoHTML2PDF/actions/workflows/deploy.yaml/badge.svg)](https://github.com/Binozo/GoHTML2PDF/actions/workflows/deploy.yaml)
## An easy-to-use HTTP endpoint which converts your html to a pdf file with state-of-the-art css support using chromium.

### Setup
1. Pull the docker image
```bash
docker pull binozoworks/gohtml2pdf:latest
```

2. Run it
```bash
docker run -d -p 7524:7524 --name gohtml2pdf binozoworks/gohtml2pdf:latest 
```

3. Convert your html to pdf:

```bash
curl --location --request POST 'http://localhost:7524/convert' --form 'html="<h1>Hi</h1>"' --output mypdf.pdf
```