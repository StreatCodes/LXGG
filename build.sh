#!/bin/bash
echo Build server
pushd lxgg;
go build .;
popd;
pushd lxgg_frontend;
npm run prod;
popd;
rm -rf build/;
mkdir build;
mkdir build/static;
cp lxgg/lxgg build;
cp -R lxgg_frontend/css build/static/;
cp -R lxgg_frontend/icons build/static/;
cp -R lxgg_frontend/index.html build/static/;
cp -R lxgg_frontend/js build/static/;
