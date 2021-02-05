sudo docker run -itd \
    --name 12306-container \
    -v $PWD:/src \
    --rm \
    --workdir /src \
    -p 8089:8080 \
    machengcheng/12306:v0.1
