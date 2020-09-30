#!/bin/bash
imageName=degauss
containerName=clever_gauss

# Delete Old Container
echo [!] Delete Old Container
docker rm -f $containerName

# Rebuild Image
echo [!] Rebuild Image
docker build -t $imageName  .

# Run New Container
echo [!] Run New Container
docker run -d --name $containerName $imageName

# Running
echo [!] Container running