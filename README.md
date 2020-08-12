# DWG Transformer
This repository aims to build a conversion tool for .dwg files (Autocad) based on microservice architecture.
The microservice called "transformer" is the main microservice. It is based on Golang technology (as well as most of microservices here except the python microservice which basically perform different validity checking on the output .dwf file)   It is strongly based on the LibreDWG (GNU licensed) project :
https://github.com/LibreDWG/libredwg
